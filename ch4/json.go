package main

import (
	"fmt"
	"log"
	"encoding/json"
)

type Movie struct {
	Title string
	Year int
	Color bool
	Actors []string
}

func main() {
	var movies = []Movie {
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string {"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string {"Paul Newman"}},
	}
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("Json marsalling failed %s", err)
	}
	fmt.Printf("%s\n", data)
}
