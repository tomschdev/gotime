package bst

import (
	"fmt"
	"testing"
)

// helper function to populate tree
func fillTree(bst *ItemBinarySearchTree) {
	bst.Insert(8, "8")
	bst.Insert(4, "4")
	bst.Insert(10, "10")
	bst.Insert(2, "2")
	bst.Insert(6, "6")
	bst.Insert(1, "1")
	bst.Insert(3, "3")
	bst.Insert(5, "5")
	bst.Insert(7, "7")
	bst.Insert(9, "9")
	bst.Insert(11, "11")
}

// returns true if two slices are identical
func isSameSlice(a, b []string) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
func TestInsert(t *testing.T) {
	var bst ItemBinarySearchTree
	t.Log("---TestInsert---")
	fillTree(&bst)
	bst.ToString()
	bst.Insert(12, "12")
	bst.ToString()
}
func TestPreOrder(t *testing.T) {
	var bst ItemBinarySearchTree
	var result []string
	fillTree(&bst)
	bst.PreOrderTraverse(func(i Item) {
		result = append(result, fmt.Sprintf("%s", i))
	})
	if !isSameSlice(result, []string{"8", "4", "2", "1", "3", "6", "5", "7", "10", "9", "11"}) {
		t.Errorf("Traversal order incorrect:\nresult: %v \nexpected: %v\n", result, []string{"8", "4", "2", "1", "3", "6", "5", "7", "10", "9", "11"})
	}
}
func TestInOrder(t *testing.T) {
	var bst ItemBinarySearchTree
	var result []string
	fillTree(&bst)
	bst.InOrderTraverse(func(i Item) {
		result = append(result, fmt.Sprintf("%s", i))
	})
	if !isSameSlice(result, []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"}) {
		t.Errorf("Traversal order incorrect:\nresult: %v \nexpected: %v\n", result, []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"})
	}
}
func TestPostOrder(t *testing.T) {
	var bst ItemBinarySearchTree
	var result []string
	fillTree(&bst)
	bst.PostOrderTraverse(func(i Item) {
		result = append(result, fmt.Sprintf("%s", i))
	})
	if !isSameSlice(result, []string{"1", "3", "2", "5", "7", "6", "4", "9", "11", "10", "8"}) {
		t.Errorf("Traversal order incorrect:\nresult: %v \nexpected: %v\n", result, []string{"1", "3", "2", "5", "7", "6", "4", "9", "11", "10", "8"})
	}
}
func TestMax(t *testing.T) {
	var bst ItemBinarySearchTree
	fillTree(&bst)
	if fmt.Sprintf("%s", *bst.Max()) != "11" {
		t.Error("Max function broken, should be 11")
	}
}
func TestMin(t *testing.T) {
	var bst ItemBinarySearchTree
	fillTree(&bst)
	if fmt.Sprintf("%s", *bst.Min()) != "1" {
		t.Error("Min function broken, should be 1")
	}
}
func TestSearch(t *testing.T) {
	var bst ItemBinarySearchTree
	fillTree(&bst)
	if !bst.Search(1) || !bst.Search(8) || !bst.Search(11) {
		t.Error("Search (or insert) function broken, items not found")
	}
}
func TestRemove(t *testing.T) {
	var bst ItemBinarySearchTree
	fillTree(&bst)
	bst.Remove(1)
	if fmt.Sprintf("%s", *bst.Min()) != "2" {
		t.Error("Remove function broken")
	}
}
