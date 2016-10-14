package api

import (
	"errors"

	"github.com/RackHD/ipam/resources"
)

//Pools can be used to query the Pools routes
type Pools struct {
	c *Client
}

//Pools returns a handle to the Pools routes
func (c *Client) Pools() *Pools {
	return &Pools{c}
}

// Index returns a list of Pools.
func (p *Pools) Index() (resources.PoolsV1, error) {
	rPools, err := p.c.ReceiveResource("GET", "/pools", "", "")
	if err != nil {
		return resources.PoolsV1{}, err
	}
	if pools, ok := rPools.(*resources.PoolsV1); ok {
		return *pools, nil
	}
	return resources.PoolsV1{}, errors.New("Pool Index call error.")
}

// Create a pool and returns the location.
func (p *Pools) Create(rPool resources.PoolV1) (string, error) {
	pLocation, err := p.c.SendResource("POST", "/pools", &rPool)
	if err != nil {
		return "", err
	}
	return pLocation, nil
}

// CreateShowPool creates a pool and then returns that pool.
func (p *Pools) CreateShowPool(rPool resources.PoolV1) (resources.PoolV1, error) {
	rPoolOut, err := p.c.SendReceiveResource("POST", "GET", "/pools", &rPool)
	if err != nil {
		return resources.PoolV1{}, err
	}
	if pool, ok := rPoolOut.(*resources.PoolV1); ok {
		return *pool, nil
	}
	return resources.PoolV1{}, errors.New("CreateShowPool call error.")
}

// Show returns the requested Pool.
func (p *Pools) Show(pLocation string, rPoolIn resources.PoolV1) (resources.PoolV1, error) {
	rPoolOut, err := p.c.ReceiveResource("GET", pLocation, rPoolIn.Type(), rPoolIn.Version())
	if err != nil {
		return resources.PoolV1{}, err
	}
	if pool, ok := rPoolOut.(*resources.PoolV1); ok {
		return *pool, nil
	}
	return resources.PoolV1{}, errors.New("Pools Show call error.")
}

// Update updates the requested Pool and returns its location.
func (p *Pools) Update(pLocation string, rPool resources.PoolV1) (string, error) {
	location, err := p.c.SendResource("PATCH", pLocation, &rPool)
	if err != nil {
		return "", err
	}
	return location, nil
}

// UpdateShowPool updates a pool and then returns that pool.
func (p *Pools) UpdateShowPool(pLocation string, rPool resources.PoolV1) (resources.PoolV1, error) {
	rPoolOut, err := p.c.SendReceiveResource("PATCH", "GET", pLocation, &rPool)
	if err != nil {
		return resources.PoolV1{}, err
	}
	if pools, ok := rPoolOut.(*resources.PoolV1); ok {
		return *pools, nil
	}
	return resources.PoolV1{}, errors.New("UpdateShowPool call error.")
}

// Delete removes the requested Pool and returns the location.
func (p *Pools) Delete(pLocation string, rPool resources.PoolV1) (string, error) {
	location, err := p.c.SendResource("DELETE", pLocation, &rPool)
	if err != nil {
		return "", err
	}
	return location, nil
}
