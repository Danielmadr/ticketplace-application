package domain

import (
	"errors"

	"github.com/google/uuid"
)

// Errors
var (
	ErrInvalidTicketType = errors.New("invalid ticket type")
	ErrTicketPriceZero   = errors.New("ticket price must be greater than zero")
)


type TicketType string

const (
	TicketTypeHalf TicketType = "HALF"
	TicketTypeFull TicketType = "FULL"
)

type Ticket struct {
	ID       string
	EventID  string
	Spot *Spot
	TicketType TicketType
	Price    float64
}

func IsValidTicketType(ticketType TicketType) bool {
	return ticketType == TicketTypeHalf || ticketType == TicketTypeFull
}

// CalculatePrice calculates the price based on the ticket type.
func (t *Ticket) CalculatePrice() {
	if t.TicketType == TicketTypeHalf {
		t.Price /= 2
	}
}


// Validate checks if the ticket price is valid.
func (t *Ticket) Validate() error {
	if t.Price <= 0 {
		return ErrTicketPriceZero
	}
	return nil
}

func NewTicket(event *Event, spot *Spot, ticketType TicketType) (*Ticket, error) {
	if !IsValidTicketType(ticketType) {
		return nil, ErrInvalidTicketType
	}

	ticket := &Ticket{
		ID:         uuid.New().String(),
		EventID:    event.ID,
		Spot:       spot,
		TicketType: ticketType,
		Price:      event.Price,
	}
	ticket.CalculatePrice()
	if err := ticket.Validate(); err != nil {
		return nil, err
	}
	return ticket, nil
}
