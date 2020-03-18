package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDockingStation(t *testing.T) {
	ds := newDockingStation(0)
	assert.Equal(t, len(ds), 0, "Docking Station should have 0 bikes")
	// change dockingStation to be able to be instantiated with x number of bikes
	ds = newDockingStation(5)
	assert.Equal(t, len(ds), 5, "Docking station should have 5 bikes")

	var b interface{} = ds[0]
	_, isBike := b.(bike)
	assert.True(t, isBike, "Should be an instance of bike")
}

func TestReleaseBike(t *testing.T) {
	ds := newDockingStation(1)

	var b, err interface{} = ds.releaseBike()
	assert.NotNil(t, b, "Bike should not be Nil")
	assert.Nil(t, err, "err should be Nil")

	_, isBike := b.(bike)
	assert.True(t, isBike, "Should be an instance of bike")

	ds = newDockingStation(0)
	var b2, err2 interface{} = ds.releaseBike()
	assert.Nil(t, b2, "Bike should be Nil")
	assert.NotNil(t, err2, "err should not be Nil")
}

func TestDockBike(t *testing.T) {
	ds := newDockingStation(0)

	ds.dockBike(bike{})
	l := len(ds)

	assert.Equal(t, l, 1, "Docking Station should now have 1 bike")
}
