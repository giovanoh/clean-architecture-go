package service

type Mailer interface {
	Send(to string, subject string, message string) error
}
