package consumer

import "context"

type KafkaConsumer interface {
	StartConsuming(ctx context.Context)
}
