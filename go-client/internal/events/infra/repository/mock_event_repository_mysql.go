package repository

import (
	"github.com/rmndvngrpslhr/app-ticket/go-client/internal/events/domain"
	"github.com/stretchr/testify/mock"
)

type mockEventRepository struct {
	mock.Mock
}

func (m *mockEventRepository) ListEvents() ([]domain.Event, error) {
	args := m.Called()
	return args.Get(0).([]domain.Event), args.Error(1)
}

func (m *mockEventRepository) FindSpotsByEventID(eventID string) ([]*domain.Spot, error) {
	args := m.Called(eventID)
	return args.Get(0).([]*domain.Spot), args.Error(1)
}

func (m *mockEventRepository) FindSpotByID(spotID string) (*domain.Spot, error) {
	args := m.Called(spotID)
	return args.Get(0).(*domain.Spot), args.Error(1)
}

func (m *mockEventRepository) FindSpotByName(eventID, name string) (*domain.Spot, error) {
	args := m.Called(eventID, name)
	return args.Get(0).(*domain.Spot), args.Error(1)
}

func (m *mockEventRepository) CreateEvent(event *domain.Event) error {
	args := m.Called(event)
	return args.Error(0)
}

func (m *mockEventRepository) CreateSpot(spot *domain.Spot) error {
	args := m.Called(spot)
	return args.Error(0)
}

func (m *mockEventRepository) CreateTicket(ticket *domain.Ticket) error {
	args := m.Called(ticket)
	return args.Error(0)
}

func (m *mockEventRepository) ReserveSpot(spotID, ticketID string) error {
	args := m.Called(spotID, ticketID)
	return args.Error(0)
}