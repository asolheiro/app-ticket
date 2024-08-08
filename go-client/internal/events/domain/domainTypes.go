package domain

import "time"

// Rating represents the age restriction for an event.
type Rating string

const (
	RatingLivre Rating = "L"
	Rating10    Rating = "L10"
	Rating12    Rating = "L12"
	Rating14    Rating = "L14"
	Rating16    Rating = "L16"
	Rating18    Rating = "L18"
)

// Event represents an event with tickets and spots.
type Event struct {
	ID           string
	Name         string
	Location     string
	Organization string
	Rating       Rating
	Date         time.Time
	ImageURL     string
	Capacity     int
	Price        float64
	PartnerID    int
	Spots        []Spot
	Tickets      []Ticket
}

// SpotStatus represents the spot status: Sold or Available
type SpotStatus string

const (
	SpotStatusAvailable SpotStatus = "available"
	SpotStatusSold      SpotStatus = "sold"
)

// Spot represents a spot for an event.
type Spot struct {
	ID       string
	EventID  string
	Name     string
	Status   SpotStatus
	TicketID string
}

// TicketKind represents the kind of a ticket: Full or Half.
type TicketKind string

const (
	TicketKindHalf TicketKind = "half"
	TicketKindFull TicketKind = "full"
)

// Ticket represents a ticket for an event.
type Ticket struct {
	ID         string
	EventID    string
	Spot       *Spot
	TicketKind TicketKind
	Price      float64
}
