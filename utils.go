package pubsub

import (
	"github.com/go-zoox/uuid"
)

// SubscribeFunc creates a Handle from a function.
func SubscribeFunc(subscriber func(payload any)) Subscriber {
	return &subscriberFuncCreator{
		id: uuid.V4(),
		fn: subscriber,
	}
}

type subscriberFuncCreator struct {
	id string
	fn func(payload any)
}

// Serve calls the function.
func (h *subscriberFuncCreator) Serve(payload any) {
	h.fn(payload)
}

// ID returns the id of the handle.
func (h *subscriberFuncCreator) ID() string {
	return h.id
}
