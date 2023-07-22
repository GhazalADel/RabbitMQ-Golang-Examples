package main

import (
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"rabbit/utils"
	"time"
)

func main() {
	/*
	  I want to avoid doing a resource-intensive task immediately and having to wait for it to complete.
	  Instead we schedule the task to be done later.
	*/
	//I want to
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
