package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBikeIsWorking(t *testing.T) {
	b := bike{}
	assert.False(t, b.isWorking, "Bike isWorking should be false by default")
	assert.False(t, b.isDocked, "Bike isDocked should be false by default")

	b = bike{true, true}
	assert.True(t, b.isWorking, "Bike isWorking should be true")
	assert.True(t, b.isDocked, "Bike isDocked should be true")
}

func TestNewBike(t *testing.T) {
	b := newBike()
	assert.True(t, b.isWorking, "Bike should be working")
	assert.True(t, b.isDocked, "Bike should be docked")
}

func TestRelease(t *testing.T) {
	b := newBike()
	assert.True(t, b.isDocked, "Bike should be docked")
	b.release()
	assert.False(t, b.isDocked, "Bike should now be released (not docked)")
}

func TestDock(t *testing.T) {
	b := bike{}
	assert.False(t, b.isDocked, "Bike should not be docked")
	b.dock()
	assert.True(t, b.isDocked, "Bike should now be docked")
}
