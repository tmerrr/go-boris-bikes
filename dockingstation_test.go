package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDockingStation(t *testing.T) {
	ds := dockingStation{}
	assert.Equal(t, len(ds), 0, "Docking Station should have 0 bikes")
	// change dockingStation to be able to be instantiated with x number of bikes
	ds = dockingStation{5}
	assert.Equal(t, len(ds), 5, "Docking station should have 5 bikes")

	var b interface{} = ds[0]
	_, isBike := b.(bike)
	assert.True(t, isBike, "Should be an instance of bike")
}

func TestReleaseBike(t *testing.T) {
	ds := dockingStation{}

	var b, err interface{} = ds.releaseBike()

	assert.NotNil(t, b, "Bike should not be Nil")
	assert.Nil(t, err, "err should be Nil")

	_, isBike := b.(bike)
	assert.True(t, isBike, "Should be an instance of bike")
}

func TestDockBike(t *testing.T) {
	ds := dockingStation{}

	ds.dockBike(bike{})
	l := len(ds)

	assert.Equal(t, l, 1, "Docking Station should now have 1 bike")
}
