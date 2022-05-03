package pubsub

// Publish trigger subscribers for a topic
func Publish(ps *PubSub, topic string, payload any) {
	ps.Publish(topic, payload)
}
