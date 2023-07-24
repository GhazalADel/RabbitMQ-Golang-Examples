package main

import (
	"context"
	"github.com/rabbitmq/amqp091-go"
	"rabbit/utils"
	"time"
)

func main() {
	/*
	  Used to Sending messages to many consumers at once
	*/
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672")
	utils.CheckError(err, "Connection To Rabbit Failed")
	defer conn.Close()

	ch, err := conn.Channel()
	utils.CheckError(err, "Channel Creation Failed")
	defer ch.Close()

	err = ch.ExchangeDeclare("MyExchange", "fanout", true, false, false, false, nil)
	utils.CheckError(err, "Exchange Declaration Failed")

	q, err := ch.QueueDeclare("", true, false, false, false, nil)
	utils.CheckError(err, "Queue Declaration Failed")

	ch.QueueBind(q.Name, "", "MyExchange", false, nil)
	utils.CheckError(err, "Bind Exchange to Queue Failed")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	inp := utils.GetInput()
	err = ch.PublishWithContext(ctx, "MyExchange", "", false, false, amqp091.Publishing{
		DeliveryMode: amqp091.Persistent,
		ContentType:  "text/plain",
		Body:         []byte(inp),
	})
}
