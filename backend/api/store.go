package api

import "fmt"

type BoardStore interface {
	GetAll() ([]*Board, error)
	GetByID(id string) (*Board, error)
}

type MemoryBoardStore struct {
	boards map[string]*Board
}

func NewMemoryBoardStore(boards []*Board) *MemoryBoardStore {
	m := &MemoryBoardStore{boards: make(map[string]*Board)}
	for _, b := range boards {
		m.boards[b.ID] = b
	}
	return m
}

func (s *MemoryBoardStore) GetAll() ([]*Board, error) {
	result := make([]*Board, 0, len(s.boards))
	for _, b := range s.boards {
		result = append(result, b)
	}
	return result, nil
}

func (s *MemoryBoardStore) GetByID(id string) (*Board, error) {
	b, ok := s.boards[id]
	if !ok {
		return nil, fmt.Errorf("board %q not found", id)
	}
	return b, nil
}
