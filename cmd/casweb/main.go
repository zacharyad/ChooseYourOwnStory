package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/zacharyad/cas"
)

func main() {

	port := flag.Int("port", 3000, "The port to start the story")

	filename := flag.String("file", "story.json", "the JSON with the story")
	flag.Parse()

	fmt.Printf("Reading from the story in: %s\n", *filename)

	f, err := os.Open(*filename)

	if err != nil {
		panic(err)
	}

	story, err := cas.JsonStory(f)

	if err != nil {
		panic(err)
	}

	h := cas.NewHandler(story)

	fmt.Printf("Starting on port: %d\n", *port)
	http.Get("/")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
