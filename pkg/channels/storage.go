package channels

import (
	"fmt"
	"sync"

	"github.com/lueurxax/upgraded-octo-robot/pkg/kv"
)

type request chan interface{}

type chnls struct {
	read  chan request
	write chan interface{}
}

type storage struct {
	mu    sync.RWMutex
	chMap map[string]chnls
}

func (s *storage) Get(key string) (interface{}, error) {
	res := make(chan interface{})
	ch, ok := s.chMap[key]
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	// possible write to the nil
	ch.read <- res
	return <-res, nil
}

func (s *storage) Set(key string, value interface{}) error {
	ch, ok := s.chMap[key]
	if !ok {
		// first write in storage
		s.mu.Lock()
		// second check in lock
		if _, ok = s.chMap[key]; !ok {
			ch = chnls{
				read:  make(chan request),
				write: make(chan interface{}),
			}
			s.chMap[key] = ch
			go s.loop(key)
		}
		s.mu.Unlock()
	}
	ch.write <- value
	return nil
}

func (s *storage) loop(key string) {
	var data interface{}
	for {
		select {
		case req := <-s.chMap[key].read:
			req <- data
		case data = <-s.chMap[key].write:
		}
	}
}

func NewStorage() kv.KV {
	return &storage{
		chMap: map[string]chnls{},
	}
}
