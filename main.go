package main

import (
	"fmt"
	"os"
)

func main() {
	ds := dockingStation{}
	b, err := ds.releaseBike()
	if err != nil {
		fmt.Println("Error releasing bike:", err)
		os.Exit(1)
	}
	fmt.Println("Is the bike working:", b.isWorking)
	ds.dockBike(b)
	fmt.Printf("Docking station has %v bikes\n", len(ds))
}
