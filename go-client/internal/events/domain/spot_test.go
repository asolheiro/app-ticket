package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewSpot(t *testing.T) {
	event, err := NewEvent(
		"Concert", "Stadium", "Music Inc.", 
		"http://x.jpg", RatingLivre, time.Now().Add(24*time.Hour), 
		100, 1, 50.0,
	)
	assert.Nil(t, err)
	assert.NotNil(t, event)

	spot := &Spot{
		EventID: event.ID,
		Name: "",
		Status: SpotStatusAvailable,
	}

	err = spot.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "invalid spot name, spot name must be at least 2 (two) characters long", err.Error())

	spot.Name = "1A"
	err = spot.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "spot name must start with a letter", err.Error())

	spot.Name = "A"
	err = spot.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "invalid spot name, spot name must be at least 2 (two) characters long", err.Error())
}

func TestSpot_Reserve(t *testing.T) {
	event, err := NewEvent(
		"Concert", "Stadium", "Music Inc.", 
		"http://x.jpg", RatingLivre, time.Now().Add(24*time.Hour), 
		100, 1, 50.0,
	)
	assert.Nil(t, err)
	assert.NotNil(t, event)

	spot, err := NewSpot(event, "A1")
	assert.Nil(t, err)
	assert.NotNil(t, spot)

	ticketID := "ticketID_123"
	err = spot.ReserveSpot(ticketID)
	assert.Nil(t, err)
	assert.Equal(t, SpotStatusSold, spot.Status)
	assert.Equal(t, ticketID, spot.TicketID)

}