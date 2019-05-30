package nlp

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

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

func loadCases(t *testing.T) []TestCase {
	var testCases []TestCase

	file, err := os.Open("tokenizer_cases.json")
	if err != nil {
		t.Fatalf("can't open test cases - %s", err)
	}
	defer file.Close()

	dec := json.NewDecoder(file)
	if err := dec.Decode(&testCases); err != nil {
		t.Fatalf("bad JSON in test cases - %s", err)
	}
	return testCases
}

func TestTokenizeMulti(t *testing.T) {
	testCases := loadCases(t)

	for _, tc := range testCases {
		name := tc.Text
		if len(name) == 0 {
			name = "<empty>"
		}
		t.Run(name, func(t *testing.T) {
			out := Tokenize(tc.Text)
			if !reflect.DeepEqual(tc.Expected, out) {
				t.Fatalf("%#v != %#v", tc.Expected, out)
			}
		})

	}
}

func TestTokenize(t *testing.T) {
	text := "Who's on first?"
	out := Tokenize(text)
	expected := []string{"who", "s", "on", "first"}

	if !reflect.DeepEqual(expected, out) {
		// See also t.Failf
		// Using %#v will output more information
		t.Fatalf("%#v != %#v", expected, out)
	}
}
