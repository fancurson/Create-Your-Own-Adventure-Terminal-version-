package main

import (
	"cyoa"
	"flag"
	"log"
	"os"
)

func main() {

	StoryFile := flag.String("FileName", "story.json", "JSON file with story")
	flag.Parse()

	file, err := os.Open(*StoryFile)
	if err != nil {
		log.Fatalf("Open file error: %v", err)
	}
	defer file.Close()

	/************************************/

	story, err := cyoa.JsonReader(file)
	if err != nil {
		panic(err)
	}

	settings := cyoa.NewStory(story, cyoa.WithFirstArc("new-york"))

	if err := cyoa.PlayStory(settings); err != nil {
		panic(err)
	}

	/************************************/
}
