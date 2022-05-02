package main

import (
	"github.com/nats-io/nats.go"
	"log"
)

const natsURL = "nats://192.168.24.147:4222"

func main() {
	// create server connection and defer close
	natsConnection, _ := nats.Connect(natsURL)
	defer natsConnection.Close()
	log.Println("connected to " + natsURL)

	// msg structure
	msg := &nats.Msg{Subject: "foo", Reply: "bar", Data: []byte("Hello World")}
	natsConnection.PublishMsg(msg)

	log.Println("published msg.Subject = "+msg.Subject, " | msg.Data = "+string(msg.Data))
}
