package main

type bike struct {
	isWorking bool
}

type dockingStation []bike

func (dockingStation) releaseBike() bike {
	return bike{}
}

func (ds *dockingStation) dockBike(b bike) {
	(*ds) = append((*ds), b)
}
