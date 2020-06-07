package main

import (
    "github.com/streadway/amqp"
    "github.com/rabbijs/rabbi-golang/rabbi"
    "log"
)


func main() {

  actor := rabbi.Actor {
    Queue: "golang_queue",
    RoutingKey: "golang_messages",
    Exchange: "rabbi_golang",
  }

  actor.Start(func (msg amqp.Delivery) {

    log.Printf("message received: %s", string(msg.Body))

    msg.Ack(false)

  })

}

