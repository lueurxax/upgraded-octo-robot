package mutexes

import (
	"fmt"
	"sync"

	"github.com/lueurxax/upgraded-octo-robot/pkg/kv"
)

type shard struct {
	mu    sync.RWMutex
	value interface{}
}

func (s *shard) Get() interface{} {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.value
}

func (s *shard) Set(value interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.value = value
}

type storage struct {
	mu sync.RWMutex
	m  map[string]*shard
}

func (s *storage) Get(key string) (interface{}, error) {
	sh, err := s.getShard(key)
	if err != nil {
		return nil, err
	}
	return sh.Get(), nil
}

func (s *storage) Set(key string, value interface{}) error {
	sh := s.getOrSetShard(key)
	sh.Set(value)
	return nil
}

func (s *storage) getShard(key string) (*shard, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	sh, ok := s.m[key]
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	return sh, nil
}

func (s *storage) getOrSetShard(key string) *shard {
	s.mu.Lock()
	defer s.mu.Unlock()
	sh, ok := s.m[key]
	if !ok {
		sh = &shard{}
		s.m[key] = sh
	}
	return sh
}

func NewStorage() kv.KV {
	return &storage{
		m: map[string]*shard{},
	}
}
