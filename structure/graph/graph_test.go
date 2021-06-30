package graph

import (
	"fmt"
	"testing"
)

var g ItemGraph

func fillGraph() {
	nA := GraphNode{"A"}
	nB := GraphNode{"B"}
	nC := GraphNode{"C"}
	nD := GraphNode{"D"}
	nE := GraphNode{"E"}
	nF := GraphNode{"F"}
	g.AddNode(&nA)
	g.AddNode(&nB)
	g.AddNode(&nC)
	g.AddNode(&nD)
	g.AddNode(&nE)
	g.AddNode(&nF)

	g.AddEdge(&nA, &nB)
	g.AddEdge(&nA, &nC)
	g.AddEdge(&nB, &nE)
	g.AddEdge(&nC, &nE)
	g.AddEdge(&nE, &nF)
	g.AddEdge(&nD, &nA)
}

func TestAddNode(t *testing.T) {
	fillGraph()
	g.String()
}

func TestTraverse(t *testing.T) {
	g.Traverse(func(n *GraphNode) {
		fmt.Printf("%v\n", n)
	})
}
