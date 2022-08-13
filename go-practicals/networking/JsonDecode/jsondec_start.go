package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name       string   `json:"fullname"` // add struct tags to query json payload for 'fullname' and assign to Name field
	Address    string   `json:"addr"`
	Age        int      `json:"age"`
	FaveColors []string `json:"favecolours"`
}

func decodeExample() {
	// declare some sample JSON data to decode
	data := []byte(`
	{
		"fullname" : "John Q Public",
	 	"addr" : "987 Main St",
	 	"age": 45,
	 	"favecolours" : ["Purple","White","Gold"]
	}
	`)

	// JSON will be decoded into a person struct
	var p person

	// Test to see if the JSON is valid, and then decode
	valid := json.Valid(data) // just checks that the json structure is legal
	if valid {
		json.Unmarshal(data, &p)
		fmt.Printf("%#v\n", p)
	}

	// Data can also be decoded into a map structure
	// Unmarshal into a map
	// Useful when we don't know which fields to expect
	var m map[string]interface{}
	json.Unmarshal(data, &m)
	fmt.Printf("%#v\n", m)
	for k, v := range m {
		fmt.Printf("key: %v, value: %v\n", k, v)
	}
}

func main() {
	// Decode JSON into Go structs
	decodeExample()
}
