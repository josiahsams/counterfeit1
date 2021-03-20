package service

import (
	"testing"

	uf "github.com/josiahsams/demo1/utils/fakes"
	"github.com/stretchr/testify/assert"
)

// type UtilsMock struct {}

// var Fn1 func(int)(int, error)

// func (m *UtilsMock)IncrBy1(x int)(int, error) {
// 	return Fn1(x)
// }

func TestMethod1(t *testing.T) {

	input := -1
	expected := 20

	// Fn1 = func(x int) (int, error) {
	// 	return 10, nil //errors.New("some")
	// }

	u := &uf.Utils{}
	sv := NewService(u)

	u.IncrBy1Returns(10, nil)
	u.IncrBy1ReturnsOnCall(input, 10, nil)

	actual, _ := sv.Method1(input)

	assert.EqualValues(t, expected, actual)

}

func TestMethod2(t *testing.T) {

	input := 10
	expected := 40

	// Fn1 = func(x int) (int, error) {
	// 	return 10, nil //errors.New("some")
	// }

	u := &uf.Utils{}
	sv := NewService(u)

	// u.IncrBy1Returns(10, nil)
	u.IncrBy1ReturnsOnCall(0, 20, nil)
	u.IncrBy1ReturnsOnCall(1, 40, nil)

	actual, _ := sv.Method1(input)
	assert.EqualValues(t, expected, actual)
	
	expected = 80
	actual2, _ := sv.Method1(input)
	assert.EqualValues(t, expected, actual2)


	expected = 2
	assert.EqualValues(t, expected, u.IncrBy1CallCount())

}