package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// get the default temporary directory path
	tempPath := os.TempDir()
	fmt.Println("default temp path is: ", tempPath)
	// create a temporary file use TempFile
	tempContent := []byte("This is some temp content")
	tmpFile, err := ioutil.TempFile(tempPath, "tempfile_*.txt")
	if err != nil {
		panic(err)
	}
	// Read and print the tempFile content
	if _, err := tmpFile.Write(tempContent); err != nil {
		panic(err)
	}

	// Remove the temp file when finished
	data, _ := ioutil.ReadFile(tmpFile.Name())
	fmt.Printf("%s\n", data)
	fmt.Println("tmp file is named", tmpFile.Name())
	defer os.Remove(tmpFile.Name())
	// create a temporary directory with ioutil's version of TempDir
	tmpDir, err := ioutil.TempDir("", "tempdir*")
	if err != nil {
		panic(err)
	}
	fmt.Println("the tempdir is named: ", tmpDir)
	defer os.RemoveAll(tmpDir)
}
