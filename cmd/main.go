package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/RackHD/ipam/resources"
	"github.com/josephgorse/ipam-client"
)

var ipam string

func init() {
	flag.StringVar(&ipam, "ipam", "127.0.0.1:8000", "address:port of ipam")
}

func main() {
	//Bootstrap
	ipamClient := api.NewClient(ipam)

	//New Pools
	p := ipamClient.Pools()

	rPool := &resources.PoolV1{
		Name:     "Pool1",
		Metadata: "yodawg I heard you like interfaces",
	}

	pool, err := p.CreateShowPool(*rPool)
	if err == nil {
		fmt.Printf("%+v\n", pool)
	}

	time.Sleep(1 * time.Second)
	//New Subnets
	s := ipamClient.Subnets()

	rSubnet := &resources.SubnetV1{
		Name:  "Subnet1",
		Pool:  pool.ID,
		Start: "192.168.1.10",
		End:   "192.168.1.20",
	}

	subnet, err := s.CreateShowSubnet(pool.ID, *rSubnet)
	if err == nil {
		fmt.Printf("%+v\n", subnet)
	}

	sub, err := s.Show(subnet.ID, resources.SubnetV1{})

	if err == nil {
		fmt.Printf("%+v\n", sub)
	}

	time.Sleep(1 * time.Second)
	//New Reservations
	r := ipamClient.Reservations()

	rReservation := &resources.ReservationV1{
		Name:   "Reservation1",
		Subnet: subnet.ID,
	}

	reservation, err := r.CreateShowReservation(subnet.ID, *rReservation)
	if err == nil {
		fmt.Printf("%+v\n\n", reservation)
	}

	time.Sleep(1 * time.Second)
	//New Leases
	l := ipamClient.Leases()

	poolIndex, err := p.Index()
	if err != nil {
	}
	fmt.Printf("%+v\n\n", poolIndex)
	time.Sleep(1 * time.Second)

	subnetIndex, err := s.Index(pool.ID)
	if err != nil {
	}
	fmt.Printf("%+v\n\n", subnetIndex)
	time.Sleep(1 * time.Second)

	reservationID, err := r.Index(subnet.ID)
	if err != nil {
	}
	fmt.Printf("%+v\n\n", reservationID)
	time.Sleep(1 * time.Second)

	leaseIndex, err := l.Index(reservation.ID)
	if err != nil {
	}
	fmt.Printf("%+v\n\n", leaseIndex)

	fmt.Println("Finished")
	return
}
