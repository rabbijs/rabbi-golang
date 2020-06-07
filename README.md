
# Rabbi Golang

Port of Node.js Rabbi Actor interface for Golang

## Usage

```

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

  actor.Start(func (_ amqp.Channel, msg amqp.Delivery) {

    log.Printf("message received: %s", string(msg.Body))

    msg.Ack(false)

  })

}

```

## About Rabbi

Rabbi is a tool to build scalable, real-time microservices for multiple processes on multiple hosts. Rabbi Actors
receive messages through AMQP messaging system and may send messages to any other actor on any host. Messages are 
delivered based on bindings on an exchange from a routing key to a queue. Actors by default process messages in their
queue one at a time and do not have access to the internal state of other ACtors.


