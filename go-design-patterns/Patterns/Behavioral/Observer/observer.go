package main

import "fmt"

// Define the interface for an observer type
type observer interface {
	onUpdate(data string)
}

// Our DataListener observer will have a name
type DataListener struct {
	Name string
}

// To conform to the interface, it must have an onUpdate function
func (ds *DataListener) onUpdate(data string) {
	fmt.Printf("The new data is: %s - according to %s\n", data, ds.Name)
}
