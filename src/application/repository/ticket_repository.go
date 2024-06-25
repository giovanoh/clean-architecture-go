package repository

import "github.com/giovanoh/clean-architecture-go/src/domain/entity"

type TicketRepository interface {
	Create(ticket *entity.Ticket) error
	Update(ticket *entity.Ticket) error
	GetTicketById(id string) (*entity.Ticket, error)
	GetTicketsByEmail(email string) ([]*entity.Ticket, error)
}
