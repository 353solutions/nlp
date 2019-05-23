package nlp

import (
	"regexp"
	"strings"
)

var (
	wordRe = regexp.MustCompile("[[:alpha:]]+")
)

func Tokenize(text string) []string {
	words := wordRe.FindAllString(text, -1)
	var tokens []string
	for _, w := range words {
		// TODO: stem
		token := strings.ToLower(w)
		tokens = append(tokens, token)
	}
	return tokens
}
