package main

/**
现在你总共有 n 门课需要选，记为 0 到 n-1。
在选修某些课程之前需要一些先修课程。
例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们: [0,1]
给定课程总量以及它们的先决条件，判断是否可能完成所有课程的学习？
*/

func canFinish(numCourses int, prerequisites [][]int) bool {
	inverseAdj := make([][]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		inverseAdj[prerequisites[i][1]] = append(inverseAdj[prerequisites[i][1]], prerequisites[i][0])
	}
	/* # 深度优先遍历，判断结点是否访问过
	   # 这里要设置 3 个状态
	   # 0 就对应 False ，表示结点没有访问过
	   # 1 就对应 True ，表示结点已经访问过，在深度优先遍历结束以后才置为 1
	   # 2 表示当前正在遍历的结点，如果在深度优先遍历的过程中，
	   # 有遇到状态为 2 的结点，就表示这个图中存在环
	*/
	nodes := make([]int, numCourses)
	for i := 0; i < numCourses; i++ {
		//在遍历的过程中，如果发现有环，就退出
		if DFS(i, inverseAdj, nodes) {
			return false
		}
	}
	return true
}

func DFS(i int, inverseAdj [][]int, nodes []int) bool {
	/*
	   注意：这个递归方法的返回值是返回是否有环
	      :param i: 结点的索引
	      :param inverse_adj: 逆邻接表，记录的是当前结点的前驱结点的集合
	      :param nodes: 记录了结点是否被访问过，2 表示当前正在 DFS 这个结点
	      :return: 是否有环
	*/
	if nodes[i] == 2 {
		// 2 表示这个结点正在访问,说明有环
		return true
	}
	if nodes[i] == 1 {
		return false
	}
	nodes[i] = 2
	for _, precursor := range inverseAdj[i] {
		// 如果有环，就返回 True 表示有环
		if DFS(precursor, inverseAdj, nodes) {
			return true
		}
	}
	// # 1 表示访问结束
	nodes[i] = 1
	return false
}

/**
DFS解题思路： 深度优先搜索
1，将边缘列表转换成逆邻接矩阵的形式，
inverse_adj[i] 的slice表示，i的所有前缀节点
2，题目可以抽象为判断有向图是否可以拓扑排序（是否有环）
3，循环从每一个顶点开始深度优先遍历
A，当前节点标记为2（正在遍历）
B，如果该节点没有前缀节点（入度为0，则标记为1）
C，如果该节点有前缀节点，深度优先遍历
D，如果该节点的所有前缀节点都标记为1，则该节点标记为1
E，如果前缀节点中有正在遍历的节点（2），说明有环，返回。
*/

func findOrder(numCourses int, prerequisites [][]int) []int {
	inverseAdj := make([][]int, numCourses)
	outDegree := make([]int, numCourses) //入度
	for i := 0; i < len(prerequisites); i++ {
		//将边缘列表转换成逆邻接矩阵的形式
		outDegree[prerequisites[i][0]]++
		inverseAdj[prerequisites[i][1]] = append(inverseAdj[prerequisites[i][1]], prerequisites[i][0])
	}
	r := BFS(inverseAdj, outDegree)
	if len(r) == numCourses {
		return r
	}
	return nil
}

func BFS(inverseAdj [][]int, outDegree []int) (r []int) {
	var q queue
	for i := 0; i < len(outDegree); i++ {
		if outDegree[i] == 0 {
			//入度为0，可以作为终点
			q.push(i)
		}
	}

	for !q.empty() {
		top := q.pop()
		r = append([]int{top}, r...)
		for _, precursor := range inverseAdj[top] {
			//将当前节点移除，所有前驱节点的出度减1
			outDegree[precursor]--
			if outDegree[precursor] == 0 {
				q.push(precursor)
			}
		}
	}
	return r
}

type queue struct {
	data []int
}

func (q *queue) empty() bool {
	return len(q.data) == 0
}

func (q *queue) push(i int) {
	q.data = append(q.data, i)
}

func (q *queue) pop() int {
	r := q.data[len(q.data)-1]
	q.data = q.data[:len(q.data)-1]
	return r
}

/**
BFS解题思路 广度优先搜索
解题思路：
对课程排序是，前一篇的递进，有向图的top排序，采用广度优先搜索（BFS）
首先将边缘列表转化成逆邻接矩阵，并记录每个前缀课程的入度
入度为0 的课程没有依赖，可以先上，放入队列
一次从队列中取节点
A. 放入返回数据
B. 将依赖此节点的所有邻接节点的入度减一（删除此节点后，邻接节点的依赖减少）
C. 将修正后入度为0 的节点放入队列
D. 循环直至队列为空
返回数据如果长度等于课程长度，说明没有环，否则有环
*/
