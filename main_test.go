package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDisplay(t *testing.T) {
	n := GetRandomNumber()

	assert.GreaterOrEqual(t, n, 0)

	assert.LessOrEqual(t, n, 1000)
}
