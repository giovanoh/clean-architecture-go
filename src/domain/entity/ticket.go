package entity

import (
	ts "github.com/giovanoh/clean-architecture-go/src/application/util"
	"github.com/google/uuid"
)

type Ticket struct {
	Id     string
	Name   string
	Email  string
	Price  float64
	Status ts.TicketStatus
}

func NewTicket(name string, email string, price float64) *Ticket {
	return &Ticket{
		Id:     uuid.New().String(),
		Name:   name,
		Email:  email,
		Price:  price,
		Status: ts.TicketStatusReserved,
	}
}

func (t *Ticket) Approve() {
	t.Status = ts.TicketStatusApproved
}

func (t *Ticket) Reject() {
	t.Status = ts.TicketStatusRejected
}
