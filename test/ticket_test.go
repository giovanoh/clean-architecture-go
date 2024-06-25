package test

import (
	"testing"

	"github.com/giovanoh/clean-architecture-go/src/domain/entity"
)

func TestTicket(t *testing.T) {

	ticket := entity.NewTicket("Ticket 1", "buyer@ticket.com", 100.0)

	if len(ticket.Id) == 0 {
		t.Errorf("Expected ticket id to be not empty")
	}
	if ticket.Name != "Ticket 1" {
		t.Errorf("Expected ticket name to be 'Ticket 1', but got %s", ticket.Name)
	}
	if ticket.Email != "buyer@ticket.com" {
		t.Errorf("Expected ticket email to be 'buyer@ticket.com', but got %s", ticket.Email)
	}
	if ticket.Price != 100.0 {
		t.Errorf("Expected ticket price to be 100.0, but got %f", ticket.Price)
	}
	if ticket.Status != "reserved" {
		t.Errorf("Expected ticket status to be 'reserved', but got %s", ticket.Status)
	}
}
