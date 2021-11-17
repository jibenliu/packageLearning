package skiptable

import (
	"math/rand"
	"time"
)

const (
	// DefaultMaxLevel 默认最大索引层数
	DefaultMaxLevel = 32
	// MinLevel 自定义层数时,最新的层数
	MinLevel = 16
)

// SkipListP 跳表加一层索引的概率
var SkipListP = 0.25

// SkipListFindRange 根据scores查找元素的条件
type SkipListFindRange struct {
	Min, Max       float64 //最大值和最小值
	MinInf, MaxInf bool    //是否是正无穷和负无穷
}

// ISkipListNode 跳表中元素必须实现的接口
type ISkipListNode interface {
	// Score 获取这个元素的分数
	Score() float64

	// Compare 根据返回值判断这两个值的大小
	//-1 this < value
	//0 this == value
	//1 this > value
	Compare(value ISkipListNode) int
}

// SkipList 跳表
type SkipList struct {
	//头结点和尾结点
	//重点,头结点是一个真实存在的,尾结点只是一个指针
	head, tail *SkipListNode

	size int64 //node总数

	level int //当前跳表的最高level

	maxLevel int //当前最大层数
}

// NewDefaultSkipTable 初始化一个默认的跳表
func NewDefaultSkipTable() *SkipList {
	rand.Seed(time.Now().UnixNano())
	return &SkipList{
		head:     NewSkipListNode(DefaultMaxLevel, 0, nil),
		size:     0,
		level:    1,
		maxLevel: DefaultMaxLevel,
	}
}

// NewSkipTable 初始化一个自定义最大层数的跳表
func NewSkipTable(maxLevel int) *SkipList {
	if maxLevel < MinLevel {
		maxLevel = MinLevel
	}
	rand.Seed(time.Now().UnixNano())
	return &SkipList{
		head:     NewSkipListNode(maxLevel, 0, nil),
		size:     0,
		level:    1,
		maxLevel: maxLevel,
	}
}

//随机索引的层数
func (list *SkipList) randLevel() int {
	level := 1
	for (rand.Uint32()&0xFFFF) < uint32(0xFFFF*SkipListP) && level < list.maxLevel {
		level++
	}
	return level
}

// Size 条表中的结点数量
func (list *SkipList) Size() int64 {
	return list.size
}

// InsertByScore 插入一个结点
func (list *SkipList) InsertByScore(score float64, value ISkipListNode) *SkipListNode {
	rank := make([]int64, list.maxLevel)
	update := make([]*SkipListNode, list.maxLevel)
	t := list.head
	for i := list.level - 1; i >= 0; i-- {
		if i == list.level-1 {
			rank[i] = 0
		} else {
			rank[i] = rank[i+1]
		}
		//当前层的下一个结点存在 && (下一个结点score<score || 当score相同时,比较这两个结点,下一个结点<新插入的结点)
		for t.Next(i) != nil && (t.Next(i).score < score || (t.Next(i).score == value.Score() && t.Next(i).value.Compare(value) < 0)) {
			rank[i] += t.level[i].span
			t = t.Next(i)
		}
		update[i] = t
	}

	level := list.randLevel()

	if level > list.level {
		//处理rand level后, level>当前level后的情况
		for i := list.level; i < level; i++ {
			rank[i] = 0
			update[i] = list.head
			update[i].SetSpan(i, list.size)
		}
		list.level = level
	}
	newNode := NewSkipListNode(level, score, value)

	for i := 0; i < level; i++ {
		newNode.SetNext(i, update[i].Next(i))
		update[i].SetNext(i, newNode)

		newNode.SetSpan(i, update[i].Span(i)-(rank[0]-rank[i]))
		update[i].SetSpan(i, rank[0]-rank[i]+1)
	}

	//处理新增结点的span
	for i := level; i < list.level; i++ {
		update[i].level[i].span++
	}
	//处理新节点的后退指针
	if update[0] == list.head {
		newNode.backward = nil
	} else {
		newNode.backward = update[0]
	}

	//判断新插入的节点是不是最后一个节点
	if newNode.Next(0) != nil {
		newNode.Next(0).backward = newNode
	} else {
		//如果是最后一个节点,就让tail指针指向这新插入的节点
		list.tail = newNode
	}
	list.size++
	return newNode
}

// UpdateScore 更新结点的score
func (list *SkipList) UpdateScore(node *SkipListNode, score float64) {
	if score == node.score {
		return
	}
	//更新后,分数还是 < next node的位置不用变
	if score > node.score {
		if node.Next(0) != nil && score < node.Next(0).score {
			node.score = score
			return
		}
	}

	//更新后,分数还是 > per node的位置不用变
	if score < node.score {
		if node.Pre() != nil && score > node.Pre().score {
			node.score = score
			return
		}
	}
	//删掉node,重新插入
	updateList := list.GetUpdateList(node)
	list.Delete(node, updateList)
	//重新插入
	list.InsertByScore(score, node.value)
}

