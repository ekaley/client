package api

import (
	"errors"

	"github.com/RackHD/ipam/resources"
)

//Subnets can be used to query the Subnets routes
type Subnets struct {
	c *Client
}

//Subnets returns a handle to the Subnets routes
func (c *Client) Subnets() *Subnets {
	return &Subnets{c}
}

// Index returns a list of Subnets.
func (s *Subnets) Index(pLocation string) (resources.SubnetsV1, error) {
	rSubnets, err := s.c.ReceiveResource("GET", pLocation+"/subnets", "", "")
	if err != nil {
		return resources.SubnetsV1{}, err
	}
	if subnets, ok := rSubnets.(*resources.SubnetsV1); ok {
		return *subnets, nil
	}
	return resources.SubnetsV1{}, errors.New("Subnet Index call error.")
}

// Create a subnet and return the location.
func (s *Subnets) Creates(pLocation string, rSubnet resources.SubnetV1) (string, error) {
	sLocation, err := s.c.SendResource("POST", pLocation+"/subnets", &rSubnet)
	if err != nil {
		return "", err
	}
	return sLocation, nil
}

// CreateShowSubnet creates a subnet and then returns that subnet.
func (s *Subnets) CreateShowSubnet(pLocation string, rSubnet resources.SubnetV1) (resources.SubnetV1, error) {
	rSubnetOut, err := s.c.SendReceiveResource("POST", "GET", pLocation+"/subnets", &rSubnet)
	if err != nil {
		return resources.SubnetV1{}, err
	}
	if subnet, ok := rSubnetOut.(*resources.SubnetV1); ok {
		return *subnet, nil
	}
	return resources.SubnetV1{}, errors.New("CreateShowSubnet call error.")
}

// Show returns the requested subnet.
func (s *Subnets) Show(sLocation string, rSubnetIn resources.SubnetV1) (resources.SubnetV1, error) {
	rSubnetOut, err := s.c.ReceiveResource("GET", sLocation, rSubnetIn.Type(), rSubnetIn.Version())
	if err != nil {
		return resources.SubnetV1{}, err
	}
	if subnet, ok := rSubnetOut.(*resources.SubnetV1); ok {
		return *subnet, nil
	}
	return resources.SubnetV1{}, errors.New("Subnet Show call error.")
}

// Update updates the requested subnet and returns its location.
func (s *Subnets) Update(sLocation string, rSubnet resources.SubnetV1) (string, error) {
	sLocation, err := s.c.SendResource("PATCH", sLocation, &rSubnet)
	if err != nil {
		return "", err
	}
	return sLocation, nil
}

// UpdateShowSubnet updates a Subnet and then returns that Subnet.
func (s *Subnets) UpdateShowSubnet(sLocation string, rSubnet resources.SubnetV1) (resources.SubnetV1, error) {
	rSubnetOut, err := s.c.SendReceiveResource("PATCH", "GET", sLocation, &rSubnet)
	if err != nil {
		return resources.SubnetV1{}, err
	}
	if subnet, ok := rSubnetOut.(*resources.SubnetV1); ok {
		return *subnet, nil
	}
	return resources.SubnetV1{}, errors.New("UpdateShowSubnet call error.")
}

// Delete removed the requested subnet and returns the location.
func (s *Subnets) Delete(sLocation string, rSubnet resources.SubnetV1) (string, error) {
	sLocation, err := s.c.SendResource("DELETE", sLocation, &rSubnet)
	if err != nil {
		return "", err
	}
	return sLocation, nil
}
