package config

import "github.com/streadway/amqp"

func NewQueue() (*amqp.Connection, error) {
	return amqp.Dial("amqp://guest:guest@127.0.0.1:5672/")
}
