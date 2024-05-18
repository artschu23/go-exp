package main

import (
	"encoding/json"
	"fmt"

	"github.com/artschu23/go-exp/response"
)

type NaiveResponse struct {
	Data interface{} `json:"data"`
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// NOTE: NaiveResponse and Response both implement the same json schema
func send(person Person) []byte {
	payload, _ := json.Marshal(response.Response[Person]{
		Data: person,
	})
	return payload
}

func receiveWithAssertion(payload []byte) {
	var response NaiveResponse

	err := json.Unmarshal(payload, &response)
	if err != nil {
		panic(err)
	}

	// NOTE: it's the pitfall when using assertion
	//
	// person, ok := response.Data.(Person)
	person, ok := response.Data.(map[string]interface{})
	if !ok {
		panic(fmt.Sprintf("invalid data: %s", response))
	}

	fmt.Printf("data: %+v\n", person)
}

func receive(payload []byte) {
	var response response.Response[Person]

	err := json.Unmarshal(payload, &response)
	if err != nil {
		panic(err)
	}

	fmt.Printf("data: %+v\n", response.Data)
}

func main() {
	person := Person{
		Name: "John",
		Age:  20,
	}

	payload := send(person)

	// show two receive function presenting the same output
	receiveWithAssertion(payload)
	receive(payload)
}
