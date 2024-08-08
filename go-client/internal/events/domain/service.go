package domain

import "fmt"

type spotService struct {}

// NewSpotService creates a new instance of SpotService.
func NewSpotService() *spotService {
	return &spotService{}
}

// GenerateSpots generates the specified number of spots for an given event.
func (s *spotService) GenerateSpots(event *Event, quantity int) error {
	if quantity <= 0 {
		return ErrServiceInvalidQuantity
	}
	for i := range quantity {
		spotName := fmt.Sprintf("%c%d", 'A'+i/10, i%10+1)
		spot, err := NewSpot(event, spotName)
		if err != nil {
			return err
		}
		event.Spots = append(event.Spots, *spot)
	}
	return nil
}