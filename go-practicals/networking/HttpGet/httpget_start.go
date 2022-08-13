package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func getRequestTest() {
	const httpbin = "https://httpbin.org/get"

	// Perform a GET operation
	resp, err := http.Get(httpbin)
	if err != nil {
		return
	}

	// The caller is responsible for closing the response
	defer resp.Body.Close() // close responses body stream

	// We can access parts of the response to get information:
	fmt.Println("Status: ", resp.Status)                // 200 OK
	fmt.Println("Status Code: ", resp.StatusCode)       // 200
	fmt.Println("Protocol: ", resp.Proto)               // HTTP/2.0
	fmt.Println("Content Length: ", resp.ContentLength) // 274

	// Use a String Builder to build the content from bytes
	var sb strings.Builder

	// Read the content and write it to the builder
	content, _ := ioutil.ReadAll(resp.Body)
	bytecount, _ := sb.Write(content)

	// Format the output
	fmt.Println(bytecount)   // 274
	fmt.Println(sb.String()) // shows full response in json format
}

func main() {
	// Execute a GET request
	getRequestTest()
}
