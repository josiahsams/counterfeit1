package service

import (
	u "github.com/josiahsams/demo1/utils"
)

type Service struct {
	obj u.UtilsInterface
}

func NewService(ob u.UtilsInterface) *Service {
	return &Service {ob}
}

func (s *Service)Method1(x int) (int, error) {
	v, err := s.obj.IncrBy1(x)
	if err == nil {
		return v * 2, nil
	} else {
		return 0, err
	}
}