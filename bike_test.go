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
