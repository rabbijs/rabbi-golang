package main

import (
    "github.com/streadway/amqp"
    "github.com/rabbijs/rabbi-golang/rabbi"
    "log"
    "time"
)


func main() {

  // this actor will automatically republish the first message it receives, forever

  actor := rabbi.Actor {
    Queue: "golang_queue",
    RoutingKey: "golang_messages",
    Exchange: "rabbi_golang",
  }

  actor.Start(func (channel *amqp.Channel, msg amqp.Delivery) {

    log.Printf("message received: %s", string(msg.Body))

    msg.Ack(false)

    time.Sleep(time.Second)

    publishMessage("rabbi_golang", "golang_messages", channel, string(msg.Body))

  })

}

func publishMessage(exchange string, routingkey string, channel *amqp.Channel, message string) {

    newMsg := amqp.Publishing{
        DeliveryMode: amqp.Persistent,
        Timestamp:    time.Now(),
        ContentType:  "text/plain",
        Body:         []byte(message),
    }

    err := channel.Publish(exchange, routingkey, false, false, newMsg)

    if err != nil {
      panic(err)
    }

}

