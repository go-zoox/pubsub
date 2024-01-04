package pubsub

import "context"

// Subcribe subscribes to a topic.
func (p *pubsub) Subscribe(ctx context.Context, topic string, handler Handler) error {
	subscribe := p.client.Subscribe(ctx, topic)
	go func() {
		<-ctx.Done()
		subscribe.Close()
	}()

	channel := subscribe.Channel()
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
