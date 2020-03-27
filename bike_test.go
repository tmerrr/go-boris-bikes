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

func TestSetIsWorking(t *testing.T) {
	b := bike{}
	assert.False(t, b.isWorking, "expected isWorking to be false")
	b.SetIsWorking(true)
	assert.True(t, b.isWorking, "expected isWorking to be true")
}

func TestSetIsDocked(t *testing.T) {
	b := bike{}
	assert.False(t, b.isDocked, "expected isDocked to be false")
	b.SetIsDocked(true)
	assert.True(t, b.isDocked, "expected isDocked to be true")
}
