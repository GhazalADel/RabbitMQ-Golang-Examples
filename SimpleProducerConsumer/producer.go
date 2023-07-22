package main

import (
	"context"
	"fmt"
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

	q, err := ch.QueueDeclare("Q", true, false, false, false, nil)
	utils.CheckError(err, "Queue Declaration Failed")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	ch.PublishWithContext(ctx, "", q.Name, false, false, amqp091.Publishing{
		Body:        []byte("Simple Producer"),
		ContentType: "text/plain",
	})
	fmt.Printf("%+v\n", q)
}
