package main

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		map[string]int{},
	}
}

type InMemoryPlayerStore struct {
	scores map[string]int
}

func (i *InMemoryPlayerStore) GetPlayerScore(id string) (int, bool) {
	score, ok := i.scores[id]
	return score, ok
}

func (i *InMemoryPlayerStore) RecordPlayerScore(id string) int {
	i.scores[id]++
	return i.scores[id]
}
