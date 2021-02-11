package nlp

//go:generate sh -c "go run gen_stop.go < stop_words.txt > stop_words.go"
//go:generate gofmt -w stop_words.go

import (
	"regexp"
	"strings"

	"github.com/353solutions/nlp/stemmer"
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
		token := strings.ToLower(w)
		token = stemmer.Stem(token)
		if len(token) > 0 && !StopWords[token] {
			tokens = append(tokens, token)
		}
	}
	return tokens
}
