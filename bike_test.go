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
	assert.False(t, b.isDocked, "Bike isDocked should be true")
}
