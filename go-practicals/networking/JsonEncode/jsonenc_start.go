package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name       string   `json:"fullname"`
	Address    string   `json:"addr"`
	Age        int      `json:"-"`                     // use - to indicate that this field should be omitted
	FaveColors []string `json:"fav-colours,omitempty"` //exclude field if map does not contain value
}

func encodeExample() {
	// create some people data
	people := []person{
		{"Jane Doe", "123 Anywhere Street", 35, nil},
		{"John Public", "456 Everywhere Blvd", 31, []string{"Purple", "Yellow", "Green"}},
	}

	// Marshal is used to convert a data structure to JSON format
	// MarshalIndent is used to format the JSON string with indentation
	result, err := json.MarshalIndent(people, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", result)

}

func main() {
	// Encode Go data as JSON
	encodeExample()
}
