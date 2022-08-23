package main

import (
	"fmt"
	"os"
)

func checkFileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err != nil { // returns a FileInfo interface which gives us details about a file
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func main() {
	// Check if a file exists
	exists := checkFileExists("sampletext.txt")
	// Use the Stat function to get file stats
	stats, err := os.Stat("sampletext.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("does file exist? ", exists)
	fmt.Println("Modification time", stats.ModTime())
	fmt.Println("Filename", stats.Name())
	fmt.Println("Size", stats.Size())
	fmt.Println("Mode", stats.Mode())
	fmt.Println("Is Dir?", stats.IsDir())
}
