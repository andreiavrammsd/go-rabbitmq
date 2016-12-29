package rabbitmq

import "github.com/streadway/amqp"

type ConsumerConfig struct {
	QueueName string
	Consumer  string
	AutoAck   bool
	Exclusive bool
	NoLocal   bool
	NoWait    bool
	Args      amqp.Table
}

type Consumer func(delivery Delivery)

type Delivery struct {
	Messages <-chan amqp.Delivery
	Queue    amqp.Queue
}

var ConsumerConfiguration = &ConsumerConfig{
	Consumer: "",
	AutoAck : false,
	Exclusive: false,
	NoLocal: false,
	NoWait: false,
	Args : nil,
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

