package rabbitmq

import "github.com/streadway/amqp"

type Channel struct {
	Channel *amqp.Channel
}

func (q *Connection) GetChannel() (*Channel, error) {
	ch, err := q.Connection.Channel()
	c := &Channel{
		Channel: ch,
	}

	return c, err
}
