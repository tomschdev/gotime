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
	if node != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "            "
		}
		format += "---[ "
		level++
		stringify(node.left, level)
		fmt.Printf(format+"%d\n", node.key)
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

func (bst *ItemBinarySearchTree) InOrderTraverse(f func(Item)) {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	inOrderTraverse(bst.root, f)
}

func inOrderTraverse(n *Node, f func(Item)) {
	if n != nil {
		inOrderTraverse(n.left, f)
		f(n.value)
		inOrderTraverse(n.right, f)
	}
}

func (bst *ItemBinarySearchTree) PreOrderTraverse(f func(Item)) {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	preOrderTraverse(bst.root, f)
}

func preOrderTraverse(n *Node, f func(Item)) {
	if n != nil {
		f(n.value)
		inOrderTraverse(n.left, f)
		inOrderTraverse(n.right, f)
	}
}

func (bst *ItemBinarySearchTree) PostOrderTraverse(f func(Item)) {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	postOrderTraverse(bst.root, f)
}

func postOrderTraverse(n *Node, f func(Item)) {
	if n != nil {
		inOrderTraverse(n.left, f)
		inOrderTraverse(n.right, f)
		f(n.value)
	}
}

// returns min value in tree, nil if empty
func (bst *ItemBinarySearchTree) Min() *Item {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	n := bst.root
	if n == nil {
		return nil
	}
	for {
		if n.left == nil {
			return &n.value
		}
		n = n.left
	}
}

// returns max value in tree, nil if empty
func (bst *ItemBinarySearchTree) Max() *Item {
	bst.lock.Lock()
	defer bst.lock.RUnlock()
	n := bst.root
	if n == nil {
		return nil
	}
	for {
		if n.right == nil {
			return &n.value
		}
		n = n.right
	}
}

// returns true of key exists in tree
func (bst *ItemBinarySearchTree) Search(key int) bool {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	return search(bst.root, key)
}

func search(n *Node, key int) bool {
	if n == nil {
		return false
	}
	if key < n.key {
		return search(n.left, key)
	}
	if key > n.key {
		return search(n.right, key)
	}
	//key equals n.key
	return true

}

func (bst *ItemBinarySearchTree) Remove(key int) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	remove(bst.root, key)
}

func remove(n *Node, key int) *Node {
	if n == nil {
		return nil
	}
	if n.key > key {
		n.left = remove(n.left, key)
		return n
	}
	if n.key < key {
		n.right = remove(n.right, key)
		return n
	}
	// key = n.key
	if n.left == nil && n.right == nil {
		n = nil
		return nil
	}
	if n.left == nil {
		n = n.right
		return n
	}
	if n.right == nil {
		n = n.left
		return n
	}
	lmrs := n.right
	for {
		if lmrs != nil && lmrs.left != nil {
			lmrs = lmrs.left
		} else {
			break
		}
	}
	n.key, n.value = lmrs.key, lmrs.value
	n.right = remove(n.right, n.key)
	return n
}

// String prints a visual representation of the tree
func (bst *ItemBinarySearchTree) String() {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	fmt.Println("------------------------------------------------")
	stringify(bst.root, 0)
	fmt.Println("------------------------------------------------")
}