// GetUpdateList 获取找到该结点的各层结点(路径)
func (list *SkipList) GetUpdateList(node *SkipListNode) (update []*SkipListNode) {
	update = make([]*SkipListNode, list.maxLevel)
	t := list.head
	for i := list.level - 1; i >= 0; i-- {
		for t.Next(i) != nil && (t.Next(i).score < node.score || (t.Next(i).score == node.score && t.Next(i).value.Compare(node.value) < 0)) {
			t = t.Next(i)
		}
		update[i] = t
	}
	return
}

// Delete 删除对应的结点
func (list *SkipList) Delete(node *SkipListNode, update []*SkipListNode) {
	if node == nil {
		return
	}
	//head 不能删
	if node == list.head {
		return
	}

	for i := 0; i < list.level; i++ {
		if update[i].Next(i) == node {
			//修改span
			update[i].SetSpan(i, update[i].Span(i)+node.Span(i)-1)
			//删除对应的结点
			update[i].SetNext(i, node.Next(i))
		} else {
			update[i].level[i].span--
		}
	}

	//处理node的后指针
	if node.Next(0) == nil { //node是最后一个,把tail指针指向node的上一个(update[0])
		list.tail = update[0]
	} else { //node不是最后一个,node的下一个指向node的上一个(update[0])
		node.Next(0).backward = update[0]
	}

	//处理删掉的是最高level的情况,当前的level要对应的--
	for list.level > 1 && list.head.Next(list.level-1) == nil {
		list.level--
	}

	list.size--
}

// GetNodeByScore 根据 score 范围 查找 node
func (list *SkipList) GetNodeByScore(findRange *SkipListFindRange) (result []*SkipListNode) {
	if findRange == nil || list.Size() == 0 {
		return
	}
	//查找范围不在这跳表中,直接return
	if !list.ScoreInRange(findRange) {
		return
	}
	t := list.head
	if findRange.MinInf {
		//从头开始查找
		t = list.head.Next(0)
	} else {
		//不是从头,找到最小的那个元素
		for i := list.level - 1; i >= 0; i-- {
			for t.Next(i) != nil && t.Next(i).score < findRange.Min {
				t = t.Next(i)
			}
		}
	}
	for {
		//符合范围的条件 (从负无穷 || 当前的score >= 查找的最小值) && (到正无穷 || 当前元素 <= 查找的最大值)
		if (findRange.MinInf || t.score >= findRange.Min) && (findRange.MaxInf || t.score <= findRange.Max) {
			result = append(result, t)
		}
		if t.Next(0) == nil || (!findRange.MaxInf && t.Next(0).score > findRange.Max) {
			//下一个元素是空(到尾了) || (不是查找到正无穷 && 下一个元素的 score > 要查找的最大值)
			break
		} else {
			//向右移动
			t = t.Next(0)
		}
	}
	return
}

// GetNodeByRank 根据排名 范围 查找 node
func (list *SkipList) GetNodeByRank(left, right int64) (result []*SkipListNode) {
	//范围出错
	if list.Size() == 0 || left == 0 || right == 0 || right < left || left > list.Size() {
		return
	}
	tRank := int64(0)
	t := list.head
	result = make([]*SkipListNode, 0, right-left+1)
	//先找到排名最小的元素,然后向右一点点查找,直到找到排名最大的元素
	for i := list.level - 1; i >= 0; i-- {
		for t.Next(i) != nil && tRank+t.level[i].span <= left {
			tRank += t.level[i].span
			t = t.Next(i)
		}
		if tRank == left {
			for ; t != nil && tRank <= right; t = t.Next(0) {
				result = append(result, t)
				tRank++
			}
			return
		}
	}
	return
}

// GetNodeRank 获取这个node的排名(排名从1开始)
func (list *SkipList) GetNodeRank(node *SkipListNode) int64 {
	t := list.head
	rank := int64(0)
	for i := list.level - 1; i >= 0; i-- {
		for t.Next(i) != nil && t.Next(i).score <= node.score {
			rank += t.level[i].span
			if t.Next(i).score == node.score && t.Next(i).value.Compare(node.value) == 0 {
				return rank
			}
			t = t.Next(i)
		}
	}
	return rank
}

// ScoreInRange 判断 这个跳表 的最大值和最小值 是否包含 要查询的score范围
func (list *SkipList) ScoreInRange(findRange *SkipListFindRange) bool {
	if !findRange.MaxInf && list.head.Next(0).score > findRange.Max {
		return false
	}
	if !findRange.MinInf && list.tail.score < findRange.Min {
		return false
	}
	return true
}
