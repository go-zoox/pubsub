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

// Message is a pubsub message.
type Message struct {
	Topic string
	Body  []byte
}

// Handler is a pubsub handler.
type Handler func(msg *Message) error

type pubsub struct {
	client *redis.Client
}

// Config is the config for a pubsub.
type Config struct {
	RedisHost     string
	RedisPort     int
	RedisUsername string
	RedisPassword string
	RedisDB       int
}

// New creates a new pubsub.
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
