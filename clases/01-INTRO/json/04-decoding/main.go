package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	rd := strings.NewReader(`{"name": "John", "age": 30}`)

	//intermediario que permite leer desde una fuente de datos y enviarlo a una variable.
	decoder := json.NewDecoder(rd)

	var p Person

	//los datos que vaya leyendo el decoder desde el reader los va pasando a la variable p.
	if err := decoder.Decode(&p); err != nil {
		fmt.Println(err.Error())
	}
}
