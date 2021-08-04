package dijkstra

type tDijkstraPathFinder struct {
}

func (t tDijkstraPathFinder) FindPath(nodes []INode, lines []ILine, srcID string, dstID string) (bool, []INode) {
	// 节点索引
	mapNodes := make(map[string]INode, 0)
	for _, it := range nodes {
		mapNodes[it.ID()] = it
	}

	srcNode, ok := mapNodes[srcID]
	if !ok {
		return false, nil
	}

	dstNode, ok := mapNodes[dstID]
	if !ok {
		return false, nil
	}

	// 边的索引
	mapFromLines := make(map[string][]ILine, 0)
	mapToLines := make(map[string][]ILine, 0)
	for _, it := range lines {
		if v, ok := mapFromLines[it.From()]; ok {
			mapFromLines[it.From()] = append(v, it)
		} else {
			mapFromLines[it.From()] = []ILine{it}
		}

		if v, ok := mapToLines[it.To()]; ok {
			mapToLines[it.To()] = append(v, it)
		} else {
			mapToLines[it.To()] = []ILine{it}
		}
	}

	// 设置from节点的weight为0, 其他节点的weight为MaxWeight
	for _, it := range nodes {
		if it.ID() == srcID {
			it.SetWeight(0)
		} else {
			it.SetWeight(MaxWeight)
		}
	}

	// 将起点push到堆
	heap := NewArrayHeap(NewNodeWeightComparator())
	heap.Push(srcNode)

	// 遍历候选节点
	for heap.IsNotEmpty() {
		_, top := heap.Pop()
		from := top.(INode)
		if from.ID() == dstID {
			break
		}

		links, ok := mapFromLines[from.ID()]
		if ok {
			for _, line := range links {
				if to, ok := mapNodes[line.To()]; ok {
					if t.updateWeight(from, to, line) {
						n := heap.IndexOf(to)
						if n >= 0 {
							heap.ShiftUp(n)
						} else {
							heap.Push(to)
						}
					}
				}
			}
		}
	}

	// 逆向查找最短路径
	if dstNode.GetWeight() >= MaxWeight {
		return false, nil
	}

	path := []INode{dstNode}
	current := dstNode
	maxRound := len(lines)
	for ; current != srcNode && maxRound > 0; maxRound-- {
		linkedLines, _ := mapToLines[current.ID()]
		for _, line := range linkedLines {
			from, _ := mapNodes[line.From()]
			if from.GetWeight() == current.GetWeight()-line.Weight() {
				current = from
				path = append(path, from)
			}
		}
	}

	if current != srcNode {
		return false, nil
	}

	t.reverse(path)
	return true, path
}

func NewDijkstraPathFinder() IPathFinder {
	return &tDijkstraPathFinder{}
}

func (t *tDijkstraPathFinder) reverse(nodes []INode) {
	for i, j := 0, len(nodes)-1; i < j; i, j = i+1, j-1 {
		nodes[i], nodes[j] = nodes[j], nodes[i]
	}
}

func (t *tDijkstraPathFinder) updateWeight(from INode, to INode, line ILine) bool {
	w := t.min(from.GetWeight()+line.Weight(), to.GetWeight())
	if to.GetWeight() > w {
		to.SetWeight(w)
		return true
	}

	return false
}

func (t *tDijkstraPathFinder) min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

var Finder = NewDijkstraPathFinder()
