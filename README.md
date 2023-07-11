# PubSub - lightweight pub/sub messaging

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/pubsub)](https://pkg.go.dev/github.com/go-zoox/pubsub)
[![Build Status](https://github.com/go-zoox/pubsub/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/pubsub/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/pubsub)](https://goreportcard.com/report/github.com/go-zoox/pubsub)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/pubsub/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/pubsub?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/pubsub.svg)](https://github.com/go-zoox/pubsub/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/pubsub.svg?label=Release)](https://github.com/go-zoox/pubsub/tags)

## Installation
To install the package, run:
```bash
go get github.com/go-zoox/pubsub
```

## Getting Started


### Subscribe
```go
import (
  "github.com/go-zoox/pubsub"
)

func main(t *testing.T) {
	ps := pubsub.New(&pubsub.Config{
		RedisHost:     <RedisHost>,
		RedisPort:     <RedisPort>,
		RedisUsername: <RedisUsername>,
		RedisPassword: <RedisPassword>,
		RedisDB:       <RedisDB>,
	})

	ps.Subscribe(context.TODO(), "default", func(msg *pubsub.Message) error {
				logger.Infof("received message: %s", string(msg.Body))
				return nil
			})
}
```

### Publish
```go
import (
  "github.com/go-zoox/pubsub"
)

func main(t *testing.T) {
	ps := pubsub.New(&pubsub.Config{
		RedisHost:     <RedisHost>,
		RedisPort:     <RedisPort>,
		RedisUsername: <RedisUsername>,
		RedisPassword: <RedisPassword>,
		RedisDB:       <RedisDB>,
	})

	ps.Publish(context.TODO(), &pubsub.Message{
				Topic: "default",
				Body:  []byte("hello world"),
			})
}
```

## License
GoZoox is released under the [MIT License](./LICENSE).
