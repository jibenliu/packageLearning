package dijkstra

import "errors"

type tNodeWeightComparator struct {
}

func (t *tNodeWeightComparator) Less(a interface{}, b interface{}) bool {
	if a == nil || b == nil {
		panic(gNullArgumentError)
	}
	n1 := a.(INode)
	n2 := b.(INode)
	return n1.GetWeight() <= n2.GetWeight()
}

func NewNodeWeightComparator() IComparator {
	return &tNodeWeightComparator{}
}

var gNullArgumentError = errors.New("null argument error")
