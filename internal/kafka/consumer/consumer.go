package consumer

import (
	"context"
	"encoding/json"
	"log"
	"wb/internal/model"

	"github.com/segmentio/kafka-go"
)

type OrderProcessor interface {
	ProcessOrder(ctx context.Context, order model.Order) error
}

type Consumer struct {
	reader    *kafka.Reader
	processor OrderProcessor
}

func NewConsumer(brokers []string, topic, groupID string, processor OrderProcessor) *Consumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: groupID,
	})
	return &Consumer{
		reader:    reader,
		processor: processor,
	}
}

func (c *Consumer) Start(ctx context.Context) {
	defer c.reader.Close()

	for {
		msg, err := c.reader.ReadMessage(ctx)
		if err != nil {
			log.Printf("Kafka read error: %v", err)
			continue
		}

		var order model.Order
		if err := json.Unmarshal(msg.Value, &order); err != nil {
			log.Printf("Invalid message: %v", err)
			c.reader.CommitMessages(ctx, msg)
			continue
		}

		if err := c.processor.ProcessOrder(ctx, order); err != nil {
			log.Printf("Process order error: %v", err)
			continue
		}

		c.reader.CommitMessages(ctx, msg)
		log.Printf("Processed order: %s", order.OrderUID)
	}
}
