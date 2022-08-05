package graph

func (g *ItemGraph) Traverse(f func(*Node)) {
	g.lock.RLock()
	defer g.lock.RUnlock()
	q := NodeQueue{}
	q.New()
	n := g.nodes[0] // assumes graph has at least one entry
	q.Enqueue(*n)   // add first node to queue to initiate queue
	visited := make(map[*Node]bool)

	for {
		if q.IsEmpty() {
			break
		}
		node := q.Dequeue() // pop from queue
		visited[node] = true
		near := g.edges[*node]
		for _, v := range near {
			if !visited[v] {
				q.Enqueue(*v) // add all connected nodes to current node to queue
				visited[v] = true
			}
		}
		if f != nil {
			f(node) // process current node
		}
	}

}
