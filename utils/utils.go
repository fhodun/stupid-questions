package utils

import (
	"github.com/agnivade/levenshtein"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// LevenshteinStringContains dupa
func LevenshteinStringContains(str string, substr string, maxDistance int) bool {
	for i := 0; i < len(str)-len(substr)+1; i++ {
		if levenshtein.ComputeDistance(substr, str[i:i+len(substr)]) < maxDistance {
			return true
		}
	}
	return false
}

// RemovePolishCharacters dupa
func RemovePolishCharacters(str string) (string, error) {
	trans := transform.Chain(
		norm.NFD,
		runes.Map(func(r rune) rune {
			switch r {
			case 'ą':
				return 'a'
			case 'ć':
				return 'c'
			case 'ę':
				return 'e'
			case 'ł':
				return 'l'
			case 'ń':
				return 'n'
			case 'ó':
				return 'o'
			case 'ś':
				return 's'
			case 'ż':
				return 'z'
			case 'ź':
				return 'z'
			}
			return r
		}),
		norm.NFC,
	)
	res, _, err := transform.String(trans, str)

	return res, err
}
