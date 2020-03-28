package main

import (
	"fmt"
	"os"
)

func checkBike(b bike) {
	fmt.Println("Is the bike working:", b.isWorking)
}

func checkBikesInDockingStation(ds *dockingStation) {
	fmt.Printf("Docking Station has %v Bikes\n", len(ds.bikes))
}

func throwError(msg string, err error) {
	fmt.Println(msg, err)
	os.Exit(1)
}

func main() {
	ds, err1 := newDockingStation(1, 5)
	if err1 != nil {
		throwError("Error creating new docking station:", err1)
	}
	checkBikesInDockingStation(ds)

	b, err2 := ds.releaseBike()
	if err2 != nil {
		throwError("Error releasing bike:", err2)
	}
	checkBikesInDockingStation(ds)

	checkBike(*b)
	b.SetIsWorking(false)
	checkBike(*b)

	ds.dockBike(b)
	b, err2 = ds.releaseBike()
	if err2 != nil {
		throwError("Error releasing bike:", err2)
	}
}
