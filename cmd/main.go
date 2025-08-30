package main

import (
	"context"
	"log"
	"net/http"

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

	db, err := pgxpool.Connect(context.Background(), pgCfg.DSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := orderRepository.NewOrderRepository(db)
	if err := repo.LoadCacheFromDB(context.Background()); err != nil {
		log.Printf("Error loading cache: %v", err)
	}

	svc := orderService.NewOrderService(repo)
	consumer := orderKafka.NewConsumer(kafkaCfg.Brokers(), kafkaCfg.Topic(), kafkaCfg.GroupID(), &svc)
	go consumer.Start(context.Background())

	r := mux.NewRouter()
	h := orderHandlers.NewOrderHandlers(&svc)
	r.HandleFunc("/order/{uid}", h.GetOrder).Methods("GET")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("web")))

	log.Printf("Starting HTTP server on :%s\n", httpCfg.Adress())
	log.Fatal(http.ListenAndServe(httpCfg.Adress(), r))
}
