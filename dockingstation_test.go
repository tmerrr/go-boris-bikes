package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDockingStation(t *testing.T) {
	// asserts docking station is instantiated with 0 bikes and capacity of 3
	ds, err := newDockingStation(0, 3)
	assert.Equal(t, len(ds.bikes), 0, "Docking Station should have 0 bikes")
	assert.Equal(t, ds.capacity, 3, "expected capacity of 3")
	assert.Nil(t, err, "expected err to be nil")
	// asserts docking station is instantiated with 5 bikes and capacity of 5
	ds, err = newDockingStation(5, 5)
	assert.Equal(t, len(ds.bikes), 5, "Docking station should have 5 bikes")
	assert.Equal(t, ds.capacity, 5, "expected capacity of 5")
	assert.Nil(t, err, "expected err to be nil")
	var b interface{} = ds.bikes[0]
	_, isBike := b.(bike)
	assert.True(t, isBike, "Should be an instance of bike")
	// asserts that an error is returned if attempting to initialise
	// a docking station with more bikes than max capacity
	ds, err = newDockingStation(2, 1)
	assert.Nil(t, ds, "expected docking station to be nil")
	assert.NotNil(t, err, "expected err not to be nil")
	_, isErr := err.(error)
	assert.True(t, isErr, "expected err to be of type error")
}

func TestReleaseBike(t *testing.T) {
	ds, _ := newDockingStation(1, 3)
	assert.Equal(t, len(ds.bikes), 1, "Docking Station should have length of 1")
	// asserts can successfully release a bike
	bp, err := ds.releaseBike()
	assert.Equal(t, len(ds.bikes), 0, "Docking Station should have length of 0")
	assert.NotNil(t, bp, "Bike should not be Nil")
	assert.Nil(t, err, "err should be Nil")
	assert.True(t, bp.isWorking, "Bike should be working")
	assert.False(t, bp.isDocked, "Bike should not be docked")
	// asserts bp is a pointer to a bike
	var b interface{} = *bp
	_, isBike := b.(bike)
	assert.True(t, isBike, "Should be an instance of bike")
	// asserts that cannot release a bike if only broken available
	ds.bikes = append(ds.bikes, bike{false, true})
	bp, err = ds.releaseBike()
	assert.Nil(t, bp, "Bike should be Nil")
	assert.NotNil(t, err, "Error should exist")
	_, isErr := err.(error)
	assert.True(t, isErr, "Should be an instance of error")
	//asserts error is returned when no bikes in docking station
	ds, _ = newDockingStation(0, 1)
	bp, err = ds.releaseBike()
	assert.Nil(t, bp, "Bike should be Nil")
	assert.NotNil(t, err, "err should not be Nil")
	_, isErr = err.(error)
	assert.True(t, isErr, "Should be an error")
}

func TestDockBike(t *testing.T) {
	ds, _ := newDockingStation(0, 3)

	b := bike{isDocked: false}
	err := ds.dockBike(&b)
	// asserts that the docking station receives a bike and that the bike is now docked
	l := len(ds.bikes)
	assert.Equal(t, l, 1, "Docking Station should now have 1 bike")
	assert.True(t, b.isDocked, "Bike should now be docked")
	assert.Nil(t, err, "expected err to be nil")

	ds, _ = newDockingStation(1, 1)
	b = bike{isDocked: false}
	err = ds.dockBike(&b)
	assert.NotNil(t, err, "expected err not to be nil")
	_, isErr := err.(error)
	assert.True(t, isErr, "expected err to be instsnce of error")
}
