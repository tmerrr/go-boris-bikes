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

func main() {
	ds, _ := newDockingStation(1, 5)
	checkBikesInDockingStation(ds)

	b, err := ds.releaseBike()
	if err != nil {
		fmt.Println("Error releasing bike:", err)
		os.Exit(1)
	}
	checkBikesInDockingStation(ds)

	checkBike(*b)
	b.SetIsWorking(false)
	checkBike(*b)

	ds.dockBike(b)
	b, err = ds.releaseBike()
	if err != nil {
		fmt.Println("Error releasing bike:", err)
		os.Exit(1)
	}
}
