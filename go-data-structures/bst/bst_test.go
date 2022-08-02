package bst

import (
	"testing"
)

var bst ItemBinarySearchTree

func fillTree(bst *ItemBinarySearchTree) {
	bst.Insert(8, "8")
	bst.Insert(7, "7")
	bst.Insert(6, "6")
	bst.Insert(5, "5")
	bst.Insert(4, "4")
	bst.Insert(3, "3")
}
func TestInsert(t *testing.T) {
	t.Log("---TestInsert---")
	fillTree(&bst)
	bst.ToString()
	bst.Insert(9, "9")
	bst.ToString()
}
