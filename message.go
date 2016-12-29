package rabbitmq

import "github.com/streadway/amqp"

type MessageConfig struct {
	Body         []byte
	ContentType  string
	DeliveryMode uint8
	Exchange     string
	RoutingKey   string
	Mandatory    bool
	Immediate    bool
}

var Message = &MessageConfig{
	ContentType: "text/plain",
	DeliveryMode: amqp.Persistent,
	Exchange: "",
	RoutingKey: "",
	Mandatory: false,
	Immediate: false,
}
