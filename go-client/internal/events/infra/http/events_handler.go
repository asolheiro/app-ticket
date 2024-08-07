package http

import (
	"encoding/json"
	"net/http"

	"github.com/rmndvngrpslhr/app-ticket/go-client/internal/events/usecase"
)

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

// TODO CreateEventsUseCase
// TODO CreateSpotUseCase

// ListEvents handles the request to list all events.
func (h *EventsHandler) ListEvents(w http.ResponseWriter, r *http.Request) {
	output, err := h.listEventsUseCase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func (h *EventsHandler) GetEvents(w http.ResponseWriter, r *http.Request) {
	eventID := r.PathValue("eventID")
	input := usecase.GetEventInputDTO{ID: eventID}

	output, err := h.getEventUseCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func (h *EventsHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateEventInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	output, err := h.createEventUseCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func (h *EventsHandler) BuyTickets(w http.ResponseWriter, r *http.Request) {
	var input usecase.BuyTicketsInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	output, err := h.buyTicketsUseCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func (h *EventsHandler) CreateSpots(w http.ResponseWriter, r *http.Request) {
	eventID := r.PathValue("eventID")
	var input usecase.CreateSpotsInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	input.EventID = eventID

	output, err := h.createSpotUseCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader((http.StatusCreated))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func (h *EventsHandler) ListSpots(w http.ResponseWriter, r *http.Request) {
	eventID := r.PathValue("eventID")
	input := usecase.ListSpotsInputDTO{EventID: eventID}

	output, err := h.listSpotsUseCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

// writeErrorresponse writes an error response in JSON format.
func (h *EventsHandler) writeErrorresponse(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Message: message})
}

// CreateSpotsRequest representes the input for creating spots.
type CreateSpotsRequest struct {
	NumberOfSpots int `json:"number_of_spots"`
}
