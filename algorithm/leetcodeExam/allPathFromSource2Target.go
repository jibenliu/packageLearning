package main

/**
给你一个有 n 个节点的 有向无环图（DAG），请你找出所有从节点 0 到节点 n-1 的路径并输出（不要求按特定顺序）

二维数组的第 i 个数组中的单元都表示有向图中 i 号节点所能到达的下一些节点，空就是没有下一个结点了。

译者注：有向图是有方向的，即规定了 a→b 你就不能从 b→a 。

示例 1：

图片
输入：graph = [[1,2],[3],[3],[]]
输出：[[0,1,3],[0,2,3]]
解释：有两条路径 0 -> 1 -> 3 和 0 -> 2 -> 3
*/
// 797. 所有可能的路径：中等
// tags：图、深度优先搜索、回溯

var target int    //搜索目标
var dag [][]int   //图的领接矩阵
var paths [][]int //存储所有路径
var path []int    //存储当前搜索到的路径

// DFS 深度优先搜索
func DFS(cur int) {
	path = append(path, cur)
	defer func() {
		path = path[:len(path)-1]
	}()
	if cur == target {
		ans := make([]int, len(path))
		copy(ans, path)
		paths = append(paths, ans)
		return
	}
	for _, next := range dag[cur] {
		DFS(next)
	}
}

func allPathsSourceTarget(graph [][]int) [][]int {
	target = len(graph) - 1
	dag = graph
	paths = make([][]int, 0)
	path = make([]int, 0)
	DFS(0)
	return paths
}
