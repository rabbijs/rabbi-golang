
package rabbi

import (
    "log"
    "os"
    "github.com/joho/godotenv"
    "github.com/streadway/amqp"
)

type Actor struct {
  Queue string
  RoutingKey string
  Exchange string
}

func (a Actor) Start(consume func(ch *amqp.Channel, message amqp.Delivery)) {

  godotenv.Load()

  connection, err := amqp.Dial(os.Getenv("AMQP_URL"))

  if err != nil {
    panic(err);
  }

  defer connection.Close()

  log.Println("amqp server connected")

  channel, err := connection.Channel()

  log.Println("amqp channel connected")

  err = channel.ExchangeDeclare(a.Exchange, "direct", true, false, false, false, nil)

  err = channel.Qos(1, 0, false)

  _, err = channel.QueueDeclare(a.Queue, true, false, false, false, nil)

  if err != nil {
    panic(err);
  }

  err = channel.QueueBind(a.Queue, a.RoutingKey, a.Exchange, false, nil)

  log.Printf("amqp queue bound %s", a)

  consumer, err := channel.Consume(a.Queue, a.Queue, false, false, false, false, nil)

  if err != nil {
    panic(err);
  }

  for msg := range consumer {

    go consume(channel, msg)

  }

}

