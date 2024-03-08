package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title       string   `json:"title,omitempty"`
	ReleaseYear int      `json:"release_year"`
	Director    string   `json:"director"`
	Actors      []string `json:"actors,onitempty"`
	Password    string   `json:"-"`
	field       string   //this field won't show on json because it isn't in capital letter. Even if it was indicated with the json tag, the field is not exportable (it's private).
}

func main() {
	movie := Movie{
		Title:       "Inception",
		ReleaseYear: 2010,
		Director:    "Christopher Nolan",
		Actors:      []string{"Leonardo DiCaprio", "Elliot Page", "Tom Hardy"},
		Password:    "123456",
		field:       "field",
	}

	bytes, err := json.Marshal(movie)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bytes))
}
