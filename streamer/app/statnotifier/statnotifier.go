package statnotifier

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func getRabbitQueue(ch *amqp.Channel) *amqp.Queue {
	q, err := ch.QueueDeclare(
		"statistic_queue", // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a queue")

	return &q
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


type StatNotifier struct {
	Channel *amqp.Channel
	Queue *amqp.Queue
	Context context.Context
	CancelContext context.CancelFunc
}

func InitStatNotifier() *StatNotifier {
	var s StatNotifier

	conn := getRabbitConnection()
	ch := getRabbitChannel(conn)
	q := getRabbitQueue(ch)

	s.Channel = ch
	s.Queue = q
	s.Context, s.CancelContext = context.WithTimeout(context.Background(), 5*time.Second)

	return &s
}

func (s *StatNotifier) Publish(songName string) {
	err := s.Channel.PublishWithContext(s.Context,
		"",         // exchange
		s.Queue.Name, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(songName),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent data of %s\n", songName)
	s.CancelContext()
}