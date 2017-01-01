package main

import (
	"github.com/andreiavrammsd/go-rabbitmq"
	"encoding/json"
	"fmt"
	"log"
)

type Email struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Message string `json:"message"`
}

var (
	conn *rabbitmq.Connection
	queueName string = "emails"
	config = &rabbitmq.Config{
		Scheme: "amqp",
		Host: "127.0.0.1",
		Port: 5672,
		Username: "guest",
		Password: "guest",
		Vhost: "/",
	}
	consumer rabbitmq.Consumer
	err error
)

func main() {
	// Establish connection to Rabbit server
	conn, err = rabbitmq.NewConnection(config)
	checkError(err)

	// Open a channel
	ch, err := conn.GetChannel()
	checkError(err)

	// Declare a queue
	q, err := ch.GetQueue(queueName)
	checkError(err)

	// We'll send a json string to queue
	email := Email{
		From: "myaddress@myprovider.tld",
		To: "youraddress@yourprovider.tld",
		Message: "Hello, GO!",
	}

	message, err := json.Marshal(email)
	checkError(err)

	// Send message to queue
	q.Publish(message)

	// Define a consumer function where we'll get all the messages from the queue
	consumer = func(delivery rabbitmq.Delivery) {
		for message := range delivery.Messages {
			email = Email{}
			err := json.Unmarshal(message.Body, &email)
			checkError(err)
			fmt.Println(email)

			// If everything ok, acknowledge message to server
			message.Ack(false)
		}
	}

	// Run one or more consumers on different threads
	go consume()
	go consume()

	// Run forever
	select {}
}

func consume() {
	// We keep a single connection to rabbit server, but we open a channel for
	// each thread we're consuming the queue from
	ch, err := conn.GetChannel()
	checkError(err)

	// Declare the queue on the new channel
	q, err := ch.GetQueue(queueName)
	checkError(err)

	// Pass our consumer function to the queue
	q.Consume(consumer)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
