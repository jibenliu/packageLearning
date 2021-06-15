package dijkstra

import "fmt"

type tNode struct {
	id     string
	weight int
}

func NewNode(id string) INode {
	return &tNode{
		id,
		MaxWeight,
	}
}

func (t *tNode) ID() string {
	return t.id
}

func (t *tNode) GetWeight() int {
	return t.weight
}

func (t *tNode) SetWeight(w int) {
	t.weight = w
}

func (t *tNode) String() string {
	return fmt.Sprintf("%s(%v)", t.id, t.weight)
}
