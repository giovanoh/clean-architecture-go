package gateway

type PaymentGatewayInputDto struct {
	TicketId        string  `json:"ticket_id"`
	Price           float64 `json:"price"`
	CreditCardToken string  `json:"credit_card_token"`
}

type PaymentGatewayOutputDto struct {
	TicketId string `json:"ticket_id"`
	Success  bool   `json:"success"`
}

type PaymentGateway interface {
	ProcessPayment(input PaymentGatewayInputDto) (PaymentGatewayOutputDto, error)
}
