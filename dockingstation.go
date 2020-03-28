package main

import (
	"errors"
)

type dockingStation struct {
	bikes    []bike
	capacity int
}

func newDockingStation(bikeCount int, capacity int) (*dockingStation, error) {
	if bikeCount > capacity {
		return nil, errors.New("Bike count cannot exceed capacity")
	}
	if bikeCount < 0 {
		return nil, errors.New("Bike count cannot be negative value")
	}

	bikes := []bike{}
	for i := 0; i < bikeCount; i++ {
		bikes = append(bikes, newBike())
	}

	return &dockingStation{bikes, capacity}, nil
}

func (ds *dockingStation) releaseBike() (*bike, error) {
	for i, b := range ds.bikes {
		if b.isWorking == true {
			ds.bikes = append(ds.bikes[:i], ds.bikes[i+1:]...)
			b.SetIsDocked(false)
			return &b, nil
		}
	}
	// if no bikes found:
	return nil, errors.New("No Working Bikes Available")
}

func (ds *dockingStation) dockBike(b *bike) error {
	if len(ds.bikes) >= ds.capacity {
		return errors.New("Unable to dock Bike. Docking Station has reached maximum capacity")
	}
	b.SetIsDocked(true)
	ds.bikes = append(ds.bikes, *b)
	return nil
}
