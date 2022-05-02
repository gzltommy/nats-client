package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

const natsURL = "nats://foo:bar@192.168.24.147:4222"

func main() {
	// create server connection and defer close
	// nats.DefaultURL
	natsConnection, _ := nats.Connect(natsURL)
	defer natsConnection.Close()
	log.Println("connected to " + natsURL)

	// publish messge on subject by name foo
	subject := "foo"
	natsConnection.Publish(subject, []byte("Hello NATS"))
	log.Printf("published message on subject " + subject)

	time.Sleep(30 * time.Second)
}
