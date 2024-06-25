package gateway

import (
	"fmt"

	"github.com/giovanoh/clean-architecture-go/src/application/gateway"
)

type PaymentGatewayPaypal struct {
}

func NewPaymentGatewayPaypal() *PaymentGatewayPaypal {
	return &PaymentGatewayPaypal{}
}

func (p *PaymentGatewayPaypal) ProcessPayment(input gateway.PaymentGatewayInputDto) (gateway.PaymentGatewayOutputDto, error) {
	fmt.Printf("payment_gateway_paypal.process_payment: %v\n", input)
	// Here we would call the real payment gateway
	return gateway.PaymentGatewayOutputDto{
		TicketId: input.TicketId,
		Success:  true,
	}, nil
}

/*
import PaymentGateway, { Input, Output } from "../../application/gateway/PaymentGateway";
import axios from "axios";

export default class PaymentGatewayHttp implements PaymentGateway {

	async processPayment(input: Input): Promise<Output> {
		const response = await axios.post("http://localhost:3001/process_payment", input);
		return response.data;
	}

}
*/
