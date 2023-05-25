package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddition(t *testing.T) {
	assert := assert.New(t)

	var a int = 2
	var b int = 3
	var expected int = 5
	result := add(a, b)
	assert.Equal(expected, result, "should be equal")
}
