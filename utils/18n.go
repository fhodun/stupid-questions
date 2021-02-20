package utils

import (
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

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
	res, _, err := transform.String(trans, "ŻóŁć")

	return res, err
}
