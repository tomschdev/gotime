package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// create a directory
	os.Mkdir("newdir", os.ModePerm)
	// Remove will remove an item
	defer os.Remove("newdir")
	// create a deep directory
	os.MkdirAll("path/to/deep/dir", os.ModePerm)
	// RemoveAll will remove an item and all children
	defer os.RemoveAll("path")
	// get the current working directory
	dir, _ := os.Getwd()
	fmt.Println("current dir is ", dir)
	// get the directory of the current process
	exedir, _ := os.Executable()
	fmt.Println("execution dir is ", exedir)
	// read the contents of a directory
	contents, _ := ioutil.ReadDir(".")
	for _, f := range contents {
		fmt.Println(f.Name(), f.IsDir())
	}

}
