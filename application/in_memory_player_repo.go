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

func (i *InMemoryPlayerStore) GetPlayerScore(id string) (*PlayerScore, bool) {
	i.lock.RLock()
	defer i.lock.RUnlock()
	score, ok := i.scores[id]
	if ok {
		return &PlayerScore{id, score}, ok
	} else {
		return nil, ok
	}

}

func (i *InMemoryPlayerStore) ListPlayerScores() []*PlayerScore {
	playerScores := []*PlayerScore{}
	for playerId, score := range i.scores {
		playerScores = append(playerScores, &PlayerScore{playerId, score})
	}

	return playerScores
}

func (i *InMemoryPlayerStore) RecordPlayerScore(id string) *PlayerScore {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.scores[id]++
	return &PlayerScore{id, i.scores[id]}
}
