package rabbitmq

import (
	"github.com/streadway/amqp"
	"fmt"
)

type Config struct {
	Address  string
	Username string
	Password string
}

type Connection struct {
	Config     *Config
	Connection *amqp.Connection
}

func NewConnection(config *Config) (*Connection, error) {
	q := &Connection{
		Config: config,
	}

	url := fmt.Sprintf("amqp://%s:%s@%s/", config.Username, config.Password, config.Address)
	conn, err := dial(url)

	q.Connection = conn

	return q, err
}

func dial(url string) (*amqp.Connection, error) {
	return amqp.Dial(url)
}
