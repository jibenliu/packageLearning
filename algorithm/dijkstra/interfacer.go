package dijkstra

type INode interface {
	ID() string
	GetWeight() int
	SetWeight(int)
}

const MaxWeight = int(0x7fffffff_ffffffff)

type ILine interface {
	From() string
	To() string
	Weight() int
}

type IPathFinder interface {
	FindPath(nodes []INode, lines []ILine, from string, to string) (bool, []INode)
}

type IComparator interface {
	Less(a interface{}, b interface{}) bool
}

type IHeap interface {
	Size() int
	IsEmpty() bool
	IsNotEmpty() bool

	Push(node interface{})
	Pop() (bool, interface{})

	IndexOf(node interface{}) int
	ShiftUp(i int)
}
