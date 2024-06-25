package usecase

import (
	"encoding/json"
	"fmt"

	"github.com/giovanoh/clean-architecture-go/src/application/queue"
	"github.com/giovanoh/clean-architecture-go/src/application/repository"
	"github.com/giovanoh/clean-architecture-go/src/domain/entity"
)

type CreateTicketInputDto struct {
	Name            string  `json:"name"`
	Email           string  `json:"email"`
	Price           float64 `json:"price"`
	CreditCardToken string  `json:"credit_card_token"`
}

type TicketCreatedEvent struct {
	TicketId        string `json:"ticket_id"`
	CreditCardToken string `json:"credit_card_token"`
}

type CreateTicketUseCase struct {
	TicketRepository repository.TicketRepository
	Queue            queue.Queue
}

func NewCreateTicket(ticketRepository repository.TicketRepository, queue queue.Queue) *CreateTicketUseCase {
	return &CreateTicketUseCase{
		TicketRepository: ticketRepository,
		Queue:            queue,
	}
}

func (u *CreateTicketUseCase) Execute(input CreateTicketInputDto) error {
	fmt.Println("create_ticket_use_case.execute: ", input)

	ticket := entity.NewTicket(input.Name, input.Email, input.Price)
	err := u.TicketRepository.Create(ticket)
	if err != nil {
		return err
	}

	ticketCreatedEvent := &TicketCreatedEvent{
		TicketId:        ticket.Id,
		CreditCardToken: input.CreditCardToken,
	}
	event, err := json.Marshal(ticketCreatedEvent)
	if err != nil {
		return err
	}
	return u.Queue.Publish("ticket.created", event)
	// Microservices na pratica: 1:51:11
	// Aqui deveria spawnar ticket created? Parece que sim. SIM, pq o process payment também faz
}
