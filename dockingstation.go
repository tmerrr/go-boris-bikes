package main

type dockingStation []bike

func (dockingStation) releaseBike() (bike, error) {
	return bike{}, nil
}

func (ds *dockingStation) dockBike(b bike) {
	(*ds) = append((*ds), b)
}
