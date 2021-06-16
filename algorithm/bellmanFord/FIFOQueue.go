package bellmanFord

type tFIFOQueue struct {
	nodes    []INode
	capacity int
	rindex   int
	windex   int
}

func newFIFOQueue() iNodeQueue {
	it := &tFIFOQueue{}
	it.Clear()
	return it
}

func (t *tFIFOQueue) Clear() {
	t.nodes = make([]INode, 0)
	t.capacity = 0
	t.rindex = -1
	t.windex = -1
}

func (t *tFIFOQueue) Size() int {
	return t.windex - t.rindex
}

func (t *tFIFOQueue) Empty() bool {
	return t.Size() <= 0
}

func (t *tFIFOQueue) Push(node INode) {
	t.ensureSpace(1)
	t.windex++
	t.nodes[t.windex] = node
}

func (t *tFIFOQueue) ensureSpace(size int) {
	for t.capacity < t.windex+size+1 {
		t.nodes = append(t.nodes, nil)
		t.capacity++
	}
}

func (t *tFIFOQueue) Poll() (bool, INode) {
	if t.Empty() {
		return false, nil
	}

	t.rindex++
	it := t.nodes[t.rindex]
	t.nodes[t.rindex] = nil

	if t.rindex > t.capacity/2 {
		size := t.Size()
		offset := t.rindex + 1
		for i := 0; i < size; i++ {
			t.nodes[i], t.nodes[i+offset] = t.nodes[i+offset], nil
		}

		t.rindex -= offset
		t.windex -= offset
	}

	return true, it
}
