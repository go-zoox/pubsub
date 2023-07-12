package pubsub

import "context"

func (p *pubsub) Subscribe(ctx context.Context, topic string, handler Handler) error {
	channel := p.client.Subscribe(ctx, topic).Channel()
	for msg := range channel {
		if err := handler(&Message{
			Topic: topic,
			Body:  []byte(msg.Payload),
		}); err != nil {
			return err
		}
	}

	return nil
}
