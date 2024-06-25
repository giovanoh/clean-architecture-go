package queue

type MemoryQueueAdapter struct {
	queue map[string][]func([]byte) error
}

func NewMemoryAdapter() *MemoryQueueAdapter {
	return &MemoryQueueAdapter{
		queue: make(map[string][]func([]byte) error),
	}
}

func (m *MemoryQueueAdapter) Connect() error {
	return nil
}

func (m *MemoryQueueAdapter) Close() error {
	return nil
}

func (m *MemoryQueueAdapter) On(queueName string, callback func([]byte) error) error {
	if _, ok := m.queue[queueName]; !ok {
		m.queue[queueName] = make([]func([]byte) error, 0)
	}
	m.queue[queueName] = append(m.queue[queueName], callback)
	return nil
}

func (m *MemoryQueueAdapter) Publish(queueName string, message []byte) error {
	if _, ok := m.queue[queueName]; !ok {
		return nil
	}
	for _, callback := range m.queue[queueName] {
		callback(message)
	}
	return nil
}
