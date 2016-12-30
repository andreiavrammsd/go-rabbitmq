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
