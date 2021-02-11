package nlp_test

import (
	"fmt"

	"github.com/353solutions/nlp"
)

func ExampleTokenize() {
	text := "Hi, how are you feeling today?"
	tokens := nlp.Tokenize(text)
	fmt.Println(tokens)

	// Output:
	// [hi feel today]
}
