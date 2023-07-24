# RabbitMQ-Golang-Examples

This repository contains Golang code samples demonstrating the usage of RabbitMQ based on the official tutorial.

## Introduction

RabbitMQ is a popular message broker that facilitates communication between different applications or services. It is widely used in distributed systems, microservices architectures, and other scenarios where decoupling and asynchronous messaging are required.

This repository contains Golang code examples and tutorials to help you get started with RabbitMQ and understand its concepts in a practical way.

## Prerequisites

Before getting started with RabbitPlayground, make sure you have the following prerequisites installed:

1. Go Programming Language: Make sure you have Go installed on your system. You can download and install it from the official website: https://golang.org/

2. Docker: RabbitMQ will be run using Docker, so ensure Docker is installed on your machine. You can download Docker from: https://www.docker.com/

## Installation

### RabbitMQ Server

To install RabbitMQ, pull the Docker image with the following command:

```bash
docker pull rabbitmq:3-management
```

### RabbitMQ Library

Install the RabbitMQ library for Golang using the following command:

```bash
docker pull rabbitmq:3-management
```

## Running RabbitMQ Server

To run the RabbitMQ container, execute the following command:

```bash
docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management
```

After running the container, you can access the RabbitMQ management dashboard at localhost:15672. The default username and password are `guest`.


## Examples

The `examples/` directory contains various Golang code examples that demonstrate different aspects of using RabbitMQ. To run the example programs, navigate to the specific example directory and execute the Go files.

For example, to execute the first program that demonstrates a simple producer-consumer setup:

```bash
cd SimpleProducerConsumer/
go run producer.go
go run consumer.go
```

Feel free to explore other examples to understand how RabbitMQ can be used for different messaging patterns and scenarios.
