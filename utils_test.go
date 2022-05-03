package pubsub

import (
	"fmt"
	"testing"
)

func TestSubscriber(t *testing.T) {
	h1 := SubscribeFunc(func(payload any) {
		t.Log("h1", payload)
	})

	h2 := SubscribeFunc(func(payload any) {
		t.Log("h2", payload)
	})

	h3 := SubscribeFunc(func(payload any) {
		t.Log("h3", payload)
	})

	if h1.ID() == h2.ID() {
		t.Error("h1 and h2 should have different ids")
	}

	if h1.ID() == h3.ID() {
		t.Error("h1 and h3 should have different ids")
	}

	if h2.ID() == h3.ID() {
		t.Error("h2 and h3 should have different ids")
	}

	fmt.Println(h1.ID(), h2.ID(), h3.ID())
}
