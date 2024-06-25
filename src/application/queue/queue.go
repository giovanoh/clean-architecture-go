package queue

type Queue interface {
	Connect() error
	Close() error
	On(queueName string, callback func([]byte) error) error
	Publish(queueName string, message []byte) error
}
