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
func (s *Reservations) Index(sID string) (resources.ReservationsV1, error) {
	rReservations, err := s.c.ReceiveResource("GET", "/subnets/"+sID+"/reservations", "", "")
	if err != nil {
		return resources.ReservationsV1{}, err
	}
	if reservations, ok := rReservations.(*resources.ReservationsV1); ok {
		return *reservations, nil
	}
	return resources.ReservationsV1{}, errors.New("Reservation Index call error.")
}

// Create a Reservation and return the location.
func (s *Reservations) Create(sID string, rReservation resources.ReservationV1) (string, error) {
	rLocation, err := s.c.SendResource("POST", "/subnets/"+sID+"/reservations", &rReservation)
	if err != nil {
		return "", err
	}
	return rLocation, nil
}

// CreateShowReservation creates a Reservation and then returns that Reservation.
func (s *Reservations) CreateShowReservation(sID string, rReservation resources.ReservationV1) (resources.ReservationV1, error) {
	rReservationOut, err := s.c.SendReceiveResource("POST", "GET", "/subnets/"+sID+"/reservations", &rReservation)
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
func (s *Reservations) Update(rID string, rReservation resources.ReservationV1) (string, error) {
	rLocation, err := s.c.SendResource("PATCH", "/reservations"+rID, &rReservation)
	if err != nil {
		return "", err
	}
	return rLocation, nil
}

// UpdateShowReservation updates a Reservation and then returns that Reservation.
func (s *Reservations) UpdateShowReservation(rID string, rReservation resources.ReservationV1) (resources.ReservationV1, error) {
	rReservationOut, err := s.c.SendReceiveResource("PATCH", "GET", "/reservations"+rID, &rReservation)
	if err != nil {
		return resources.ReservationV1{}, err
	}
	if reservation, ok := rReservationOut.(*resources.ReservationV1); ok {
		return *reservation, nil
	}
	return resources.ReservationV1{}, errors.New("UpdateShowReservation call error.")
}

// Delete removed the requested Reservation and returns the location.
func (s *Reservations) Delete(rID string, rReservation resources.ReservationV1) (string, error) {
	rLocation, err := s.c.SendResource("DELETE", "/reservations"+rID, &rReservation)
	if err != nil {
		return "", err
	}
	return rLocation, nil
}
