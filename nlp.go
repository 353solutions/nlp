package nlp

import (
	"regexp"
	"strings"
)

var (
	wordRe = regexp.MustCompile("[[:alpha:]]+")
)

// Tokenize returns a list of tokens in text (lower case)
func Tokenize(text string) []string {
	words := wordRe.FindAllString(text, -1)
	// Pre allocate 20 tokens (average sentence length)
	tokens := make([]string, 0, 20)
	for _, w := range words {
		// TODO: stem
		token := strings.ToLower(w)
		tokens = append(tokens, token)
	}
	return tokens
}
