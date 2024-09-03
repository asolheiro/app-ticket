package http

import "github.com/rmndvngrpslhr/app-ticket/go-client/internal/events/usecase"

// EventsHandlerhandles HTTP requests for events.
type EventsHandler struct {
	buyTicketsUseCase  *usecase.BuyTicketsUseCase
	createEventUseCase *usecase.CreateEventUseCase
	createSpotUseCase  *usecase.CreateSpotsUseCase
	getEventUseCase    *usecase.GetEventUseCase
	listEventsUseCase  *usecase.ListEventsUseCase
	listSpotsUseCase   *usecase.ListSpotsUseCase
}

// NewEventshandler creates a new instance of EventsHandler.
func NewEventsHandler(
	buyTicketsUseCase *usecase.BuyTicketsUseCase,
	createEventUseCase *usecase.CreateEventUseCase,
	createSpotUseCase *usecase.CreateSpotsUseCase,
	getEventUseCase *usecase.GetEventUseCase,
	listEventsUseCase *usecase.ListEventsUseCase,
	listSpotsUseCase *usecase.ListSpotsUseCase,
) *EventsHandler {
	return &EventsHandler{
		buyTicketsUseCase:  buyTicketsUseCase,
		createEventUseCase: createEventUseCase,
		createSpotUseCase:  createSpotUseCase,
		getEventUseCase:    getEventUseCase,
		listEventsUseCase:  listEventsUseCase,
		listSpotsUseCase:   listSpotsUseCase,
	}
}