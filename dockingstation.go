package main

import (
	"errors"
)

type dockingStation []bike

func newDockingStation(n int) dockingStation {
	ds := dockingStation{}
	if n < 1 {
		return ds
	}

	for i := 0; i < n; i++ {
		b := newBike()
		ds = append(ds, b)
	}

	return ds
}

func (ds *dockingStation) releaseBike() (*bike, error) {
	for i, b := range *ds {
		if b.isWorking == true {
			(*ds) = append((*ds)[:i], (*ds)[i+1:]...)
			b.SetIsDocked(false)
			return &b, nil
		}
	}
	// if no bikes found:
	return nil, errors.New("No Working Bikes Available")
}

func (ds *dockingStation) dockBike(b *bike) {
	b.SetIsDocked(true)
	(*ds) = append((*ds), *b)
}
