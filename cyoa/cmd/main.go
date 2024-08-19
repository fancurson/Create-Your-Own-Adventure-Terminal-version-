package main

import (
	"cyoa"
	"flag"
	"log"
	"os"
)

func main() {
	fileName := flag.String("storyFile", "story.json", "fill the file with the story")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatalf("Open file error: %v", err)
	}
	defer file.Close()

	story, err := cyoa.JsonReader(file)
	if err != nil {
		log.Fatalf("Read file error: %v", err)
	}

	params := cyoa.NewStory(story/*, WithCustomArc("new-york")*/)
	if cyoa.PlayStory(*params); err != nil {
		panic(err)
	}
}

func WithCustomArc(customArc string) func(*cyoa.StoryParams) {
	return func(st *cyoa.StoryParams) {
		st.ArcName = customArc
	}
}
