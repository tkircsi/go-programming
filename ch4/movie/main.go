package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

func main() {

	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s\n", err)
	}
	// fmt.Printf("%s\n", data)
	err = writeJSON("movies.json", data)
	if err != nil {
		log.Fatalf("Saving JSON failed: %s\n", err)
	}

	formattedData, err := json.MarshalIndent(movies, "", "  ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s\n", err)
	}
	err = writeJSON("movies_format.json", formattedData)
	if err != nil {
		log.Fatalf("Saving JSON failed: %s\n", err)
	}

	d, err := ioutil.ReadFile("movies_format.json")
	if err != nil {
		log.Fatal(err)
	}
	var titles []struct{ Title string }
	err = json.Unmarshal(d, &titles)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", titles)

}

func writeJSON(filename string, data []byte) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	_, err = w.Write(data)
	w.Flush()
	if err != nil {
		return err
	}
	return nil
}
