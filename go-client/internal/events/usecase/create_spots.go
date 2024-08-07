package usecase

import (
	"fmt"

	"github.com/rmndvngrpslhr/app-ticket/go-client/internal/events/domain"
)

// CreateSpotsInputDTO represents the input data required to create a spot.
type CreateSpotsInputDTO struct {
	EventID       string `json:"event_id"`
	NumberOfSpots int    `json:"number_of_spots"`
}

// CreateSpotsOutputDTO represents the output data after creating a spot.
type CreateSpotsOutputDTO struct {
	Spots []SpotDTO `json:"spots"`
}

// CreateSpotsUseCase defines the use case for creting spots.
type CreateSpotsUseCase struct {
	repo domain.EventRepository
}

// NewCreateSpotsUseCase creates a new instance of CreateSpotsUseCase.
func NewCreateSpotsUseCase(repo domain.EventRepository) *CreateSpotsUseCase {
	return &CreateSpotsUseCase{repo: repo}
}

// Execute executes the use case to create spots
func (uc *CreateSpotsUseCase) Execute(input CreateSpotsInputDTO) (*CreateSpotsOutputDTO, error) {
	event, err := uc.repo.FindEventByID(input.EventID)
	if err != nil {
		return nil, err
	}

	spots := make([]domain.Spot, input.NumberOfSpots)
	for i := 0; i < input.NumberOfSpots; i++ {
		spotName := generateSpotName(i)
		spot, err := domain.NewSpot(event, spotName)
		if err != nil {
			return nil, err
		}
		if err := uc.repo.CreateSpot(spot); err != nil {
			return nil, err
		}
		spots[i] = *spot
	}

	spotsDTO := make([]SpotDTO, len(spots))
	for i, spot := range spots {
		spotsDTO[i] = SpotDTO{
			ID:       spot.ID,
			Name:     spot.Name,
			Status:   string(spot.Status),
			TicketID: spot.TicketID,
		}
	}

	return &CreateSpotsOutputDTO{Spots: spotsDTO}, nil
}

func generateSpotName(index int) string {
	letter := 'A' + rune(index/10)
	number := index%10 + 1
	return fmt.Sprintf("%c%d", letter, number)
}
