package main

import "fmt"

// Iterate using a callback function
func main() {
	// use IterateBooks to iterate via a callback function
	lib.IterateBooks(myBookCallback)

	// Use IterateBooks to iterate via anonymous function
	lib.IterateBooks(func(b Book) error {
		fmt.Printf("Iterating book with anon callback: %s by %s\n", b.Name, b.Author)
		return nil
	})

	// create a BookIterator
	iter := lib.createIterator()
	for iter.hasNext() {
		book := iter.next()
		fmt.Println("book: ", book.Name, book.Author)
	}
}

// This callback function processes an individual Book object
func myBookCallback(b Book) error {
	fmt.Println("Book title:", b.Name)
	return nil
}
