package model

import "time"

// BookingRequest represents the request to bookings adatpter service.
type BookingRequest struct {
	Items []Booking `json:"items"`
}

// Booking represents a booking that needs to be booked on bookings adatpter service side.
type Booking struct {
	ID          string    `json:"id"`
	Amount      float64   `json:"amount"`
	BookingType string    `json:"type"`
	BookingDate time.Time `json:"booking_date"`
}
