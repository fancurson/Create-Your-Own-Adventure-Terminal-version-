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

/****************************/
// functional options

type StoryParams struct {
	Story   Story
	ArcName string
}

type OptionFunc func(sp *StoryParams)

func NewStory(s Story, options ...OptionFunc) *StoryParams {
	sp := &StoryParams{
		Story:   s,
		ArcName: "intro",
	}

	for _, option := range options {
		option(sp)
	}
	return sp
}

/****************************/

func PlayStory(s StoryParams) error {
	for {
		value := s.Story[s.ArcName]
		if len(value.Options) <= 0 {
			fmt.Printf("Congratulations! You end this story")
			break
		}

		//Title output
		fmt.Printf("\n%s\n\n", strings.ToUpper(s.ArcName))
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
		typingAnswer(&s, value)
	}
	return nil
}

func typingAnswer(s *StoryParams, value Chapter) {
	for {
		var answer string
		fmt.Printf("I want to choose ")
		fmt.Scanf("%s\n", &answer)
		answer = strings.TrimSpace(strings.ToLower(answer))

		for _, val := range value.Options {
			if val.Arc == answer {
				s.ArcName = answer
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
