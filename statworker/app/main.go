package main

import (
	"context"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func getRabbitQueue(ch *amqp.Channel) amqp.Queue {
	q, err := ch.QueueDeclare(
		"statistic_queue", // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a queue")

	return q
}

func getRabbitChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	return ch
}

func getRabbitConnection() *amqp.Connection {
	conn, err := amqp.Dial("amqp://rmuser:rmpassword@rabbitmq/")
	failOnError(err, "Failed to connect to RabbitMQ")
	return conn
}

func initRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "change-me",            
		DB:       0,              
	})
}

func main() {
	ctx := context.Background()
	redisClient := initRedisClient()
	conn := getRabbitConnection()
	ch := getRabbitChannel(conn)
	queue := getRabbitQueue(ch)

	msgs, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			redisClient.ZAddArgsIncr(ctx, "statistics", redis.ZAddArgs{Members: []redis.Z{{Score: 1, Member: string(d.Body)}}})

			log.Printf("Received a listening message of %s", d.Body)
			if err != nil {
				fmt.Println(err)
			}
		}
	}()

	<-forever
}