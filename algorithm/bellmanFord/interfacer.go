package bellmanFord

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

type iNodeQueue interface {
	Clear()
	Size() int
	Empty() bool
	Push(node INode)
	Poll() (bool, INode)
}
