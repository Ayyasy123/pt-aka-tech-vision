package main

import (
	"log"

	"github.com/streadway/amqp"
)

func receiveMessage() {
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
		"my-queue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Gagal mendeklarasikan queue: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,  // Auto-acknowledge
		false, // Exclusive
		false, // No-local
		false, // No-wait
		nil,
	)
	if err != nil {
		log.Fatalf("Gagal mendaftarkan consumer: %v", err)
	}

	log.Println("Menunggu pesan...")

	for msg := range msgs {
		log.Printf("Pesan diterima: %s", msg.Body)
	}
}

func main() {
	receiveMessage()
}
