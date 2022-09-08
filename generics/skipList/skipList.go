package skipList

import (
	"math/rand"
	"skipList/constraints"
)

type Record[K constraints.Ordered, V any] struct {
	Key   K
	Value V
}

type SkipNode[K constraints.Ordered, V any] struct {
	record  *Record[K, V]
	forward []*SkipNode[K, V]
}

func NewRecord[K constraints.Ordered, V any](key K, value V) *Record[K, V] {
	return &Record[K, V]{
		Key:   key,
		Value: value,
	}
}

func NewSkipNode[K constraints.Ordered, V any](key K, value V, level int) *SkipNode[K, V] {
	return &SkipNode[K, V]{
		record:  NewRecord(key, value),
		forward: make([]*SkipNode[K, V], level+1),
	}
}

type SkipList[K constraints.Ordered, V any] struct {
	head  *SkipNode[K, V]
	level int
	size  int
}

func NewSkipList[K constraints.Ordered, V any]() *SkipList[K, V] {
	return &SkipList[K, V]{
		head:  NewSkipNode[K, V](*new(K), *new(V), 0),
		level: -1,
		size:  0,
	}
}

func (s *SkipList[K, V]) Find(key K) (V, bool) {
	x := s.head

	for i := s.level; i >= 0; i-- {

		for {
			if x.forward[i] == nil || x.forward[i].record.Key > key {
				break
			} else if x.forward[i].record.Key == key {
				return x.forward[i].record.Value, true
			} else {
				x = x.forward[i]
			}
		}
	}

	return *new(V), false
}

func (s *SkipList[K, V]) getRandomLevel() int {
	level := 0
	for rand.Int31()%2 == 0 {
		level += 1
	}
	return level
}

func (s *SkipList[K, V]) adjustLevel(level int) {
	temp := s.head.forward

	s.head = NewSkipNode(*new(K), *new(V), level)
	s.level = level

	copy(s.head.forward, temp)
}

func (s *SkipList[K, V]) Insert(key K, value V) {
	newLevel := s.getRandomLevel()

	if newLevel > s.level {
		s.adjustLevel(newLevel)
	}

	newNode := NewSkipNode[K, V](key, value, newLevel)
	updates := make([]*SkipNode[K, V], newLevel+1)
	x := s.head

	for i := s.level; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].record.Key < key {
			x = x.forward[i]
		}
		updates[i] = x
	}

	for i := 0; i <= newLevel; i++ {
		newNode.forward[i] = updates[i].forward[i]
		updates[i].forward[i] = newNode
	}

	s.size += 1
}

func (s *SkipList[K, V]) Delete(key K) {
	x := s.head
	for i := s.level; i >= 0; i-- {
		for {
			if x.forward[i] == nil || x.forward[i].record.Key > key {
				break
			} else if x.forward[i].record.Key == key {
				x.forward[i] = x.forward[i].forward[i]
			} else {
				x = x.forward[i]
			}
		}
	}
}
