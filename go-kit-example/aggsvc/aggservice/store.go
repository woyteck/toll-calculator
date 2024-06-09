package aggservice

import (
	"fmt"

	"github.com/woyteck/toll-calculator/types"
)

type MemoryStore struct {
	data map[int]float64
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: map[int]float64{},
	}
}

func (m *MemoryStore) Insert(d types.Distance) error {
	m.data[d.OBUID] += d.Value

	return nil
}

func (m *MemoryStore) Get(obuid int) (float64, error) {
	dist, ok := m.data[obuid]
	if !ok {
		return 0.0, fmt.Errorf("could not find distance for obuid %d", obuid)
	}

	return dist, nil
}
