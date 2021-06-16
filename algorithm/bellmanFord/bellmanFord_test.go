package bellmanFord

import (
	"fmt"
	"strings"
	"testing"
)

func Test_BellmanFold(t *testing.T) {
	fnAssertTrue := func(b bool, msg string) {
		if !b {
			t.Fatal(msg)
		}
	}

	nodes := []INode{
		NewNode("a"),
		NewNode("b"),
		NewNode("c"),
		NewNode("d"),
		NewNode("e"),
		NewNode("f"),
		NewNode("g"),
	}

	lines := []ILine{
		NewLine("a", "b", 9),
		NewLine("a", "c", 2),

		NewLine("b", "c", 6),
		NewLine("b", "d", 3),
		NewLine("b", "e", 1),

		NewLine("c", "d", 2),
		NewLine("c", "f", 9),

		NewLine("d", "e", 5),
		NewLine("d", "f", 6),

		NewLine("e", "f", 3),
		NewLine("e", "g", 7),

		NewLine("f", "g", 4),
	}

	for _, it := range lines[:] {
		lines = append(lines, NewLine(it.To(), it.From(), it.Weight()))
	}

	ok, path := Finder.FindPath(nodes, lines, "a", "g")
	if !ok {
		t.Fatal("failed to find min path")
	}
	fnPathToString := func(nodes []INode) string {
		items := make([]string, len(nodes))
		for i, it := range nodes {
			items[i] = fmt.Sprintf("%s", it)
		}
		return strings.Join(items, " ")
	}
	pathString := fnPathToString(path)
	fnAssertTrue(pathString == "a(0) c(2) d(4) f(10) g(14)", "incorrect path")
}
