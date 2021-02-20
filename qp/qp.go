package qp

import (
	"strings"

	"github.com/agnivade/levenshtein"
	"github.com/fhodun/stupid-questions/utils"
)

// SentenceTag dupa
type SentenceTag struct {
	Weight uint
	Text   string
}

// Sentence dupa, word is used for defining banned words
type Sentence struct {
	PrimaryWord string
	Tags        []SentenceTag
	Answer      string
}

// QuestionParser dupa
type QuestionParser struct {
	Sentences   []Sentence
	MinWeight   uint
	MaxDistance int
}

// ParseString dupa
func (qp QuestionParser) ParseString(str string) *Sentence {
	// Split words by space
	words := strings.Split(str, " ")

	for _, sentence := range qp.Sentences {
		// Check if string containts primary word
		// If it doesn't contain, continue looping over the rest of sentences
		// e.g if primary word == "testportal"
		// "hur dur testportal to guwno"   will output true
		// "hur dur to guwno"              will output false
		if !utils.LevenshteinStringContains(str, sentence.PrimaryWord, qp.MaxDistance) {
			continue
		}

		// Set weight, that will determine if message is stupid enough to respond it
		weight := uint(0)

		// Loop over every word in message
		for _, word := range words {
			// If word is equal to the primary word skip this loop
			if word == sentence.PrimaryWord {
				continue
			}
			// Loop over all sentence tags
			for _, tag := range sentence.Tags {
				// Check if word matches with tag
				if distance := levenshtein.ComputeDistance(word, tag.Text); distance < qp.MaxDistance {
					// Add weight if match, and if this is enought to say that this message is stupid, simply return
					weight += tag.Weight
					if weight > qp.MinWeight {
						return &sentence
					}
				}
			}
		}

		// Skip rest of sentences if no matching tags are present
		if weight == 0 {
			return nil
		}
	}
	return nil
}
