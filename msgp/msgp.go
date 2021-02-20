package msgp

import (
	"fmt"
	"strings"

	"github.com/agnivade/levenshtein"
	"github.com/fhodun/stupid-questions/utils"
)

type SentenceTag struct {
	Weight uint
	Text   string
}

// Word is used for defining banned words
type Sentence struct {
	PrimaryWord string
	Tags        []SentenceTag
	Answer      string
}

type MessageParser struct {
	Sentences   []Sentence
	MinWeight   uint
	MaxDistance int
}

func (mp MessageParser) ParseString(str string) *Sentence {
	words := strings.Split(str, " ")

	for _, sentence := range mp.Sentences {
		if !utils.LevenshteinStringContains(str, sentence.PrimaryWord, mp.MaxDistance) {
			continue
		}

		weight := uint(0)
		hasAnyTag := false

		for _, word := range words {
			if word == sentence.PrimaryWord {
				continue
			}
			for _, tag := range sentence.Tags {
				fmt.Printf("distance: %d from %s to %s\n", levenshtein.ComputeDistance(word, tag.Text), word, tag.Text)
				if distance := levenshtein.ComputeDistance(word, tag.Text); distance < mp.MaxDistance {
					hasAnyTag = true
					weight += tag.Weight
					if weight > mp.MinWeight {
						return &sentence
					}
				}
			}
		}
		if !hasAnyTag {
			return nil
		}
	}
	return nil
}
