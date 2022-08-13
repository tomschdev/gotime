package main

import (
	"fmt"
	"net/url"
)

func main() {
	// Define a URL
	s := "https://www.example.com:8000/user?username=joemarini"

	// Use the Parse function to parse the URL content
	result, _ := url.Parse(s)
	fmt.Println(result.Scheme)   //https
	fmt.Println(result.Host)     //www.example.com:8000
	fmt.Println(result.Port())   //8000
	fmt.Println(result.Path)     // /user
	fmt.Println(result.RawQuery) //username=joemarini

	// Extract the query components into a Values struct
	vals := result.Query()
	fmt.Println(vals["username"]) // joemarini

	// Create a URL from components
	newUrl := &url.URL{
		Scheme:   "https",
		Host:     "www.example.com",
		Path:     "/args",
		RawQuery: "x=1&y=2",
	}
	s = newUrl.String()
	fmt.Println(s) // https://www.example.com/args?x=1&y=2
	newUrl.Host = "www.tomschdevupdate.com"
	s = newUrl.String() // https://www.tomschdevupdate.com/args?x=1&y=2
	fmt.Println(s)

	// Create a new Values struct and encode it as a query string
	newVals := url.Values{}
	newVals.Add("z", "this")
	newVals.Add("x", "that")
	newUrl.RawQuery = newVals.Encode()
	s = newUrl.String()
	fmt.Println(s) // https://www.tomschdevupdate.com/args?x=that&z=this

}
