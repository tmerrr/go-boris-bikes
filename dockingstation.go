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

func (ds *dockingStation) releaseBike() (bike, error) {
	var bikeFound bool
	var index int
	var bk bike
	for i, b := range *ds {
		if b.isWorking == true {
			bikeFound = true
			index = i
			bk = b
			break
		}
	}
	if bikeFound == true {
		(*ds) = append((*ds)[:index], (*ds)[index+1:]...)
		bk.release()
		return bk, nil
	}

	return bike{}, errors.New("No Working Bikes Available")
}

func (ds *dockingStation) dockBike(b bike) {
	(*ds) = append((*ds), b)
}
