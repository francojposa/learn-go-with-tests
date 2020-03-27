package players

import (
	"sort"
	"sync"
)

func NewInMemoryPlayerRepo() *InMemoryPlayerRepo {
	return &InMemoryPlayerRepo{
		map[string]PlayerScore{},
		sync.RWMutex{},
	}
}

type InMemoryPlayerRepo struct {
	scores map[string]PlayerScore
	// A mutex is used to synchronize read/write access to the map
	lock sync.RWMutex
}

func (r *InMemoryPlayerRepo) GetPlayerScore(id string) (PlayerScore, bool) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	playerScore, ok := r.scores[id]
	return playerScore, ok

}

func (r *InMemoryPlayerRepo) ListPlayerScores() []PlayerScore {
	ids := make([]string, len(r.scores))
	i := 0
	for id := range r.scores {
		ids[i] = id
		i++
	}

	sort.Strings(ids)

	playerScores := make([]PlayerScore, len(r.scores))
	for i, id := range ids {
		playerScores[i] = r.scores[id]
	}
	return playerScores
}

func (r *InMemoryPlayerRepo) RecordPlayerScore(id string) PlayerScore {
	r.lock.Lock()
	defer r.lock.Unlock()
	playerScore, ok := r.scores[id]
	if ok {
		playerScore.Score++
		r.scores[id] = playerScore
		return playerScore
	}
	newPlayerScore := PlayerScore{id, 1}
	r.scores[id] = newPlayerScore
	return newPlayerScore
}
