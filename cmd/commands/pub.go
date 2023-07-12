package commands

import (
	"context"

	"github.com/go-zoox/cli"
	"github.com/go-zoox/pubsub"
)

// Pub is the command for publishing a message to a topic.
func Pub(app *cli.MultipleProgram) {
	app.Register("pub", &cli.Command{
		Name:  "pub",
		Usage: "the pub of pubsub",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "topic",
				Usage:   "the topic to publish",
				EnvVars: []string{"TOPIC"},
				Value:   "default",
			},
			&cli.StringFlag{
				Name:     "message",
				Usage:    "the message to publish",
				EnvVars:  []string{"MESSAGE"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "redis-host",
				Usage:    "the redis host",
				EnvVars:  []string{"REDIS_HOST"},
				Required: true,
			},
			&cli.IntFlag{
				Name:    "redis-port",
				Usage:   "the redis port",
				EnvVars: []string{"REDIS_PORT"},
				Value:   6379,
			},
			&cli.StringFlag{
				Name:    "redis-username",
				Usage:   "the redis username",
				EnvVars: []string{"REDIS_USERNAME"},
			},
			&cli.StringFlag{
				Name:    "redis-password",
				Usage:   "the redis password",
				EnvVars: []string{"REDIS_PASSWORD"},
			},
			&cli.IntFlag{
				Name:    "redis-db",
				Usage:   "the redis db",
				EnvVars: []string{"REDIS_DB"},
				Value:   0,
			},
		},
		Action: func(ctx *cli.Context) error {
			ps := pubsub.New(&pubsub.Config{
				RedisHost:     ctx.String("redis-host"),
				RedisPort:     ctx.Int("redis-port"),
				RedisUsername: ctx.String("redis-username"),
				RedisPassword: ctx.String("redis-password"),
				RedisDB:       ctx.Int("redis-db"),
			})

			return ps.Publish(context.TODO(), &pubsub.Message{
				Topic: ctx.String("topic"),
				Body:  []byte(ctx.String("message")),
			})
		},
	})
}
