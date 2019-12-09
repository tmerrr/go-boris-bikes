package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReleaseBike(t *testing.T) {
	ds := dockingStation{}

	var b interface{} = ds.releaseBike()

	_, isBike := b.(bike)
	assert.True(t, isBike, "Should be an instance of bike")
}

func TestBikeIsWorking(t *testing.T) {
	b := bike{}
	assert.False(t, b.isWorking, "Bike isWorking should be false by default")

	b = bike{true}
	assert.True(t, b.isWorking, "Bike isWorking should be true")
}
