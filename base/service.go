package base

import ()

type Service interface {
}

type BaseService struct {
	name string
}

func NewBaseService(name string) *BaseService {
	return &BaseService{name}
}
