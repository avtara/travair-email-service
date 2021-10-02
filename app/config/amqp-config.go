package config

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
)

func CreateConnectionAMQP() *amqp.Channel {
	amqpURI := os.Getenv("AMQP_URI")
	conn, err := amqp.Dial(amqpURI)
	if err != nil {
		fmt.Println("Failed Initializing Broker Connection")
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	return ch
}
