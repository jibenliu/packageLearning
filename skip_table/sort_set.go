package skiptable

// ISortSet 插入有序集合必须实现的接口
type ISortSet interface {
	// Key 获取这个元素唯一确定的key,用来唯一标识这个元素
	Key() string
	ISkipListNode
}

// NewDefaultSortSet 初始化一个默认的有序集合
func NewDefaultSortSet() *SortSet {
	return &SortSet{
		member: make(map[string]*SkipListNode),
		sl:     NewDefaultSkipTable(),
	}
}

// NewSortSet 初始化一个有序集合, 可以设置底层跳表的最大层数 没有特殊情况,不建议自定义层数
func NewSortSet(level int) *SortSet {
	return &SortSet{
		member: make(map[string]*SkipListNode),
		sl:     NewSkipTable(level),
	}
}

// SortSet 集合
type SortSet struct {
	//使用map记录当前集合所有的元素
	member map[string]*SkipListNode
	//底层的跳表
	sl *SkipList
}

//获取map中的元素
func (set *SortSet) getMember(key string) *SkipListNode {
	return set.member[key]
}

//向map中添加元素
func (set *SortSet) addMember(key string, member *SkipListNode) {
	set.member[key] = member
}

//删除map中的元素
func (set *SortSet) delMember(key string) {
	delete(set.member, key)
}

// Add 向sortSet中添加元素
func (set *SortSet) Add(items ...ISortSet) int {
	l := len(items)
	if l == 0 {
		return 0
	}
	//记录添加了多少个元素
	op := make(map[string]struct{})

	l--
	//想一下,为啥是从后向前遍历
	for l >= 0 {
		//如果这次已经操作过这个元素了,就不用再操作了
		if _, e := op[items[l].Key()]; e {
			l--
			continue
		}
		if member := set.getMember(items[l].Key()); member == nil {
			//如果当前集合中没有这个元素了,就添加
			node := set.sl.InsertByScore(items[l].Score(), items[l])
			set.addMember(items[l].Key(), node)
		} else {
			//如果当前集合中已经有这个元素了,就只更新分数就好了
			set.sl.UpdateScore(member, items[l].Score())
		}
		op[items[l].Key()] = struct{}{}
		l--
	}
	return len(op)
}

// Count sortSet中元素数量
func (set *SortSet) Count() int64 {
	return set.sl.Size()
}

// Rank 返回有序集合中指定成员的索引(从0开始)不存在返回 -1
func (set *SortSet) Rank(key string) int64 {
	member := set.getMember(key)
	if member == nil {
		return 0
	}
	return set.sl.GetNodeRank(member) - 1
}

// RevRank 返回有序集合中指定成员的索引(从0开始)不存在返回 -1
func (set *SortSet) RevRank(key string) int64 {
	member := set.getMember(key)
	if member == nil {
		return -1
	}
	rank := set.sl.GetNodeRank(member)
	return set.sl.Size() - rank
}

// Score 获取元素分数
func (set *SortSet) Score(key string) float64 {
	member := set.getMember(key)
	if member == nil {
		return 0
	}
	return member.score
}

// Remove 移除有序集合中的一个或多个成员
func (set *SortSet) Remove(keys ...string) int {
	for _, key := range keys {
		if member := set.getMember(key); member != nil {
			set.delMember(key)
			set.sl.Delete(member, set.sl.GetUpdateList(member))
		}
	}
	return 0
}

// RemoveRangeByRank 移除有序集合中给定的排名区间的所有成员
func (set *SortSet) RemoveRangeByRank(min, max int64) int {
	//先根据rank范围 查找node
	result := set.Range(min, max)
	if len(result) == 0 {
		return 0
	}

	//删除数据需要的各层结点信息(路径)
	//想一下,为啥只需要获取一次路径就行呢?????
	var updateList []*SkipListNode
	for _, key := range result {
		if member := set.getMember(key.Key()); member != nil {
			if updateList == nil {
				updateList = set.sl.GetUpdateList(member)
			}
			set.delMember(key.Key())
			set.sl.Delete(member, updateList)
		}
	}
	return len(result)
}

