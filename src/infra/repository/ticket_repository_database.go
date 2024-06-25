package repository

import (
	"database/sql"

	"github.com/giovanoh/clean-architecture-go/src/domain/entity"
)

type TicketRepositoryDatabase struct {
	Db *sql.DB
}

func NewTicketRepositoryDatabase(db *sql.DB) *TicketRepositoryDatabase {
	return &TicketRepositoryDatabase{
		Db: db,
	}
}

func (t *TicketRepositoryDatabase) Create(ticket *entity.Ticket) error {
	_, err := t.Db.Exec("INSERT INTO tickets (id, name, price) VALUES (?, ?, ?)", ticket.Id, ticket.Name, ticket.Price)
	if err != nil {
		return err
	}
	return nil
}

func (t *TicketRepositoryDatabase) Update(ticket *entity.Ticket) error {
	_, err := t.Db.Exec("UPDATE tickets SET status = ? WHERE id = ?", ticket.Status, ticket.Id)
	if err != nil {
		return err
	}
	return nil
}

func (t *TicketRepositoryDatabase) GetTicketById(id string) (*entity.Ticket, error) {
	row := t.Db.QueryRow("SELECT id, name, price FROM tickets WHERE id = ?", id)
	ticket := &entity.Ticket{}
	err := row.Scan(&ticket.Id, &ticket.Name, &ticket.Price)
	if err != nil {
		return nil, err
	}
	return ticket, nil
}

func (t *TicketRepositoryDatabase) GetTicketsByEmail(email string) ([]*entity.Ticket, error) {
	rows, err := t.Db.Query("SELECT id, name, price FROM tickets WHERE email = ?", email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tickets := []*entity.Ticket{}
	for rows.Next() {
		ticket := &entity.Ticket{}
		err := rows.Scan(&ticket.Id, &ticket.Name, &ticket.Price)
		if err != nil {
			return nil, err
		}
		tickets = append(tickets, ticket)
	}
	return tickets, nil
}
