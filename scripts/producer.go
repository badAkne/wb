package main

import (
	"context"
	"encoding/json"
	"log"

	"wb/internal/model"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/segmentio/kafka-go"
)

func main() {
	w := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "orders",
		Balancer: &kafka.LeastBytes{},
	}

	order := model.Order{
		OrderUID:    gofakeit.UUID(),
		TrackNumber: gofakeit.UUID(),
		Entry:       gofakeit.BeerName(),
		Delivery: model.Delivery{
			Name:    gofakeit.Name(),
			Phone:   gofakeit.Phone(),
			Zip:     gofakeit.Zip(),
			City:    gofakeit.Address().City,
			Address: gofakeit.Address().Address,
			Region:  gofakeit.Address().State,
			Email:   gofakeit.Email(),
		},
		Payment: model.Payment{
			Transaction:  gofakeit.UUID(),
			Currency:     gofakeit.CurrencyShort(),
			RequestId:    gofakeit.FarmAnimal(),
			Provider:     gofakeit.BankType(),
			Amount:       int64(gofakeit.Number(1000, 5000)),
			Paymentdt:    gofakeit.Int64(),
			Bank:         gofakeit.BankName(),
			DeliveryCost: int64(gofakeit.Number(200, 1000)),
			GoodsTotal:   int64(gofakeit.Number(0, 500)),
			CustomFee:    int64(gofakeit.Number(100, 500)),
		},
		Items: []model.Item{
			{
				ChrtId:      gofakeit.Int64(),
				TrackNumber: gofakeit.DigitN(10),
				Price:       int64(gofakeit.Product().Price),
				Rid:         gofakeit.DigitN(100),
				Name:        gofakeit.ProductName(),
				Sale:        int64(gofakeit.Number(100, 10000)),
				Size:        gofakeit.RandomString([]string{"xs", "s", "m", "l", "xl"}),
				TotalPrice:  int64(gofakeit.Number(1, 1000)),
				NmId:        int64(gofakeit.Number(1, 100000)),
				Brand:       gofakeit.Company(),
				Status:      int64(gofakeit.HTTPStatusCode()),
			},
		},
		Locale:            gofakeit.CountryAbr(),
		InternalSignature: "",
		CustomerId:        gofakeit.DigitN(10),
		DeliveryService:   gofakeit.Company(),
		ShardKey:          gofakeit.DigitN(100),
		SmId:              int64(gofakeit.Number(1, 100)),
		DateCreated:       "2021-11-26T06:22:19Z",
		OofShard:          gofakeit.DigitN(10),
	}

	log.Printf(order.OrderUID)
	data, _ := json.Marshal(order)
	err := w.WriteMessages(context.Background(), kafka.Message{Value: data})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Message sent")
}
