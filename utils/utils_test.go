package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncrBy1(t *testing.T) {

	input := 5
	expected := 6

	u := &Utils{}

	actual, _ := u.IncrBy1(input)
	assert.EqualValues(t, expected, actual)

}

func TestIncrBy2(t *testing.T) {

	input := -1
	expected := 0

	u := &Utils{}
	actual, _ := u.IncrBy1(input)

	assert.EqualValues(t, expected, actual)

}