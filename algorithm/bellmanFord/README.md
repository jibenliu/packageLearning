## 最短路径 贝尔曼-福特算法

## refer:[https://studygolang.com/articles/33601](https://studygolang.com/articles/33601)

### 流程

```
1.给定若干顶点, 以及顶点间的若干条边, 寻找从指定起点from到指定终点to的最小权重路径
2.设定from的权重为0, 其他顶点的权重为无穷大
3.将from节点送入候选队列
4.for 候选队列不为空:
    1.从候选队列出队一个顶点node
    2.遍历从node出发的所有边, 将边的终点权重, 更新为min(终点权重, node.权重+边.权重)
    3.如果终点权重 > node.权重+边.权重, 说明更新有效, 则将终点push到候选队列
5.判断终点的权重是否被更新(!=无穷大), 如果是则说明存在最短路径
6.反向查找最短路径:
    1.设定当前节点current = 终点
    2.push节点current进路径队列
    3.遍历终点为current的边, 查找符合条件的node:边的起点.权重 = current.权重-边.权重
    4.push节点node进路径队列
    5.循环1-4, 直到current == from, 查找完成
```

### 设计
```
INode: 顶点接口
ILine: 边接口
IPathFinder: 最短路径查找算法接口
iNodeQueue: 顶点队列接口, FIFO队列
tNode: 顶点, 实现INode
tLine: 边, 实现ILine
tFIFOQueue: FIFO节点队列的实现
tBellmanFoldFinder: 贝尔曼-福特算法的实现
```