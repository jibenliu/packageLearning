package bellmanFord

type tBellmanFoldFinder struct {
}

func newBellmanFoldFinder() IPathFinder {
	return &tBellmanFoldFinder{}
}

func (t *tBellmanFoldFinder) FindPath(nodes []INode, lines []ILine, fromID string, toID string) (bool, []INode) {
	// 节点索引
	mapNodes := make(map[string]INode, 0)
	for _, it := range nodes {
		mapNodes[it.ID()] = it
	}

	fromNode, ok := mapNodes[fromID]
	if !ok {
		return false, nil
	}

	toNode, ok := mapNodes[toID]
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
		if it.ID() == fromID {
			it.SetWeight(0)
		} else {
			it.SetWeight(MaxWeight)
		}
	}

	// 循环更新所有节点的权重 直到不再变化
	fromNode.SetWeight(0)
	queue := newFIFOQueue()
	queue.Push(fromNode)
	for !queue.Empty() {
		ok, from := queue.Poll()
		if !ok {
			panic("unexpected !ok")
		}

		affectedLines, ok := mapFromLines[from.ID()]
		if ok {
			for _, line := range affectedLines {
				if to, ok := mapNodes[line.To()]; ok {
					if t.updateWeight(from, to, line) {
						queue.Push(to)
					}
				}
			}
		}
	}

	// 逆向查找最短路径
	if toNode.GetWeight() >= MaxWeight {
		return false, nil
	}

	queue.Clear()
	queue.Push(toNode)
	current := toNode
	maxRound := len(lines)
	for ; current != fromNode && maxRound > 0; maxRound-- {
		linkedLines, _ := mapToLines[current.ID()]
		for _, line := range linkedLines {
			from, _ := mapNodes[line.From()]
			if from.GetWeight() == current.GetWeight()-line.Weight() {
				current = from
				queue.Push(from)
			}
		}
	}

	if current != fromNode {
		return false, nil
	}

	// 返回
	result := make([]INode, queue.Size())
	for i := queue.Size() - 1; i >= 0; i-- {
		_, result[i] = queue.Poll()
	}
	return true, result
}

func (t *tBellmanFoldFinder) updateWeight(from INode, to INode, line ILine) bool {
	w := t.min(from.GetWeight()+line.Weight(), to.GetWeight())
	if to.GetWeight() > w {
		to.SetWeight(w)
		return true
	}

	return false
}

func (t *tBellmanFoldFinder) min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

var Finder = newBellmanFoldFinder()
