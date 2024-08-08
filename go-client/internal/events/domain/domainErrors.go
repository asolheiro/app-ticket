package domain

import "errors"

//Erros Event
var (
	ErrEventNameRequired = errors.New("event name is required")
	ErrEventFutureDateRequired = errors.New("event date must be in the future")
	ErrEventCapacityInvalid = errors.New("event capacity must be greater than zero")
	ErrEventPriceInvalid = errors.New("ticket price must be greater than zero")
)

// Errors Spot
var (
	ErrSpotInvalidName = errors.New("invalid spot name, spot name must be at least 2 (two) characters long")
	ErrSpotInitialCharacterInvalid = errors.New("spot name must start with a letter")
	ErrSpotFinalCharacterInvalid = errors.New("Spot name must end with a number")
	ErrSpotInvalidNumber = errors.New("invalid spot number")
	ErrSpotNotFound = errors.New("spot not found")
	ErrSpotAlreadyReserved = errors.New("invalid spot Spot already reserved")
)

// Errors Ticket
var (
	ErrTicketPrice = errors.New("ticket price must be greater than zero")
	ErrTicketInvalidKind = errors.New("invalid ticket kind")
)

// Errors Service
var (
	ErrServiceInvalidQuantity = errors.New("quantity must be at least 1")
)