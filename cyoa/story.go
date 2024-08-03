package cyoa

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

/****************************/
type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type StoryParams struct {
	story   Story
	arcName string
}

type StoryOption func(*StoryParams)

/****************************/
// functional options
func NewStory(s Story, options ...StoryOption) StoryParams {
	sp := StoryParams{
		story:   s,
		arcName: "intro",
	}

	for _, opt := range options {
		opt(&sp)
	}
	return sp
}

func WithFirstArc(arc string) StoryOption {
	return func(sp *StoryParams) {
		sp.arcName = arc
	}
}

/****************************/

func PlayStory(s StoryParams) error {
	for {
		value := s.story[s.arcName]
		if len(value.Options) > 0 {
			fmt.Printf("Congratulations! You end this story")
			break
		}

		//Title output
		fmt.Printf("\n%s\n\n", strings.ToUpper(s.arcName))
		fmt.Println(value.Title)

		// Paragraphs output
		for _, val := range value.Paragraphs {
			fmt.Printf("%s\n", val)
		}

		// Options output
		fmt.Printf("\n You have %d option(s)\n", len(value.Options))
		for i, val := range value.Options {
			fmt.Printf("Option %d: %v\n", i+1, val.Text)
			fmt.Printf("(Print '%s' to choose this variant)\n\n", val.Arc)
		}
		// Requesting for the answer
		TypingAnswer(&s, value)
	}
	return nil
}

func TypingAnswer(s *StoryParams, value Chapter) {
	for {
		var answer string
		fmt.Printf("I want to choose ")
		fmt.Scanf("%s\n", &answer)
		answer = strings.TrimSpace(strings.ToLower(answer))

		for _, val := range value.Options {
			if val.Arc == answer {
				s.arcName = answer
				return
			}
		}
		fmt.Println("Invalid option. Please choose a valid variant.")
	}
}

func JsonReader(f io.Reader) (Story, error) {
	decoder := json.NewDecoder(f)
	var story Story
	err := decoder.Decode(&story)
	if err != nil {
		return nil, fmt.Errorf("decode json file error: %v", err)
	}

	return story, nil
}
