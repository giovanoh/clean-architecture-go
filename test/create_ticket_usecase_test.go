package test

import (
	"encoding/json"
	"testing"

	"github.com/giovanoh/clean-architecture-go/src/application/usecase"
	"github.com/giovanoh/clean-architecture-go/src/infra/gateway"
	"github.com/giovanoh/clean-architecture-go/src/infra/queue"
	"github.com/giovanoh/clean-architecture-go/src/infra/repository"
	"github.com/giovanoh/clean-architecture-go/src/infra/service"
)

func TestCreateTicketUseCase(t *testing.T) {
	memoryRepository := repository.NewTicketRepositoryMemory()
	memoryMailer := service.NewMemoryMailer()
	memoryPaymentGateway := gateway.NewPaymentGatewayPaypal()

	memoryQueue := queue.NewMemoryAdapter()
	memoryQueue.Connect()
	defer memoryQueue.Close()

	createTicket := usecase.NewCreateTicket(memoryRepository, memoryQueue)
	processPayment := usecase.NewProcessPayment(memoryRepository, memoryPaymentGateway, memoryQueue)
	notifyUser := usecase.NewNotifyUser(memoryRepository, memoryMailer)

	memoryQueue.On("ticket.created", func(message []byte) error {
		var ticketDto usecase.TicketCreatedEvent
		if err := json.Unmarshal(message, &ticketDto); err != nil {
			return err
		}
		input := usecase.ProcessPaymentInputDto(ticketDto)
		err := processPayment.Execute(input)
		return err
	})
	memoryQueue.On("payment.changed", func(message []byte) error {
		var input usecase.PaymentChangedEvent
		if err := json.Unmarshal(message, &input); err != nil {
			return err
		}

		ticket, err := memoryRepository.GetTicketById(input.TicketId)
		if err != nil {
			return err
		}
		switch input.Status {
		case "success":
			ticket.Approve()
		case "rejected":
			ticket.Reject()
		}
		if err = memoryRepository.Update(ticket); err != nil {
			return err
		}

		inputNotify := usecase.NotifyUserInputDto{
			TicketId: ticket.Id,
			Status:   ticket.Status,
		}
		err = notifyUser.Execute(inputNotify)
		return err
	})

	input := usecase.CreateTicketInputDto{
		Name:  "Ticket 1",
		Email: "buyer@ticket.com",
		Price: 100.0,
	}
	err := createTicket.Execute(input)
	if err != nil {
		t.Errorf("Error creating ticket: %v", err)
	}

	tickets, err := memoryRepository.GetTicketsByEmail("buyer@ticket.com")
	if err != nil {
		t.Errorf("Error getting ticket: %v", err)
	}
	if len(tickets) == 0 {
		t.Errorf("Expected ticket to be get, but got none")
	}
	ticket := tickets[0]
	if ticket == nil {
		t.Errorf("Expected ticket to be get, but got nil")
	}
}
