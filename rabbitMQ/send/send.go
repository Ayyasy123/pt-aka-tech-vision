package main

import (
	"log"

	"github.com/streadway/amqp"
)

func sendMessage() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Gagal terhubung ke RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Gagal membuka channel: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"my-queue", // Nama queue
		false,      // Durable
		false,      // Delete when unused
		false,      // Exclusive
		false,      // No-wait
		nil,        // Arguments
	)
	if err != nil {
		log.Fatalf("Gagal mendeklarasikan queue: %v", err)
	}

	messageBody := "Hello, RabbitMQ!"
	err = ch.Publish(
		"",     // Exchange
		q.Name, // Routing key (queue name)
		false,  // Mandatory
		false,  // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(messageBody),
		},
	)
	if err != nil {
		log.Fatalf("Gagal mengirim pesan: %v", err)
	}

	log.Printf("Pesan terkirim: %s", messageBody)
}

func main() {
	sendMessage()
}
