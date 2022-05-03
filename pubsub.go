package pubsub

import (
	"sync"
)

// PubSub is a simple pubsub.
type PubSub struct {
	ch          chan *Action
	subscribers map[string][]Subscriber
	m           sync.Mutex
}

// Subscriber is a function that can be registered to an event.
type Subscriber interface {
	ID() string
	Serve(payload any)
}

// Action is an action for a PubSub.
type Action struct {
	Topic   string
	Payload any
}

// New creates a new PubSub.
func New() *PubSub {
	return &PubSub{
		subscribers: make(map[string][]Subscriber),
	}
}

// Subscribe subscribes to a topic.
// returns a function to unsubscribe.
func (e *PubSub) Subscribe(topic string, subscriber Subscriber) func() {
	e.m.Lock()
	defer e.m.Unlock()

	e.subscribers[topic] = append(e.subscribers[topic], subscriber)

	return func() {
		e.UnSubscribe(topic, subscriber)
	}
}

// Publish trigger subscribers for a topic
func (e *PubSub) Publish(topic string, payload any) {
	if e.ch == nil {
		panic("pubsub is not started or stopped	")
	}

	e.ch <- &Action{
		Topic:   topic,
		Payload: payload,
	}
}

// SubscribeOnce performs exactly one action.
func (e *PubSub) SubscribeOnce(topic string, subscriber Subscriber) {
	var once sync.Once
	e.Subscribe(topic, SubscribeFunc(func(payload any) {
		once.Do(func() {
			subscriber.Serve(payload)
		})
	}))
}

// UnSubscribe unsubscribes a subscriber from a topic.
func (e *PubSub) UnSubscribe(typ string, subscriber Subscriber) {
	e.m.Lock()
	defer e.m.Unlock()

	for i, s := range e.subscribers[typ] {
		if s.ID() == subscriber.ID() {
			e.subscribers[typ] = append(e.subscribers[typ][:i], e.subscribers[typ][i+1:]...)
			break
		}
	}
}

func (e *PubSub) consume(action *Action) {
	e.m.Lock()
	defer e.m.Unlock()

	for _, subscriber := range e.subscribers[action.Topic] {
		subscriber.Serve(action.Payload)
	}
}

// Start starts the event worker.
func (e *PubSub) Start() {
	e.ch = make(chan *Action)

	go func() {
		for {
			select {
			case action := <-e.ch:
				e.consume(action)
			}
		}
	}()
}

// Stop stops the event worker.
func (e *PubSub) Stop() {
	close(e.ch)
	e.ch = nil
}
