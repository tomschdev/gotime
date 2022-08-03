package dictionary

import (
	"fmt"
	"testing"
)

func populateDictionary(count, start int) *ValueDictionary {
	dict := ValueDictionary{}
	for i := start; i < (count + start); i++ {
		dict.Set(
			fmt.Sprintf("key-%d", i),
			fmt.Sprintf("value-%d", i),
		)
	}
	return &dict
}
func TestSet(t *testing.T) {
	dict := populateDictionary(5, 0)
	if size := dict.Size(); size != 5 {
		t.Errorf("wrong size returned, Set may be at fault: expected:%d result:%d", 5, size)
	}
	dict.Set("key-1", "value-1") // should only update, not append
	if size := dict.Size(); size != 5 {
		t.Errorf("wrong size returned, Set may be at fault: expected:%d result:%d", 5, size)
	}
	dict.Set("key-6", "value-6")
	if size := dict.Size(); size != 6 {
		t.Errorf("wrong size returned, Set may be at fault): expected:%d result:%d", 5, size)
	}
}
func TestHas(t *testing.T) {
	dict := populateDictionary(5, 0)
	ok := dict.Has("key-1")
	if !ok {
		t.Error("expected entry to be there, Has may be at fault")
	}
	dict.Delete("key-1")
	ok = dict.Has("key-1")
	if ok {
		t.Error("entry should not be there, Has may be at fault")
	}
}
func TestGet(t *testing.T) {
	dict := populateDictionary(5, 0)
	value := dict.Get("key-1")
	if fmt.Sprintf("%s", value) != "value-1" {
		t.Error("get failed to return the correct value")
	}
}
func TestClear(t *testing.T) {
	dict := populateDictionary(5, 0)
	dict.Clear()
	if size := dict.Size(); size != 0 {
		t.Error("wrong size retuned, Clear may be at fault")
	}
}
func TestDelete(t *testing.T) {
	dict := populateDictionary(5, 0)
	dict.Delete("key-1")
	if size := dict.Size(); size != 4 {
		t.Error("wrong size returned, Delete may be at fault")
	}
}
func TestKeys(t *testing.T) {
	dict := populateDictionary(5, 0)
	items := dict.Keys()
	if len(items) != 5 {
		t.Error("Keys method is faulty")
	}
}
func TestValues(t *testing.T) {
	dict := populateDictionary(5, 0)
	items := dict.Values()
	if len(items) != 5 {
		t.Error("Values method is faulty")
	}
}
func TestSize(t *testing.T) {
	dict := populateDictionary(100, 0)
	items := dict.Values()
	if size := dict.Size(); size != len(items) {
		t.Error("size failed to return the correct item count")
	}
}
