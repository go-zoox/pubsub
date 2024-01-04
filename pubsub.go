package pubsub

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// PubSub is a simple pubsub.
type PubSub interface {
	Publish(ctx context.Context, msg *Message) error
	Subscribe(ctx context.Context, topic string, handler Handler) error
}

type Message struct {
	Topic string
	Body  []byte
}

type Handler func(msg *Message) error

type pubsub struct {
	client *redis.Client
}

type Config struct {
	RedisHost     string
	RedisPort     int
	RedisUsername string
	RedisPassword string
	RedisDB       int
}

func New(cfg *Config) PubSub {
	return &pubsub{
		client: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort),
			Username: cfg.RedisUsername,
			Password: cfg.RedisPassword,
			DB:       cfg.RedisDB,
		}),
	}
}

func (p *pubsub) Publish(ctx context.Context, msg *Message) error {
	return p.client.Publish(ctx, msg.Topic, msg.Body).Err()
}

func (p *pubsub) Subscribe(ctx context.Context, topic string, handler Handler) error {
	subscribe := p.client.Subscribe(ctx, topic)
	go func() {
		<-ctx.Done()
		subscribe.Close()
	}()

	channel := subscribe.Channel()
	for msg := range channel {
		if err := handler(&Message{
			Topic: topic,
			Body:  []byte(msg.Payload),
		}); err != nil {
			return err
		}
	}

	return nil
}
