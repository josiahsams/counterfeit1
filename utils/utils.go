package utils

import "errors"

//go:generate counterfeiter -o fakes/utils.go --fake-name Utils . UtilsInterface
type UtilsInterface interface {
	IncrBy1(int) (int, error)
}

type Utils struct {}

func (u *Utils)IncrBy1(x int) (int, error) {
	if (x < 0) {
		return 0, errors.New("value is negative")
	}
	return x + 1, nil
}