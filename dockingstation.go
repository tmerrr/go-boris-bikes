package main

type dockingStation []bike

func newDockingStation(n int) dockingStation {
	ds := dockingStation{}
	if n < 1 {
		return ds
	}

	for i := 0; i < n; i++ {
		b := bike{true}
		ds = append(ds, b)
	}

	return ds
}

func (ds *dockingStation) releaseBike() (bike, error) {
	if len(ds) == 0 {
		// Can't return struct as nil.
		// The zero value of an interface is nil
		// Need to rewrite the Bike as an interface
		return nil, error
	}
	return bike{}, nil
}

func (ds *dockingStation) dockBike(b bike) {
	(*ds) = append((*ds), b)
}
