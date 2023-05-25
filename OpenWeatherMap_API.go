package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type SearchResults struct {
	ready   bool
	Query   string
	Results []Result
}
type Result struct {
	Name, Description, URL string
}

func main() {
	file, err := os.Open("api_key.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	b, err := ioutil.ReadAll(file)
	var apiKey = (string(b))
	fmt.Println(string(apiKey))
}
