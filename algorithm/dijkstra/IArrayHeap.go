package dijkstra

import (
	"errors"
	"fmt"
	"strings"
)

type tArrayHeap struct {
	comparator IComparator
	items      []interface{}
	size       int
	version    int64
}

func (t *tArrayHeap) Size() int {
	return t.size
}

func (t *tArrayHeap) IsEmpty() bool {
	return t.size <= 0
}

func (t *tArrayHeap) IsNotEmpty() bool {
	return !t.IsEmpty()
}

func (t *tArrayHeap) Push(value interface{}) {
	t.version++
	t.ensureSize(t.size + 1)
	t.items[t.size] = value
	t.size++

	t.ShiftUp(t.size - 1)
	t.version++
}

func (t *tArrayHeap) ensureSize(size int) {
	for len(t.items) < size {
		t.items = append(t.items, nil)
	}
}

func (t *tArrayHeap) parentOf(i int) int {
	return (i - 1) / 2
}

func (t *tArrayHeap) leftChildOf(i int) int {
	return i*2 + 1
}

func (t *tArrayHeap) rightChildOf(i int) int {
	return t.leftChildOf(i) + 1
}

func (t *tArrayHeap) last() (i int, v interface{}) {
	if t.IsEmpty() {
		return -1, nil
	}
	i = t.size - 1
	v = t.items[i]
	return i, v
}

func (t *tArrayHeap) Pop() (b bool, i interface{}) {
	if t.IsEmpty() {
		return false, nil
	}
	t.version++
	top := t.items[0]
	li, lv := t.last()
	t.items[0] = nil
	t.size--
	if t.IsEmpty() {
		return true, top
	}
	t.items[0] = lv
	t.items[li] = nil
	t.ShiftDown(0)
	t.version++
	return true, top
}

func (t *tArrayHeap) IndexOf(node interface{}) int {
	n := -1
	for i, it := range t.items {
		if it == node {
			n = i
			break
		}
	}
	return n
}

func (t *tArrayHeap) ShiftUp(i int) {
	if i <= 0 {
		return
	}
	v := t.items[i]
	pi := t.parentOf(i)
	pv := t.items[pi]

	if t.comparator.Less(v, pv) {
		t.items[pi], t.items[i] = v, pv
		t.ShiftUp(pi)
	}
}

func (t *tArrayHeap) ShiftDown(i int) {
	pv := t.items[i]
	ok, ci, cv := t.minChildOf(i)
	if ok && t.comparator.Less(cv, pv) {
		t.items[i], t.items[ci] = cv, pv
		t.ShiftDown(ci)
	}
}
func (t *tArrayHeap) minChildOf(p int) (ok bool, i int, v interface{}) {
	li := t.leftChildOf(p)
	if li >= t.size {
		return false, 0, nil
	}
	lv := t.items[li]

	ri := t.rightChildOf(p)
	if ri >= t.size {
		return true, li, lv
	}
	rv := t.items[ri]

	if t.comparator.Less(lv, rv) {
		return true, li, lv
	} else {
		return true, ri, rv
	}
}

func (t *tArrayHeap) String() string {
	level := 0
	lines := make([]string, 0)
	lines = append(lines, "")

	for {
		n := 1 << level
		min := n - 1
		max := n + min - 1
		if min >= t.size {
			break
		}

		line := make([]string, 0)
		for i := min; i <= max; i++ {
			if i >= t.size {
				break
			}
			line = append(line, fmt.Sprintf("%4d", t.items[i]))
		}
		lines = append(lines, strings.Join(line, ","))

		level++
	}

	return strings.Join(lines, "\n")
}

func NewArrayHeap(comparator IComparator) IHeap {
	return &tArrayHeap{
		comparator: comparator,
		items:      make([]interface{}, 0),
		size:       0,
		version:    0,
	}
}

var gNoMoreElementsError = errors.New("no more elements")
