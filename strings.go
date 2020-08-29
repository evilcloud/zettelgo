package main

import (
	"github.com/mozillazg/go-unidecode"
	"strings"
	"unicode"
)

func cleanText(text string) string {
	unidecoded := unidecode.Unidecode(text)
	x := func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r) && !unicode.IsSpace(r)
	}
	ascii := strings.Join(strings.FieldsFunc(unidecoded, x), "")
	return ascii
}
