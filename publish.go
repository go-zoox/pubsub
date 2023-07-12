package pubsub

import "context"

func (p *pubsub) Publish(ctx context.Context, msg *Message) error {
	return p.client.Publish(ctx, msg.Topic, msg.Body).Err()
}
