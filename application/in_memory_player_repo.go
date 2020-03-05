package main

import "sync"

func NewInMemoryPlayerRepo() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		map[string]*PlayerScore{},
		sync.RWMutex{},
	}
}

type InMemoryPlayerStore struct {
	scores map[string]*PlayerScore
	// A mutex is used to synchronize read/write access to the map
	lock sync.RWMutex
}

func (i *InMemoryPlayerStore) GetPlayerScore(id string) (*PlayerScore, bool) {
	i.lock.RLock()
	defer i.lock.RUnlock()
	playerScore, ok := i.scores[id]
	if ok {
		return playerScore, ok
	}
	return nil, ok

}

func (i *InMemoryPlayerStore) ListPlayerScores() []*PlayerScore {
	playerScores := []*PlayerScore{}
	for _, playerScore := range i.scores {
		playerScores = append(playerScores, playerScore)
	}
	return playerScores
}

func (i *InMemoryPlayerStore) RecordPlayerScore(id string) *PlayerScore {
	i.lock.Lock()
	defer i.lock.Unlock()
	playerScore, ok := i.scores[id]
	if ok {
		playerScore.Score++
		return playerScore
	}
	newPlayerScore := &PlayerScore{id, 1}
	i.scores[id] = newPlayerScore
	return newPlayerScore
}