// RemoveRangeByScore 移除有序集合中给定的分数区间的所有成员
func (set *SortSet) RemoveRangeByScore(min, max float64) int {
	//先根据score范围获取node
	result := set.RangeByScore(&SkipListFindRange{
		Min:    min,
		Max:    max,
		MinInf: false,
		MaxInf: false,
	})

	if len(result) == 0 {
		return 0
	}
	//删除数据需要的各层结点信息(路径)
	//想一下,为啥只需要获取一次路径就行呢?????
	var updateList []*SkipListNode
	for _, key := range result {
		if member := set.getMember(key.Key()); member != nil {
			if updateList == nil {
				updateList = set.sl.GetUpdateList(member)
			}
			set.delMember(key.Key())
			set.sl.Delete(member, updateList)
		}
	}
	return len(result)
}

// Range 通过索引区间返回有序集合指定区间内的成员,分数从低到高
func (set *SortSet) Range(min, max int64) (result []ISortSet) {
	if set.sl.Size() == 0 {
		return
	}

	//处理范围时负数的情况
	if min < 0 {
		min = set.sl.Size() + min
	}
	if max < 0 {
		max = set.sl.Size() + max
	}

	//给定的范围出错了
	if min > max {
		return
	}
	//索引是从0开始的, 跳表中的rank是从1开始,所以这里要 +1
	nodes := set.sl.GetNodeByRank(min+1, max+1)
	if len(nodes) == 0 {
		return
	}
	result = make([]ISortSet, len(nodes))
	for i, node := range nodes {
		result[i] = node.value.(ISortSet)
	}
	return
}

// RevRange 返回有序集中指定区间内的成员，通过索引，分数从高到低排序
func (set *SortSet) RevRange(min, max int64) (result []ISortSet) {
	if set.sl.Size() == 0 {
		return
	}
	//反向查找也是按照正向查找来做的
	//只不过是把反向查找的范围转换成正向查找的范围
	//最后把查找的结果反向
	if min < 0 {
		min = -min
	} else {
		if set.sl.Size() >= min {
			min = set.sl.Size() - min
		} else {
			min = set.sl.Size()
		}
	}

	if max < 0 {
		max = -max
	} else {
		if set.sl.Size() > max {
			max = set.sl.Size() - max
		} else {
			max = 1
		}
	}

	if max > min {
		return
	}
	nodes := set.sl.GetNodeByRank(max, min)
	l := len(nodes)
	if l == 0 {
		return
	}
	result = make([]ISortSet, l)

	l--
	//翻转查找结果
	for i, node := range nodes {
		result[l-i] = node.value.(ISortSet)
	}
	return
}

// RangeByScore 返回有序集中指定分数区间内的成员，分数从低到高排序
func (set *SortSet) RangeByScore(findRange *SkipListFindRange) (result []ISortSet) {
	if findRange == nil || set.sl.Size() == 0 {
		return
	}
	nodes := set.sl.GetNodeByScore(findRange)
	if len(nodes) == 0 {
		return
	}
	result = make([]ISortSet, len(nodes))
	for i, node := range nodes {
		result[i] = node.value.(ISortSet)
	}
	return
}

// RevRangeByScore 返回有序集中指定分数区间内的成员，分数从高到低排序
func (set *SortSet) RevRangeByScore(findRange *SkipListFindRange) (result []ISortSet) {
	if findRange == nil || set.sl.Size() == 0 {
		return
	}

	//分数从高到低的查找, 本质上和从低到高查找是一样的
	//不同点是,从高到低查找时,给的范围要调换
	//最后再把查找的结果翻转一下
	findRange.Max, findRange.Min = findRange.Min, findRange.Max
	findRange.MaxInf, findRange.MinInf = findRange.MinInf, findRange.MaxInf

	nodes := set.sl.GetNodeByScore(findRange)
	l := len(nodes)
	if l == 0 {
		return
	}
	result = make([]ISortSet, l)

	l--
	//翻转查找结果
	for i, node := range nodes {
		result[l-i] = node.value.(ISortSet)
	}
	return
}
