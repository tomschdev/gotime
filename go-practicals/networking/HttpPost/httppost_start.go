package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func postRequestTest() {
	const httpbin = "https://httpbin.org/post"

	// POST operation using Post
	reqBody := strings.NewReader(`
	{
		"x": "this",
		"y": "that"
	}
	`)
	resp, err := http.Post(httpbin, "application/json", reqBody)
	if err != nil {
		return
	}
	content, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Printf("%s\n", content)

	// POST operation using PostForm
	data := url.Values{}
	data.Add("field1", "fielded added via values")
	data.Add("field2", "2509")
	// PostForm issues a POST to the specified URL, with data's keys and
	// values URL-encoded as the request body.
	resp, err = http.PostForm(httpbin, data)
	content, _ = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Printf("%s\n", content)

}

func main() {
	// Execute a POST
	postRequestTest()
}
