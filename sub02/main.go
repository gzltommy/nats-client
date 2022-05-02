package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"runtime"
	"time"
)

const natsURL = "nats://192.168.24.147:4222"

func main() {
	// create server connection
	// nats.DefaultURL
	natsConnection, _ := nats.Connect(natsURL)
	log.Println("connected to " + natsURL)

	// subscribe to subject
	log.Printf("subscribing to subject 'foo' \n")

	//natsConnection.Subscribe("foo", func(msg *nats.Msg) {
	//	//handle the message
	//	log.Printf("received message '%s\n", string(msg.Data)+"'")
	//})

	sub, err := natsConnection.SubscribeSync("foo")
	if err != nil {
		return
	}
	msg, err := sub.NextMsg(time.Second * 5)
	if err != nil {
		return
	}
	//handle the message
	log.Printf("received message '%s\n", string(msg.Data)+"'")

	// keep the connection alive
	runtime.Goexit()
}
