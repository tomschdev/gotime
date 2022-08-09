package linkedlist

import (
	"fmt"
	"testing"
)

var ll ItemLinkedList

func TestAppend(t *testing.T) {
	if !ll.IsEmpty() {
		t.Error("list should be empty")
	}
	ll.Append("first")
	if ll.IsEmpty() {
		t.Error("list should not be empty")
	}
	if size := ll.Size(); size != 1 {
		t.Errorf("wrong count, expceted 1 - got %d", size)
	}
	ll.Append("second")
	ll.Append("third")
}

func TestRemove(t *testing.T) {
	_, err := ll.RemoveAt(1)
	if err != nil {
		t.Errorf("unexpected error thrown from RemoveAt: %s", err)
	}
	if size := ll.Size(); size != 2 {
		t.Errorf("wrong count, expected 2 - got %d", size)
	}
}
func TestInsert(t *testing.T) {
	err := ll.Insert(2, "second2")
	if err != nil {
		t.Errorf("unexpected error thrown from Insert: %s", err)
	}
	if size := ll.Size(); size != 3 {
		t.Errorf("wrong count, expected 3 - got %d", size)
	}
	err = ll.Insert(0, "zero")
	if err != nil {
		t.Errorf("unexpected error thrown from Insert(0, ): %s", err)
	}
}
func TestIndexOf(t *testing.T) {
	if i := ll.IndexOf("zero"); i != 0 {
		t.Errorf("expected position 0 - got %d", i)
	}
	if i := ll.IndexOf("second2"); i != 2 {
		t.Errorf("expected positon 2 - got %d", i)
	}
	if i := ll.IndexOf("third"); i != 3 {
		t.Errorf("expected positon 3 - got %d", i)
	}
}
func TestHead(t *testing.T) {
	ll.Append("zero")
	h := ll.Head()
	if "zero" != fmt.Sprint(h.content) {
		t.Errorf("expected `zero` - got %s", fmt.Sprint(h.content))
	}
}
