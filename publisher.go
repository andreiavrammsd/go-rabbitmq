package rabbitmq

import "github.com/streadway/amqp"

func (q *Queue) Publish(body []byte) error {
	err := q.Channel.Publish(
		Message.Exchange,
		q.Queue.Name,
		Message.Mandatory,
		Message.Immediate,
		amqp.Publishing{
			DeliveryMode: Message.DeliveryMode,
			ContentType: Message.ContentType,
			Body: body,
		},
	)

	return err
}
