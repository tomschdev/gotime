package main

// Define the interface for the observable type
type observable interface {
	registerObserver(obs observer)
	unregisterObserver(obs observer)
	notifyAll()
}

// The DataSubject will have a list of listeners
// and a field that gets changed, triggering them
type DataSubject struct {
	observers []DataListener
	field     string
}

// The ChangeItem function will cause the Listeners to be called
func (ds *DataSubject) ChangeItem(data string) {
	ds.field = data
	ds.notifyAll()
}

// This function adds an observer to the list
func (ds *DataSubject) registerObserver(o DataListener) {
	ds.observers = append(ds.observers, o)
}

// This function removes an observer from the list
func (ds *DataSubject) unregisterObserver(o DataListener) {
	var newObservers []DataListener
	for _, ob := range ds.observers {
		if o.Name != ob.Name {
			newObservers = append(newObservers, o)
		}
	}
	ds.observers = newObservers
}

// The notifyAll function calls all the listeners with the changed data
func (ds *DataSubject) notifyAll() {
	for _, ob := range ds.observers {
		ob.onUpdate(ds.field)
	}
}
