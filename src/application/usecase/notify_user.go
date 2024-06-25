package usecase

import (
	"fmt"

	"github.com/giovanoh/clean-architecture-go/src/application/repository"
	"github.com/giovanoh/clean-architecture-go/src/application/service"
	ts "github.com/giovanoh/clean-architecture-go/src/application/util"
)

type NotifyUserInputDto struct {
	TicketId string          `json:"ticket_id"`
	Status   ts.TicketStatus `json:"status"`
}

type NotifyUserUseCase struct {
	repository repository.TicketRepository
	mailer     service.Mailer
}

func NewNotifyUser(ticketRepository repository.TicketRepository, mailer service.Mailer) *NotifyUserUseCase {
	return &NotifyUserUseCase{
		repository: ticketRepository,
		mailer:     mailer,
	}
}

func (u *NotifyUserUseCase) Execute(input NotifyUserInputDto) error {
	fmt.Println("notify_user_use_case.execute: ", input)

	ticket, err := u.repository.GetTicketById(input.TicketId)
	if err != nil {
		return err
	}

	if ticket.Status == "approved" {
		err = u.mailer.Send(ticket.Email, "Ticket purchase", "Congratulations on your purchase!")
	}
	if ticket.Status == "rejected" {
		err = u.mailer.Send(ticket.Email, "Ticket purchase", "Sorry, your purchase was not approved.")
	}

	return err
}
