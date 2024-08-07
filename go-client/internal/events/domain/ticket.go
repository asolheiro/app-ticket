package domain

import "github.com/google/uuid"

type TicketKind string

const (
	TicketKindHalf TicketKind = "half"
	TicketKindFull TicketKind = "full"
)

type Ticket struct {
	ID         string
	EventID    string
	Spot       *Spot
	TicketKind TicketKind
	Price      float64
}

func IsValidTicketKind(TicketKind TicketKind) bool {
	return TicketKind == TicketKindHalf || TicketKind == TicketKindFull
}
// CalculatePrice calculates the ticket price based on the ticket kind.
func (t *Ticket) CalculatePrice() {
	if t.TicketKind == TicketKindHalf {
		t.Price /= 2
	}
}

// Validate checks if the ticket data is valid.
func (t *Ticket) Validate() error {
	if t.Price <= 0 {
		return ErrTicketPrice
	}
	return nil
}

// NewTicket creates a new ticket with the given parameters.
func NewTicket(event *Event, spot *Spot, ticketKind TicketKind) (*Ticket, error) {
	if !IsValidTicketKind(ticketKind) {
		return nil, ErrTicketInvalidKind
	}

	ticket := &Ticket {
		ID: uuid.New().String(),
		EventID: event.ID,	
		Spot: spot,
		TicketKind: ticketKind,
		Price: event.Price,
	}
	ticket.CalculatePrice()
	if err := ticket.Validate(); err != nil {
		return nil, err
	}
	return ticket, nil
}
