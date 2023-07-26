package main

import (
	"context"
	"github.com/rabbitmq/amqp091-go"
	"rabbit/utils"
	"time"
)

func main() {

	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672")
	utils.CheckError(err, "Connection To Rabbit Failed")
	defer conn.Close()

	ch, err := conn.Channel()
	utils.CheckError(err, "Channel Creation Failed")
	defer ch.Close()

	err = ch.ExchangeDeclare("MyExchange", "topic", true, false, false, false, nil)
	utils.CheckError(err, "Exchange Declaration Failed")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	inp := utils.GetInput()
	err = ch.PublishWithContext(ctx, "MyExchange", "important.*", false, false, amqp091.Publishing{
		DeliveryMode: amqp091.Persistent,
		ContentType:  "text/plain",
		Body:         []byte(inp),
	})
}
