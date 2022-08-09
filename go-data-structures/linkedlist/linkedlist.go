package linkedlist

import (
	"fmt"
	"sync"

	"github.com/cheekybits/genny/generic"
)

type Item generic.Type
type Node struct {
	content Item
	next    *Node
}

type ItemLinkedList struct {
	head *Node
	size int
	lock sync.RWMutex
}

func (ll *ItemLinkedList) Append(t Item) {
	ll.lock.Lock()
	defer ll.lock.Unlock()
	node := Node{t, nil}
	if ll.head == nil {
		ll.head = &node
	} else {
		tmp := ll.head
		for {
			if tmp.next == nil {
				break
			}
			tmp = tmp.next
		}
		tmp.next = &node
	}
	ll.size++
}
func (ll *ItemLinkedList) Insert(i int, t Item) error {
	ll.lock.Lock()
	defer ll.lock.Unlock()
	if i < 0 || i > ll.size {
		return fmt.Errorf("index %d out of bounds", i)
	}
	node := Node{t, nil} // set next when placed in list
	if i == 0 {
		node.next = ll.head
		ll.head = &node
		return nil
	}
	tmp := ll.head
	j := 0
	for j < i-2 { //-1 because zero indexing and size has 1 extra - another -1 because we have to avoid nil pointer exeception
		tmp = tmp.next
		j++
	}
	node.next = tmp.next
	tmp.next = &node
	ll.size++
	return nil
}
func (ll *ItemLinkedList) RemoveAt(i int) (*Item, error) {
	ll.lock.Lock()
	defer ll.lock.Unlock()
	if i < 0 || i > ll.size {
		return nil, fmt.Errorf("index %d out of bounds", i)
	}
	if i == 0 {
		remove := ll.head
		ll.head = ll.head.next
		ll.size--
		return &remove.content, nil
	}
	j := 0
	node := ll.head
	for j < i-1 {
		j++
		node = node.next
	}
	remove := node.next
	node.next = node.next.next
	ll.size--
	return &remove.content, nil
}
func (ll *ItemLinkedList) IndexOf(t Item) int {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	j := 0
	tmp := ll.head
	for {
		if tmp.content == t {
			return j
		}
		if tmp.next == nil {
			return -1
		}
		tmp = tmp.next
		j++
	}
}
func (ll *ItemLinkedList) IsEmpty() bool {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	if ll.head == nil {
		return true
	}
	return false
}
func (ll *ItemLinkedList) Size() int {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	j := 0
	for tmp := ll.head; tmp != nil; tmp = tmp.next {
		j++
	}
	return j
}
func (ll *ItemLinkedList) String() {
	defer ll.lock.RUnlock()
	node := ll.head
	j := 0
	for {
		if node == nil {
			break
		}
		j++
		fmt.Print(node.content)
		fmt.Print(" ")
		node = node.next
	}
	fmt.Println()
}
func (ll *ItemLinkedList) Head() *Node {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	return ll.head
}
