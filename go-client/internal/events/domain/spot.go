package domain

import "github.com/google/uuid"


type SpotStatus string

const (
	SpotStatusAvailable SpotStatus = "available"
	SpotStatusSold      SpotStatus = "sold"
)

type Spot struct {
	ID       string
	EventID  string
	Name     string
	Status   SpotStatus
	TicketID string
}

// Validate if spots data is valid
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

// Create a new spot
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
