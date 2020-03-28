package main

import (
	"errors"
)

type dockingStation struct {
	bikes    []bike
	capacity int
}

func newDockingStation(bikeCount int, capacity int) dockingStation {
	bikes := []bike{}
	if bikeCount < 1 {
		return dockingStation{bikes, capacity}
	}

	for i := 0; i < bikeCount; i++ {
		b := newBike()
		bikes = append(bikes, b)
	}

	return dockingStation{bikes, capacity}
}

func (ds *dockingStation) releaseBike() (*bike, error) {
	for i, b := range ds.bikes {
		if b.isWorking == true {
			(ds.bikes) = append((ds.bikes)[:i], (ds.bikes)[i+1:]...)
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
