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
	assert.Equal(t, len(ds), 1, "Docking Station should have length of 1")
	// asserts can successfully release a bike
	b, err := ds.releaseBike()
	assert.Equal(t, len(ds), 0, "Docking Station should have length of 0")
	assert.NotNil(t, b, "Bike should not be Nil")
	assert.Nil(t, err, "err should be Nil")
	assert.True(t, b.isWorking, "Bike should be working")
	assert.False(t, b.isDocked, "Bike should not be docked")
	// asserts that cannot release a bike if only broken available
	ds = append(ds, bike{false, true})
	b, err = ds.releaseBike()
	assert.NotNil(t, b, "Error should exist")

	ds = newDockingStation(1)
	var b2, _ interface{} = ds.releaseBike()
	_, isBike := b2.(bike)
	assert.True(t, isBike, "Should be an instance of bike")

	ds = newDockingStation(0)
	var _, err2 interface{} = ds.releaseBike()
	assert.NotNil(t, err2, "err should not be Nil")
	_, isErr := err2.(error)
	assert.True(t, isErr, "Should be an error")
}

func TestDockBike(t *testing.T) {
	ds := newDockingStation(0)

	b := bike{isDocked: false}
	ds.dockBike(&b)

	l := len(ds)
	assert.Equal(t, l, 1, "Docking Station should now have 1 bike")
	assert.True(t, b.isDocked, "Bike should now be docked")
}
