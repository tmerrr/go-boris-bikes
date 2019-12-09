package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBikeIsWorking(t *testing.T) {
	b := bike{}
	assert.False(t, b.isWorking, "Bike isWorking should be false by default")

	b = bike{true}
	assert.True(t, b.isWorking, "Bike isWorking should be true")
}

func TestDockingStation(t *testing.T) {
	ds := dockingStation{}
	l := len(ds)

	assert.Equal(t, l, 0, "Docking Station should have 0 bikes")
}

func TestReleaseBike(t *testing.T) {
	ds := dockingStation{}

	var b interface{} = ds.releaseBike()

	_, isBike := b.(bike)
	assert.True(t, isBike, "Should be an instance of bike")
}

func TestDockBike(t *testing.T) {
	ds := dockingStation{}

	b := ds.releaseBike()

	ds.dockBike(b)
	l := len(ds)

	assert.Equal(t, l, 1, "Docking Station should now have 1 bike")
}
