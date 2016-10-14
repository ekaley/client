package api

import (
	"errors"

	"github.com/RackHD/ipam/resources"
)

//Reservations can be used to query the Reservations routes
type Reservations struct {
	c *Client
}

//Reservations returns a handle to the Reservations routes
func (c *Client) Reservations() *Reservations {
	return &Reservations{c}
}

// Index returns a list of Reservations.
func (s *Reservations) Index(sLocation string) (resources.ReservationsV1, error) {
	rReservations, err := s.c.ReceiveResource("GET", sLocation+"/reservations", "", "")
	if err != nil {
		return resources.ReservationsV1{}, err
	}
	if reservations, ok := rReservations.(*resources.ReservationsV1); ok {
		return *reservations, nil
	}
	return resources.ReservationsV1{}, errors.New("Reservation Index call error.")
}

// Create a Reservation and return the location.
func (s *Reservations) Create(sLocation string, rReservation resources.ReservationV1) (string, error) {
	rLocation, err := s.c.SendResource("POST", sLocation+"/reservations", &rReservation)
	if err != nil {
		return "", err
	}
	return rLocation, nil
}

// CreateShowReservation creates a Reservation and then returns that Reservation.
func (s *Reservations) CreateShowReservation(sLocation string, rReservation resources.ReservationV1) (resources.ReservationV1, error) {
	rReservationOut, err := s.c.SendReceiveResource("POST", "GET", sLocation+"/reservations", &rReservation)
	if err != nil {
		return resources.ReservationV1{}, err
	}
	if reservation, ok := rReservationOut.(*resources.ReservationV1); ok {
		return *reservation, nil
	}
	return resources.ReservationV1{}, errors.New("CreateShowReservation call error.")
}

// Show returns the requested Reservation.
func (s *Reservations) Show(rID string, rReservationIn resources.ReservationV1) (resources.ReservationV1, error) {
	rReservationOut, err := s.c.ReceiveResource("GET", "/reservations"+rID, rReservationIn.Type(), rReservationIn.Version())
	if err != nil {
		return resources.ReservationV1{}, err
	}
	if reservation, ok := rReservationOut.(*resources.ReservationV1); ok {
		return *reservation, nil
	}
	return resources.ReservationV1{}, errors.New("Reservation Show call error.")
}

// Update updates the requested Reservation and returns its location.
func (s *Reservations) Update(rLocation string, rReservation resources.ReservationV1) (string, error) {
	rLocation, err := s.c.SendResource("PATCH", rLocation, &rReservation)
	if err != nil {
		return "", err
	}
	return rLocation, nil
}

// UpdateShowReservation updates a Reservation and then returns that Reservation.
func (s *Reservations) UpdateShowReservation(rLocation string, rReservation resources.ReservationV1) (resources.ReservationV1, error) {
	rReservationOut, err := s.c.SendReceiveResource("PATCH", "GET", rLocation, &rReservation)
	if err != nil {
		return resources.ReservationV1{}, err
	}
	if reservation, ok := rReservationOut.(*resources.ReservationV1); ok {
		return *reservation, nil
	}
	return resources.ReservationV1{}, errors.New("UpdateShowReservation call error.")
}

// Delete removed the requested Reservation and returns the location.
func (s *Reservations) Delete(rLocation string, rReservation resources.ReservationV1) (string, error) {
	rLocation, err := s.c.SendResource("DELETE", rLocation, &rReservation)
	if err != nil {
		return "", err
	}
	return rLocation, nil
}
