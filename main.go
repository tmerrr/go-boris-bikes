package main

import (
	"fmt"
	"os"
)

func main() {
	ds := newDockingStation(10)
	ds.releaseBike()
	b, err := ds.releaseBike()
	if err != nil {
		fmt.Println("Error releasing bike:", err)
		os.Exit(1)
	}
	fmt.Println("Is the bike working:", b.isWorking)
	fmt.Printf("Docking station has %v bikes\n", len(ds))
	ds.dockBike(&b)
	fmt.Printf("Docking station has %v bikes\n", len(ds))
}
