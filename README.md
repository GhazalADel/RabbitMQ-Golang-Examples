# RabbitPlayground

RabbitPlayground is a repository for Implementing RabbitMQ with Golang.

## Installation

Use this command to install RabbitMQ.

```bash
docker pull rabbitmq:3-management
```
Use this command to install RabbitMQ library.
```bash
go get github.com/rabbitmq/amqp091-go
```
## Installation
To Run the container, Enter this command:
```bash
docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management
```
Go to [localhost:15672](http://localhost:15672) . default username and password is guest.

## Run Program
For Example For executing first program:
```bash
cd SimpleProducerConsumer/
go run producer.go
go run consumer.go
```