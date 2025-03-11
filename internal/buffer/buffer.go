package buffer

import (
	"errors"
	"target-management/internal/buffer/model"
)

// Buffer ...
type Buffer interface {
	Push(fact *model.Fact)
	Pop() *model.Fact
	GetLength() int
	IsEmpty() bool
	Peek() (*model.Fact, error)
}

type buffer struct {
	facts []*model.Fact
}

// NewBuffer ...
func NewBuffer(size int) Buffer {
	return &buffer{
		facts: make([]*model.Fact, 0, size),
	}
}

func (b *buffer) Push(fact *model.Fact) {
	b.facts = append(b.facts, fact)
}

func (b *buffer) Pop() *model.Fact {
	if b.IsEmpty() {
		return nil
	}
	fact := b.facts[0]
	if b.GetLength() == 1 {
		b.facts = nil
		return fact
	}
	b.facts = b.facts[1:]
	return fact
}

func (b *buffer) GetLength() int {
	return len(b.facts)
}

func (b *buffer) IsEmpty() bool {
	return b.GetLength() == 0
}

func (b *buffer) Peek() (*model.Fact, error) {
	if b.IsEmpty() {
		return nil, errors.New("empty buffer")
	}
	return b.facts[0], nil
}
