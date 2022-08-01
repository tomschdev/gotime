package bst

import (
	"fmt"
	"sync"
	"github.com/cheekybits/genny/generic"
)

type Item generic.Type
type Node struct {
	key   int
	value Item
	right *Node
	left  *Node
}

type ItemBinarySearchTree struct {
	root *Node
	lock sync.RWMutex
}

func (bst *ItemBinarySearchTree) ToString() {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	fmt.Println("---------------------")
	stringify(bst.root, 0)
	fmt.Println("---------------------")
}

func stringify(node *Node, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "            "
		}
		format += "---[ "
		level++
		stringify(node.left, level)
		fmt.Printf(format+"%d\n", n.key)
		stringify(node.right, level)

	}
}
func (bst *ItemBinarySearchTree) Insert(key int, value Item) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	n := &Node{key, value, nil, nil}

	if bst.root == nil {
		bst.root = n
	} else {
		insertNode(bst.root, n)
	}
}

func insertNode(node *Node, newNode *Node) {
	if newNode.key < node.key {
		if node.left == nil {
			node.left = newNode
		} else {
			insertNode(node.left, newNode)
		}
	} else {
		if node.right == nil {
			node.right = newNode
		} else {
			insertNode(node.right, newNode)
		}
	}

}
func (bst *ItemBinarySearchTree) InOrderTraverse(f func(item)) {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	inOrderTraverse(bst.root, f)
}

func inOrderTraverse(n *Node, f func(item)) {
	if n != nil {
		inOrderTraverse(n.left, f)
		f(n.value)
		inOrderTraverse(n.right, f)
	}
}func (bst *ItemBinarySearchTree) InOrderTraverse(f func(item)) {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	inOrderTraverse(bst.root, f)
}

func inOrderTraverse(n *Node, f func(item)) {
	if n != nil {
		inOrderTraverse(n.left, f)
		f(n.value)
		inOrderTraverse(n.right, f)
	}
}
