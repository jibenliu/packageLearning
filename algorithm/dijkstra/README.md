# 求图的最短路径 
### refer:[https://studygolang.com/articles/33624](https://studygolang.com/articles/33624)

##流程
```
1.给定若干顶点, 以及顶点间的若干条边, 寻找从指定起点srcNode到指定终点dstNode的最小权重路径
2.设定srcNode的权重为0, 其他顶点的权重为无穷大
3.将srcNode节点送入候选堆
4.for 候选堆不为空:
    1.从候选堆pop顶点node
    2.如果node.id == dstNode.id, 循环结束
    3.遍历从node出发的所有边, 将边的终点to的权重, 更新为min(终点权重, node.权重+边.权重)
    4.如果to.权重 > node.权重+边.权重, 说明更新有效
    5.如果更新有效, 判断to是否在堆中, 如果是, 则上浮以维护堆秩序, 否则, 将to节点push入候选堆
5.判断终点的权重是否被更新(!=无穷大), 如果是则说明存在最短路径
6.反向查找最短路径:
    1.设定当前节点current = 终点
    2.push节点current进路径队列
    3.遍历终点为current的边, 查找符合条件的node:边的起点.权重 = current.权重-边.权重
    4.push节点node进路径队列
    5.循环1-4, 直到current == srcNode, 查找完成
```
## 设计
```
    INode: 顶点接口
    ILine: 边接口
    IPathFinder: 最短路径查找算法接口
    IComparator: 顶点比较接口
    IHeap: 顶点堆接口
    tNode: 顶点, 实现INode
    tLine: 边, 实现ILine
    tNodeWeightComparator: 基于权重的顶点比较器, 实现IComparator接口
    tArrayHeap: 堆的实现
    tDijkstraPathFinder: 狄克斯特拉算法的实现
```