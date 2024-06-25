package usecase

import (
	"encoding/json"
	"fmt"

	"github.com/giovanoh/clean-architecture-go/src/application/gateway"
	"github.com/giovanoh/clean-architecture-go/src/application/queue"
	"github.com/giovanoh/clean-architecture-go/src/application/repository"
)

type ProcessPaymentInputDto struct {
	TicketId        string `json:"ticket_id"`
	CreditCardToken string `json:"credit_card_token"`
}

type PaymentChangedEvent struct {
	TicketId string `json:"ticket_id"`
	Status   string `json:"status"`
}

type ProcessPaymentUseCase struct {
	queue      queue.Queue
	repository repository.TicketRepository
	gateway    gateway.PaymentGateway
}

func NewProcessPayment(ticketRepository repository.TicketRepository, paymentGateway gateway.PaymentGateway, queue queue.Queue) *ProcessPaymentUseCase {
	return &ProcessPaymentUseCase{
		queue:      queue,
		repository: ticketRepository,
		gateway:    paymentGateway,
	}
}

func (u *ProcessPaymentUseCase) Execute(input ProcessPaymentInputDto) error {
	fmt.Println("process_payment_use_case.execute: ", input)

	ticket, err := u.repository.GetTicketById(input.TicketId)
	if err != nil {
		return err
	}

	inputPayment := gateway.PaymentGatewayInputDto{
		TicketId:        ticket.Id,
		Price:           ticket.Price,
		CreditCardToken: input.CreditCardToken,
	}
	output, erro := u.gateway.ProcessPayment(inputPayment)
	if erro != nil {
		return erro
	}

	paymentDto := &PaymentChangedEvent{
		TicketId: ticket.Id,
		Status:   (map[bool]string{true: "success", false: "rejected"})[output.Success],
	}
	jsonDto, err := json.Marshal(paymentDto)
	if err != nil {
		return err
	}
	err = u.queue.Publish("payment.changed", jsonDto)
	return err
}
