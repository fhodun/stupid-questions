package qp

import (
	"strings"

	"github.com/agnivade/levenshtein"
	"github.com/fhodun/stupid-questions/utils"
	"github.com/sirupsen/logrus"
)

// SentenceTag dupa
type SentenceTag struct {
	Weight int
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
	MinWeight   int
	MaxDistance int
}

func splitWords(str string) []string {
	words := make([]string, 0)

	last := 0
	for i, c := range str {
		if c == ' ' || c == '?' {
			word := strings.TrimSpace(str[last:i])
			if len(word) == 0 {
				continue
			}
			words = append(words, word)
			last = i + 1
		}
		if c == '?' {
			words = append(words, "?")
		}
	}
	lastWord := str[last:]
	if len(lastWord) > 0 {
		words = append(words, str[last:])
	}

	return words
}

// ParseString dupa
func (qp QuestionParser) ParseString(str string) *Sentence {
	str, err := utils.RemovePolishCharacters(str)
	if err != nil {
		logrus.Errorf("Fail removing polish shit %s\n", err.Error())
	}

	words := splitWords(str)

	for _, sentence := range qp.Sentences {
		// Check if string containts primary word
		// If it doesn't contain, continue looping over the rest of sentences
		// e.g if primary word == "testportal"
		// "hur dur testportal to guwno"   will output true
		// "hur dur to guwno"              will output false
		if !utils.LevenshteinStringContains(str, sentence.PrimaryWord, qp.MaxDistance) {
			continue
		}
		if len(sentence.Tags) == 0 {
			return &sentence
		}

		// Set weight, that will determine if message is stupid enough to respond it
		weight := 0

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
