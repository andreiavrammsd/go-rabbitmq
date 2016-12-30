package rabbitmq

import (
	"github.com/streadway/amqp"
)

type Config struct {
	Scheme   string
	Host     string
	Port     int
	Username string
	Password string
	Vhost    string
}

type Connection struct {
	Config     *Config
	Connection *amqp.Connection
}

type Channel struct {
	Channel *amqp.Channel
}

var dial = amqp.Dial

func NewConnection(config *Config) (*Connection, error) {
	uri := amqp.URI{
		Scheme: config.Scheme,
		Host: config.Host,
		Port: config.Port,
		Username: config.Username,
		Password: config.Password,
		Vhost: config.Vhost,
	}
	conn, err := dial(uri.String())

	q := &Connection{
		Config: config,
		Connection: conn,
	}

	return q, err
}

func (c *Connection) GetChannel() (*Channel, error) {
	newChannel, err := c.Connection.Channel()
	ch := &Channel{
		Channel: newChannel,
	}

	return ch, err
}
