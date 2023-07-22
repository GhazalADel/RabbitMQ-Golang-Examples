package main

import (
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"rabbit/utils"
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

	done := make(chan bool)
	go func() {
		msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
		utils.CheckError(err, "Failed to Receive Messages")
		for msg := range msgs {
			fmt.Println(string(msg.Body))
			//To Delete From Queue
			msg.Ack(false)
		}
	}()
	<-done
}
