package nlp

import (
	"regexp"
	"strings"
)

var (
	// "Who's on first" -> ["Who", "s", "on", "first"]
	wordRe = regexp.MustCompile("[[:alpha:]]+")
)

// Tokenize returns list of tokens in text
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
