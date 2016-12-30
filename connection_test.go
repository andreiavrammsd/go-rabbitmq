package rabbitmq

import (
	"testing"
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
)

func TestNewConnectionSuccess(t *testing.T) {
	dial = func(url string) (*amqp.Connection, error) {
		return &amqp.Connection{}, nil
	}

	config := &Config{
		Scheme: "amqp",
		Host: "127.0.0.1",
		Port: 5672,
		Username: "guest",
		Password: "guest",
		Vhost: "/",
	}
	conn, err := NewConnection(config)

	expectedConnection := &Connection{
		Config: config,
		Connection: &amqp.Connection{},
	}
	assert.Equal(t, expectedConnection, conn)
	assert.Nil(t, err)
}

func TestNewConnectionFail(t *testing.T) {
	dial = func(url string) (*amqp.Connection, error) {
		return nil, nil
	}

	config := &Config{
		Scheme: "invalid scheme",
		Host: "127.0.0.1",
		Port: 5672,
		Username: "guest",
		Password: "guest",
		Vhost: "/",
	}
	conn, err := NewConnection(config)

	expectedConnection := &Connection{
		Config: config,
		Connection: nil,
	}
	assert.Equal(t, expectedConnection, conn)
	assert.Nil(t, err)
}
