package main

type bike struct {
	isWorking bool
}

type dockingStation []bike

func (dockingStation) releaseBike() bike {
	return bike{}
}
