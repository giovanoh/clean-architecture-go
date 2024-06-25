package repository

import "github.com/giovanoh/clean-architecture-go/src/domain/entity"

type TicketRepositoryMemory struct {
	Tickets []*entity.Ticket
}

func NewTicketRepositoryMemory() *TicketRepositoryMemory {
	return &TicketRepositoryMemory{
		Tickets: make([]*entity.Ticket, 0),
	}
}

func (t *TicketRepositoryMemory) Create(ticket *entity.Ticket) error {
	t.Tickets = append(t.Tickets, ticket)
	return nil
}

func (t *TicketRepositoryMemory) Update(ticket *entity.Ticket) error {
	for _, t := range t.Tickets {
		if t.Id == ticket.Id {
			t.Status = ticket.Status
		}
	}
	return nil
}

func (t *TicketRepositoryMemory) GetTicketById(id string) (*entity.Ticket, error) {
	for _, ticket := range t.Tickets {
		if ticket.Id == id {
			return ticket, nil
		}
	}
	return nil, nil
}

func (t *TicketRepositoryMemory) GetTicketsByEmail(email string) ([]*entity.Ticket, error) {
	tickets := []*entity.Ticket{}
	for _, ticket := range t.Tickets {
		if ticket.Email == email {
			tickets = append(tickets, ticket)
		}
	}
	return tickets, nil
}
