package main

import (
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

	err = ch.ExchangeDeclare("MyExchange", "fanout", true, false, false, false, nil)
	utils.CheckError(err, "Exchange Declaration Failed")

	q, err := ch.QueueDeclare("", true, false, false, false, nil)
	utils.CheckError(err, "Queue Declaration Failed")

	err = ch.QueueBind(q.Name, "log", "MyExchange", false, nil)
	utils.CheckError(err, "Bind Exchange to Queue Failed")

	done := make(chan bool)
	go func() {
		msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
		utils.CheckError(err, "Failed to Receive Messages")
		for msg := range msgs {
			fmt.Println(string(msg.Body) + " START")
			starCount := utils.StarCount(msg.Body)
			t := time.Duration(starCount)
			time.Sleep(t * time.Second)
			fmt.Println(string(msg.Body) + " DONE")
			//To Delete From Queue
			msg.Ack(false)
		}
	}()
	<-done
}
