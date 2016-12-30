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
	q := &Connection{
		Config: config,
	}

	uri := amqp.URI{
		Scheme: config.Scheme,
		Host: config.Host,
		Port: config.Port,
		Username: config.Username,
		Password: config.Password,
		Vhost: config.Vhost,
	}

	conn, err := dial(uri.String())

	q.Connection = conn

	return q, err
}

func (q *Connection) GetChannel() (*Channel, error) {
	ch, err := q.Connection.Channel()
	c := &Channel{
		Channel: ch,
	}

	return c, err
}
