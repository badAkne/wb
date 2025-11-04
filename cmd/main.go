package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	config "wb/internal/config"
	envConfig "wb/internal/config/env"
	orderHandlers "wb/internal/handlers/order"
	orderKafka "wb/internal/kafka/consumer"
	orderRepository "wb/internal/repository/order"
	orderService "wb/internal/service/order"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	mainCtx, cancel := context.WithCancel(context.Background())

	sigCh, stop := signal.NotifyContext(mainCtx, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer stop()

	if err := config.Load(".env"); err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	pgCfg, err := envConfig.NewPGConfig()
	if err != nil {
		log.Fatalf("error loading pg config: %v", err)
	}

	httpCfg, err := envConfig.NewHTTPConfig()
	if err != nil {
		log.Fatalf("error loading http config: %v", err)
	}

	kafkaCfg, err := envConfig.NewKafkaConfig()
	if err != nil {
		log.Fatalf("error loading kafka config: %v", err)
	}

	db, err := pgxpool.Connect(mainCtx, pgCfg.DSN())
	if err != nil {
		log.Fatal(err)
	}

	repo := orderRepository.NewOrderRepository(db)
	if err := repo.LoadCacheFromDB(mainCtx); err != nil {
		log.Printf("Error loading cache: %v", err)
	}

	svc := orderService.NewOrderService(repo)

	consumer, err := orderKafka.NewConsumer(kafkaCfg.Brokers(), kafkaCfg.Topic(), kafkaCfg.GroupID(), &svc)
	if err != nil {
		log.Printf("err occured while making kafka consumer: %v", err)
	}

	go consumer.Start(mainCtx)

	r := mux.NewRouter()
	h := orderHandlers.NewOrderHandlers(&svc)
	r.HandleFunc("/order/{uid}", h.GetOrder).Methods("GET")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("web")))

	srv := &http.Server{
		Addr:    httpCfg.Adress(),
		Handler: r,
	}

	log.Printf("Starting HTTP server on :%s\n", httpCfg.Adress())

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("error while trying closing http server: %v", err)
		}
	}()

	<-sigCh.Done()
	stop()

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := consumer.Stop(); err != nil {
		log.Printf("%v", err)
	}

	srv.Shutdown(shutdownCtx)
	db.Close()
	cancel()

	log.Printf("server shutdown graccefully")
}
