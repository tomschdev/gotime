package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getFileLength(filepath string) int64 {
	f, err := os.Stat(filepath)
	handleErr(err)
	return f.Size()
}

func main() {
	// Use the ReadFile function to read an entire file
	content, err := ioutil.ReadFile("sampletext.txt")
	handleErr(err)
	// ReadFile reads the data as bytes, we have to convert to a string
	fmt.Println(string(content))
	// The Read function can perform a generic read
	const BuffSize = 20
	f, _ := os.Open("sampletext.txt")
	defer f.Close()

	b1 := make([]byte, BuffSize)
	for {
		n, err := f.Read(b1)
		if err != nil {
			if err != io.EOF {
				handleErr(err)
			}
			break
		}
		fmt.Println("bytes read: ", n)
		fmt.Println("content: ", string(b1[:n]))
	}

	// Get the length of a file
	l := getFileLength("sampletext.txt")
	fmt.Println("file length: ", l)

	// Use ReadAt to read from a specific index
	b2 := make([]byte, 10)
	_, err3 := f.ReadAt(b2, l-int64(len(b2))) // read 1- bytes from end of file
	handleErr(err3)
	fmt.Println("last 10 bytes of file as string: ", string(b2))
}
