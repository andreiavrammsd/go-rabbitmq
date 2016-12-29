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
