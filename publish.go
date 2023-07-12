package pubsub

import "context"

// Publish publishes a message to a topic.
func (p *pubsub) Publish(ctx context.Context, msg *Message) error {
	return p.client.Publish(ctx, msg.Topic, msg.Body).Err()
}
