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

```go
import (
  "testing"
  "github.com/go-zoox/pubsub"
)

func main(t *testing.T) {
	e := pubsub.New()
	count := 0
	e.On("send.notify", pubsub.HandleFunc(func(payload any) {
		count++
		t.Log("send.notify", payload)
	}))

	e.Start()

	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		index := i
		wg.Add(1)
		go func() {
			e.Emit("send.notify", index)
			wg.Done()
		}()
	}

	wg.Wait()
}
```

## License
GoZoox is released under the [MIT License](./LICENSE).
