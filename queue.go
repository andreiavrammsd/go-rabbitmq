package rabbitmq

import "github.com/streadway/amqp"

type Queue struct {
	Name       string
	Channel    *amqp.Channel
	Queue      amqp.Queue
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
	Args       amqp.Table
}

func (ch *Channel) GetQueue(name string) (*Queue, error) {
	q := &Queue{
		Name: name,
		Channel: ch.Channel,
		Durable: true,
		AutoDelete: false,
		Exclusive: false,
		NoWait: false,
		Args : nil,
	}
	queue, err := ch.Channel.QueueDeclare(
		q.Name,
		q.Durable,
		q.AutoDelete,
		q.Exclusive,
		q.NoWait,
		q.Args,
	)
	q.Queue = queue

	return q, err
}

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

func (q *Queue) Consume(consumer Consumer) (error) {
	config := ConsumerConfiguration
	config.QueueName = q.Queue.Name

	messages, err := q.Channel.Consume(
		config.QueueName,
		config.Consumer,
		config.AutoAck,
		config.Exclusive,
		config.NoLocal,
		config.NoWait,
		config.Args,
	)

	if err != nil {
		return err
	}

	delivery := Delivery{
		Messages: messages,
		Queue: q.Queue,
	}
	consumer(delivery)

	return nil
}
