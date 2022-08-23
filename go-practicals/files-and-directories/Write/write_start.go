package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func checkFileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func writeFileData() {
	// create a new file
	f, err := os.Create("testfile.txt")
	handleErr(err)
	// it is idiomatic in Go to defer the close after you open the file
	defer f.Close()
	// get the Name of the file
	fmt.Println("The filename is: ", f.Name())
	// write some string data to the file
	f.WriteString("here is text\n")
	// write some individual bytes to the file
	data2 := []byte{'a', 'b', 'c', '\n'}
	f.Write(data2)
	// the Sync function forces the data to be written out
	f.Sync()
}

func appendFileData(fname string, data string) {
	// Use the lower-level OpenFile function to open the file in append mode
	f, err := os.OpenFile(fname, os.O_APPEND|os.O_WRONLY, 0644)
	handleErr(err)
	defer f.Close()
	_, err = f.WriteString(data)
}

func main() {
	// Simple example: dump some data to a file
	data := []byte("this is my data to go to file\n")
	ioutil.WriteFile("datafile.txt", data, 0644)

	// More complex example: Granular writing of data
	if !checkFileExists("testfile.txt") {
		writeFileData()
	} else {
		os.Truncate("testfile.txt", 10) // if file exists, trim to 10 bytes in length
	}
	// append data to a file
	appendFileData("testfile.txt", "stringeroo")
}
