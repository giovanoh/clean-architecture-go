package service

import "fmt"

type MemoryMailer struct {
}

func NewMemoryMailer() *MemoryMailer {
	return &MemoryMailer{}
}

func (m *MemoryMailer) Send(to string, subject string, message string) error {
	fmt.Printf("Sending email to %s with subject %s and message %s\n", to, subject, message)

	return nil
}
