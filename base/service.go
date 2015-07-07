package base

import (
	"fmt"
	"sync/atomic"
)

type Service interface {
	Name() string

	Started() bool

	Start() error

	Stop() error
}

type ServiceSlots struct {
	OnStart   Slot
	PreServe  Slot
	PostServe Slot
	OnStop    Slot
}

type ServiceStatus int32

const (
	ServiceUnused ServiceStatus = iota
	ServiceStarted
	ServiceStopped
)

var (
	serviceStatusNames = []string{"unused", "started", "stopped"}
)

func (s ServiceStatus) String() string {
	return serviceStatusNames[int(s)]
}

type BaseService struct {
	name  string
	state int32
	Serve func() error
	Slots ServiceSlots
}

func NewBaseService(name string) *BaseService {
	return &BaseService{name: name}
}

func (s *BaseService) Name() string { return s.name }

func (s *BaseService) Started() bool {
	return atomic.LoadInt32((*int32)(&s.state)) == int32(ServiceStarted)
}

func (s *BaseService) Start() error {
	if !atomic.CompareAndSwapInt32((*int32)(&s.state), int32(ServiceUnused), int32(ServiceStarted)) {
		return fmt.Errorf("service `%s` %s", s.Name(), s.state)
	}

	s.Slots.OnStart.Publish(Service(s))

	go func() {
		if s.Serve != nil {
			s.Slots.PreServe.Publish(Service(s))

			if err := s.Serve(); err != nil {

			}

			s.Slots.PostServe.Publish(Service(s))
		}
	}()

	return nil
}

func (s *BaseService) Stop() error {
	if !atomic.CompareAndSwapInt32((*int32)(&s.state), int32(ServiceStarted), int32(ServiceStopped)) {
		return fmt.Errorf("service `%s` %s", s.Name(), s.state)
	}

	s.Slots.OnStop.Publish(Service(s))

	return nil
}
