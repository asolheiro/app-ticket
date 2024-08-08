package domain

import "github.com/google/uuid"



// Validate checks if spots data is valid.
func (s Spot) Validate() error {
	if len(s.Name) <= 1 {
		return ErrSpotInvalidName
	}
	if s.Name[0] < 'A' || s.Name[0] > 'Z' {
		return ErrSpotInitialCharacterInvalid
	}
	if s.Name[1] < '0' || s.Name[1] > '9' {
		return ErrSpotFinalCharacterInvalid
	}
	return nil
}

// NewSpot create a new spot with the given parameters.
func NewSpot(event *Event, name string) (*Spot, error) {
	spot := &Spot{
		ID:      uuid.New().String(),
		EventID: event.ID,
		Name:    name,
		Status:  SpotStatusAvailable,
	}

	if err := spot.Validate(); err != nil {
		return nil, err
	}
	return spot, nil
}

// ReserveSpot updates a spot Status on database
func (s *Spot) ReserveSpot(ticketId string) error {
	if s.Status == SpotStatusSold {
		return ErrSpotAlreadyReserved
	}
	s.Status = SpotStatusSold
	s.TicketID = ticketId
	return nil
}