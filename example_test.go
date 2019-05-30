package nlp_test

import (
	"fmt"

	"github.com/353solutions/nlp"
)

func ExampleTokenize() {
	text := "I don't know on third"
	tokens := nlp.Tokenize(text)
	fmt.Println(tokens)

	// Output:
	// [i don t know on third]
}
