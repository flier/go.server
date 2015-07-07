package base

import (
	"errors"
	"sync"
)

type Value interface{}

type Subscriber interface {
	Unsubscribe() error

	Chan() chan Value

	Value() Value
}

var (
	ErrSubscriberNotFound = errors.New("subscriber not found")
)

type subscriber struct {
	slot *Slot
	c    chan Value
}

func (s *subscriber) Unsubscribe() error { return s.slot.unsubscribe(s) }

func (s *subscriber) Chan() chan Value { return s.c }

func (s *subscriber) Value() Value { return <-s.c }

type Slot struct {
	sync.RWMutex
	subscribers []*subscriber
}

func (s *Slot) Subscribe() Subscriber {
	subscriber := &subscriber{s, make(chan Value)}

	s.Lock()

	s.subscribers = append(s.subscribers, subscriber)

	s.Unlock()

	return subscriber
}

func (s *Slot) unsubscribe(rhs *subscriber) error {
	s.Lock()
	defer s.Unlock()

	for i, lhs := range s.subscribers {
		if lhs == rhs {
			s.subscribers = append(s.subscribers[:i], s.subscribers[i:]...)

			return nil
		}
	}

	return ErrSubscriberNotFound
}

func (s *Slot) Publish(v Value) {
	s.RLock()

	for _, sub := range s.subscribers {
		sub.c <- v
	}

	s.RUnlock()
}

func (s *Slot) NotifyAll() {
	s.Publish(s)
}

func (s *Slot) On(callback func(Value)) {
	subscriber := s.Subscribe()

	go func() {
		defer subscriber.Unsubscribe()

		for {
			v, ok := <-subscriber.Chan()

			if !ok {
				break
			}

			callback(v)
		}
	}()
}

func (s *Slot) Once(callback func(Value)) {
	subscriber := s.Subscribe()

	go func() {
		defer subscriber.Unsubscribe()

		v, ok := <-subscriber.Chan()

		if ok {
			callback(v)
		}
	}()
}
