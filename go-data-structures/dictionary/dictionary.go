package dictionary

import (
	"sync"

	"github.com/cheekybits/genny/generic"
)

type Key generic.Type
type Value generic.Type
type ValueDictionary struct {
	dict map[Key]Value
	lock sync.RWMutex
}

// creates entry of value v at key k
func (d *ValueDictionary) Set(k Key, v Value) {
	d.lock.Lock()
	defer d.lock.Unlock()
	if d.dict == nil {
		d.dict = make(map[Key]Value)
	}
	d.dict[k] = v
}

// deletes entry associated with key k
func (d *ValueDictionary) Delete(k Key) bool {
	d.lock.Lock()
	defer d.lock.Unlock()
	_, ok := d.dict[k]
	if ok {
		delete(d.dict, k)
	}
	return ok
}

// returns true if key k maps to a value, else false
func (d *ValueDictionary) Has(k Key) bool {
	d.lock.RLock()
	defer d.lock.RUnlock()
	_, ok := d.dict[k]
	return ok
}

// returns value v given key k
func (d *ValueDictionary) Get(k Key) Value {
	d.lock.RLock()
	defer d.lock.RUnlock()
	value, ok := d.dict[k]
	if ok {
		return value
	}
	return nil
}

// empties the dictioanry
func (d *ValueDictionary) Clear() {
	d.lock.Lock()
	defer d.lock.Unlock()
	d.dict = make(map[Key]Value)
}

// returns the number of items in the dictionary
func (d *ValueDictionary) Size() int {
	d.lock.RLock()
	defer d.lock.RUnlock()
	return len(d.dict)
}

// returns slice of all keys present
func (d *ValueDictionary) Keys() []Key {
	d.lock.RLock()
	defer d.lock.RUnlock()
	keys := []Key{}
	for i := range d.dict {
		keys = append(keys, i)
	}
	return keys
}

// returns slice of all values present
func (d *ValueDictionary) Values() []Value {
	d.lock.RLock()
	defer d.lock.RUnlock()
	values := []Value{}
	for i := range d.dict {
		values = append(values, d.dict[i])
	}
	return values
}
