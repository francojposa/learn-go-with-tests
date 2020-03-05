package main

import "sync"

func NewInMemoryPlayerRepo() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		map[string]int{},
		sync.RWMutex{},
	}
}

type InMemoryPlayerStore struct {
	scores map[string]int
	// A mutex is used to synchronize read/write access to the map
	lock sync.RWMutex
}

func (i *InMemoryPlayerStore) GetPlayerScore(id string) (int, bool) {
	i.lock.RLock()
	defer i.lock.RUnlock()
	score, ok := i.scores[id]
	return score, ok
}

func (i *InMemoryPlayerStore) RecordPlayerScore(id string) int {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.scores[id]++
	return i.scores[id]
}
