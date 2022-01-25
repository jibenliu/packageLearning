package skipList

import "math/rand"

const (
	// maxLevel 表示元素最多可以有多少个向前的指针，它决定了跳跃表的容量。p 是一个概率因子，它决定了随机生成 level 的大小分布。
	maxLevel int     = 16 // Should be enough for 2^16 elements
	p        float32 = 0.25
)

// Element is an Element of a skiplist.
type Element struct {
	Score   float64
	Value   interface{}
	forward []*Element
}

func newElement(score float64, value interface{}, level int) *Element {
	return &Element{
		Score:   score,
		Value:   value,
		forward: make([]*Element, level),
	}
}

// SkipList represents a skiplist.
// The zero value from SkipList is an empty skiplist ready to use.
type SkipList struct {
	header *Element // header is a dummy element
	len    int      // current skiplist length，header not included
	level  int      // current skiplist level，header not included
}

// New returns a new empty SkipList.
func New() *SkipList {
	return &SkipList{
		header: &Element{forward: make([]*Element, maxLevel)},
	}
}

func randomLevel() int {
	level := 1
	for rand.Float32() < p && level < maxLevel {
		level++
	}
	return level
}

// Front returns first element in the skiplist which maybe nil.
func (sl *SkipList) Front() *Element {
	return sl.header.forward[0]
}

// Len returns the element length of the skiplist
func (sl *SkipList) Len() int {
	return sl.len
}

// Next returns first element after e.
func (e *Element) Next() *Element {
	if e != nil {
		return e.forward[0]
	}
	return nil
}

// Search the skiplist to findout element with the given score.
// Returns (*Element, true) if the given score present, otherwise returns (nil, false).
func (sl *SkipList) Search(score float64) (element *Element, ok bool) {
	x := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].Score < score {
			x = x.forward[i]
		}
	}
	x = x.forward[0]
	if x != nil && x.Score == score {
		return x, true
	}
	return nil, false
}

// Insert (score, value) pair to the skiplist and returns pointer of element.
func (sl *SkipList) Insert(score float64, value interface{}) *Element {
	update := make([]*Element, maxLevel)
	x := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].Score < score {
			x = x.forward[i]
		}
		update[i] = x
	}
	x = x.forward[0]

	// Score already presents, replace with new value then return
	if x != nil && x.Score == score {
		x.Value = value
		return x
	}

	level := randomLevel()
	if level > sl.level {
		level = sl.level + 1
		update[sl.level] = sl.header
		sl.level = level
	}
	e := newElement(score, value, level)
	for i := 0; i < level; i++ {
		e.forward[i] = update[i].forward[i]
		update[i].forward[i] = e
	}
	sl.len++
	return e
}

// Delete remove and return element with given score, return nil if element not present
func (sl *SkipList) Delete(score float64) *Element {
	update := make([]*Element, maxLevel)
	x := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].Score < score {
			x = x.forward[i]
		}
		update[i] = x
	}
	x = x.forward[0]

	if x != nil && x.Score == score {
		for i := 0; i < sl.level; i++ {
			if update[i].forward[i] != x {
				return nil
			}
			update[i].forward[i] = x.forward[i]
		}
		sl.len--
	}
	return x
}
