package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	//writer: mi fuente de datos es el std output.
	var writer io.Writer = os.Stdout

	//encoder:
	encoder := json.NewEncoder(writer)

	p := Person{
		Name: "John",
		Age:  30,
	}

	//encodes the variable to a json and writes it to the std output.
	if err := encoder.Encode(p); err != nil {
		fmt.Println(err.Error())
	}
}
