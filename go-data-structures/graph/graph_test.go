package graph

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {

	var g ItemGraph
	fillGraph(&g)
	g.String()
}
func TestTraverse(t *testing.T) {
	var g ItemGraph
	fillGraph(&g)
	g.Traverse(func(n *Node) {
		fmt.Printf("%v\n", n)

	})
}

func fillGraph(g *ItemGraph) {
	na := Node{"A"}
	nb := Node{"B"}
	nc := Node{"C"}
	nd := Node{"D"}
	ne := Node{"E"}
	nf := Node{"F"}
	g.AddNode(&na)
	g.AddNode(&nb)
	g.AddNode(&nc)
	g.AddNode(&nd)
	g.AddNode(&ne)
	g.AddNode(&nf)
	g.AddEdge(&na, &nb)
	g.AddEdge(&na, &nc)
	g.AddEdge(&na, &nd)
	g.AddEdge(&nb, &ne)
	g.AddEdge(&nc, &nf)
	g.AddEdge(&nd, &ne)
}
