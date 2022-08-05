package graph

import (
	"fmt"
	"sync"

	"github.com/cheekybits/genny/generic"
)

type Item generic.Type

type Node struct {
	value Item
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.value)
}

type ItemGraph struct {
	nodes []*Node
	edges map[Node][]*Node
	lock  sync.RWMutex
}

func (g *ItemGraph) AddNode(n *Node) {
	g.lock.Lock()
	defer g.lock.Unlock()
	g.nodes = append(g.nodes, n)
}

func (g *ItemGraph) AddEdge(n1, n2 *Node) {
	g.lock.Lock()
	defer g.lock.Unlock()
	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}
	g.edges[*n1] = append(g.edges[*n1], n2)
	g.edges[*n2] = append(g.edges[*n2], n1) // *n2 must be dereferencing n2 which is a pointer to a Node, because the key should be a Node
}
func (g *ItemGraph) String() {
	g.lock.Lock()
	defer g.lock.Unlock()
	s := ""
	for _, v := range g.nodes {
		s += v.String() + " -> "
		near := g.edges[*v]
		for _, vv := range near {
			s += vv.String() + " "
		}
		s += "\n"
	}
	fmt.Println(s)
}
