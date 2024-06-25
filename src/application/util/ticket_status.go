package util

type TicketStatus string

const (
	TicketStatusReserved TicketStatus = "reserved"
	TicketStatusApproved TicketStatus = "approved"
	TicketStatusRejected TicketStatus = "rejected"
)
