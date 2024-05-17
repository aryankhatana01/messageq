package main

import (
	"fmt"
	"sync"
)

type Storer interface {
	Push([]byte) (int, error)
	Fetch(int) ([]byte, error)
}

type Storage struct {
	mu   sync.RWMutex
	data [][]byte
}

func NewMemoryStorage() *Storage {
	return &Storage{
		data: make([][]byte, 0),
	}
}

func (s *Storage) Push(d []byte) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data = append(s.data, d)
	return len(s.data) - 1, nil
}

func (s *Storage) Fetch(offset int) ([]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if offset >= len(s.data) {
		return nil, fmt.Errorf("index out of range")
	}
	return s.data[offset], nil
}
