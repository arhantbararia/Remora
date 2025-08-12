package store

import (
	"remora/pkg/resp"
	"sync"
)

type Store struct {
	mu   sync.RWMutex
	data map[string]resp.Value
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]resp.Value),
	}
}

func (s *Store) Set(key string, value resp.Value) (oldValue resp.Value, existed bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	oldValue, existed = s.data[key]
	s.data[key] = value
	return
}

func (s *Store) Get(key string) (value resp.Value, existed bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	value, existed = s.data[key]
	return
}
