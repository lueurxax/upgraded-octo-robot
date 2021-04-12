package sync_map

import (
	"fmt"
	"sync"

	"github.com/lueurxax/upgraded-octo-robot/pkg/kv"
)

type storage struct {
	m sync.Map
}

func (s *storage) Get(key string) (interface{}, error) {
	data, ok := s.m.Load(key)
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	return data, nil
}

func (s *storage) Set(key string, value interface{}) error {
	s.m.Store(key, value)
	return nil
}

func NewKv() kv.KV {
	return &storage{}
}
