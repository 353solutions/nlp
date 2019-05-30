package nlp

import (
	"encoding/json"
	"os"
	//	"reflect"
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
)

func TestTokenizerQuick(t *testing.T) {
	fn := func(text string) bool {
		tokens := Tokenize(text)
		// Sanity check
		numTokens := len(wordRe.FindAllString(text, -1))
		return numTokens == len(tokens)
	}

	err := quick.Check(fn, nil)
	require.NoError(t, err, "quick")
}

/* Place the test cases in tokenizer_cases.json
[
	{"text": "hi", "output": ["hi"]},
	{"text": "HI", "output": ["hi"]},
	{"text": "What's on second?", "output": ["what", "s", "on", "second"]},
	{"text": "", "output": []}
]
*/

type TestCase struct {
	Text     string
	Expected []string `json:"output"`
}

/* Why we have require.InDelta
func TestFloat(t *testing.T) {
	a, b := 1.1, 1.1
	out := a * b
	expected := 1.21
	//require.Equal(t, expected, out, "float")
	require.InDelta(t, expected, out, 0.00001, "float")

}
*/

func loadCases(t *testing.T) []TestCase {
	require := require.New(t)
	var testCases []TestCase

	casesFile := "tokenizer_cases.json"
	file, err := os.Open(casesFile)
	require.NoErrorf(err, "open test cases - %#v", casesFile)
	/*
		if err != nil {
			t.Fatalf("can't open test cases - %s", err)
		}
	*/
	defer file.Close()

	dec := json.NewDecoder(file)
	require.NoError(dec.Decode(&testCases), "bad JSON in test cases")
	/*
		if err := dec.Decode(&testCases); err != nil {
			t.Fatalf("bad JSON in test cases - %s", err)
		}
	*/
	return testCases
}

func TestTokenizeMulti(t *testing.T) {
	require := require.New(t)
	testCases := loadCases(t)

	for _, tc := range testCases {
		name := tc.Text
		if len(name) == 0 {
			name = "<empty>"
		}
		t.Run(name, func(t *testing.T) {
			out := Tokenize(tc.Text)

			require.Equal(tc.Expected, out)
			/*
				if !reflect.DeepEqual(tc.Expected, out) {
					t.Fatalf("%#v != %#v", tc.Expected, out)
				}
			*/
		})

	}
}

func TestTokenize(t *testing.T) {
	require := require.New(t)
	text := "Who's on first?"
	out := Tokenize(text)
	expected := []string{"who", "s", "on", "first"}
	require.Equal(expected, out)

	/*
		if !reflect.DeepEqual(expected, out) {
			// See also t.Failf
			// Using %#v will output more information
			t.Fatalf("%#v != %#v", expected, out)
		}
	*/
}

// Typical text length
var tokBenchText = `
Software engineering is what happens to programming when you add
time and other programmers.
	- Russ Cox
`

func BenchmarkTokenize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		toks := Tokenize(tokBenchText)
		// Using toks so compiler won't optimize them away
		if len(toks) != 16 {
			b.Fatal(len(toks))
		}
	}
}
