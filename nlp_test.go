package nlp

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
)

func TestTokenizeFuzz(t *testing.T) {
	fn := func(text string) bool {
		//		t.Log(text)
		tokens := Tokenize(text)

		lText := strings.ToLower(text)
		for _, tok := range tokens {
			if !strings.Contains(lText, tok) {
				return false
			}
		}

		return true
	}

	err := quick.Check(fn, nil)
	if err != nil {
		t.Fatal(err)
	}
}

type TokenizeTestCase struct {
	Text   string
	Tokens []string `json:"output"`
}

func loadTokenizeCases(t *testing.T, fileName string) []TokenizeTestCase {
	require := require.New(t)

	file, err := os.Open(fileName)
	require.NoError(err)
	defer file.Close()

	var cases []TokenizeTestCase
	err = json.NewDecoder(file).Decode(&cases)
	require.NoError(err)
	return cases
}

// Exercise: read test cases from tokenize_cases.json
func TestTokenize(t *testing.T) {
	/*
		testCases := []struct {
			text   string
			tokens []string
		}{
			{"", nil},
			{"hi", []string{"hi"}},
			{"HI", []string{"hi"}},
			{"Who's on first?", []string{"who", "s", "on", "first"}},
		}
	*/

	for _, tc := range loadTokenizeCases(t, "testdata/tokenizer_cases.json") {
		name := tc.Text
		if name == "" {
			name = "<empty>"
		}
		t.Run(name, func(t *testing.T) {
			// require := require.New(t)
			tokens := Tokenize(tc.Text)
			// require.Equal(tc.Tokens, tokens)
			require.Equal(t, tc.Tokens, tokens)
		})
	}
}

// Example of "global" fixture
func TestMain(m *testing.M) {
	fmt.Println("Before main")
	// containerID = docker run -d -p 3306:3306 353solutions/testdb
	// wait for container
	out := m.Run()
	fmt.Println("After main")
	// docker rm -f containerID
	os.Exit(out)
}
