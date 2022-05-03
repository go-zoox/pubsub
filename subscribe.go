package pubsub

// Subscribe subscribes to a topic.
// returns a function to unsubscribe.
func Subscribe(ps *PubSub, topic string, subscriber Subscriber) func() {
	return ps.Subscribe(topic, subscriber)
}
