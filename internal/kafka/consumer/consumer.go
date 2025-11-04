package consumer

import (
	"context"
	"encoding/json"
	"log"
	"strings"
	"wb/internal/help"
	"wb/internal/model"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type OrderProcessor interface {
	ProcessOrder(ctx context.Context, order model.Order) error
}

type Consumer struct {
	consumer  *kafka.Consumer
	processor OrderProcessor
	stop      bool
}

func NewConsumer(brokers []string, topic, groupID string, processor OrderProcessor) (*Consumer, error) {
	cfg := &kafka.ConfigMap{
		"group.id":          groupID,
		"bootstrap.servers": strings.Join(brokers, ","),
	}

	c, err := kafka.NewConsumer(cfg)
	if err != nil {
		return nil, err
	}

	if err := c.Subscribe(topic, nil); err != nil {
		return nil, err
	}

	/*
		reader := kafka.NewConsumer(kafka.ReaderConfig{
			Brokers: brokers,
			Topic:   topic,
			GroupID: groupID,
		})
	*/

	return &Consumer{
		consumer:  c,
		processor: processor,
	}, nil
}

func (c *Consumer) Start(ctx context.Context) {

	for {

		if c.stop {
			break
		}

		kafkaMsg, err := c.consumer.ReadMessage(0)
		if err != nil {
			log.Printf("%v", err)
		}

		if kafkaMsg == nil {
			continue
		}

		var order model.Order
		if err := json.Unmarshal(kafkaMsg.Value, &order); err != nil {
			log.Printf("Invalid message: %v", err)
			c.consumer.CommitMessage(kafkaMsg)
			continue
		}

		if err := help.ValidateOrder(order); err != nil {
			log.Printf("invalid_order_parametrs:%s\n", err.Error())
			continue
		}

		if err := c.processor.ProcessOrder(ctx, order); err != nil {
			log.Printf("Process order error: %v", err)
			continue
		}

		if _, err := c.consumer.CommitMessage(kafkaMsg); err != nil {
			log.Printf("Error while commiting a message: %v", err)
			continue
		}

		log.Printf("Processed order: %s", order.OrderUID)
	}

	/*
		for {

			if c.stop {
				break
			}

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

			if err := help.ValidateOrder(order); err != nil {
				log.Printf("invalid_order_parametrs:%s\n", err.Error())
				continue
			}

			if err := c.processor.ProcessOrder(ctx, order); err != nil {
				log.Printf("Process order error: %v", err)
				continue
			}

			if err := c.reader.CommitMessages(ctx, msg); err != nil {
				log.Printf("Error while commiting a message: %v", err)
				continue
			}

			log.Printf("Processed order: %s", order.OrderUID)
		}
	*/
}

func (c *Consumer) Stop() error {
	c.stop = true
	return c.consumer.Close()
}
