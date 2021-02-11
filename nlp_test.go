package nlp

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// Exercise: read test cases from tokenize_cases.json
func TestTokenize(t *testing.T) {
	testCases := []struct {
		text   string
		tokens []string
	}{
		{"", nil},
		{"hi", []string{"hi"}},
		{"HI", []string{"hi"}},
		{"Who's on first?", []string{"who", "s", "on", "first"}},
	}

	for _, tc := range testCases {
		name := tc.text
		if name == "" {
			name = "<empty>"
		}
		t.Run(name, func(t *testing.T) {
			// require := require.New(t)
			tokens := Tokenize(tc.text)
			// require.Equal(tc.tokens, tokens)
			require.Equal(t, tc.tokens, tokens)
		})
	}
}

func TestMain(m *testing.M) {
	fmt.Println("Before main")
	// containerID = docker run -d -p 3306:3306 353solutions/testdb
	// wait for container
	out := m.Run()
	fmt.Println("After main")
	// docker rm -f containerID
	os.Exit(out)
}
