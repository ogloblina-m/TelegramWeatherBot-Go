package main

import (
	"fmt"
	"log"
	"os"

	owm "github.com/briandowns/openweathermap"
)

type SearchResults struct {
	ready   bool
	Query   string
	Results []Result
}
type Result struct {
	Name, Description, URL string
}

var apiKey = os.Getenv("OWM_API_KEY")

func main() {
	w, err := owm.NewCurrent("F", "ru", apiKey)
	if err != nil {
		log.Fatalln(err)
	}

	w.CurrentByName("Ekb")
	fmt.Println(w)
}
