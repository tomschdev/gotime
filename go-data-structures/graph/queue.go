package graph

import "sync"

// data structure to hold nodes while traversing graph
type NodeQueue struct {
	items []Node
	lock  sync.RWMutex
}

func (s *NodeQueue) New() *NodeQueue {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.items = []Node{}
	return s
}

func (s *NodeQueue) Enqueue(t Node) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.items = append(s.items, t)
}

func (s *NodeQueue) Dequeue() *Node {
	s.lock.Lock()
	defer s.lock.Unlock()

	item := s.items[0]
	s.items = s.items[1:len(s.items)]
	return &item
}
func (s *NodeQueue) Front() *Node {
	s.lock.RLock()
	defer s.lock.RUnlock()
	item := s.items[0]
	return &item
}

func (s *NodeQueue) IsEmpty() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.items) == 0
}
func (s *NodeQueue) Size() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.items)
}
