package queue

import (
	"context"
	"time"

	rabbitmq "github.com/rabbitmq/amqp091-go"
)

type RabbitMQAdapter struct {
	conn    *rabbitmq.Connection
	channel *rabbitmq.Channel
	queue   rabbitmq.Queue
	uri     string
}

func NewRabbitMQAdapter(uri string) *RabbitMQAdapter {
	return &RabbitMQAdapter{
		uri: uri,
	}
}

func (r *RabbitMQAdapter) Connect() error {
	var err error
	r.conn, err = rabbitmq.Dial(r.uri)
	if err != nil {
		return err
	}

	r.channel, err = r.conn.Channel()
	if err != nil {
		return err
	}

	return nil
}

func (r *RabbitMQAdapter) Close() error {
	r.channel.Close()
	return r.conn.Close()
}

func (r *RabbitMQAdapter) On(queueName string, callback func([]byte) error) error {
	msgs, err := r.channel.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			callback(d.Body)
		}
	}()

	return nil
}

func (r *RabbitMQAdapter) Publish(queueName string, message []byte) error {
	var err error
	r.queue, err = r.channel.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = r.channel.PublishWithContext(ctx,
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		rabbitmq.Publishing{
			ContentType: "text/json",
			Body:        message,
		})
	return err
}
