package stemmer

/* Test our stemmer with:
- runs -> run
- working -> work
- sleep -> sleep
*/

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var testCases = []struct {
	word string
	stem string
}{
	{"runs", "run"},
	{"working", "work"},
	{"sleep", "sleep"},
}

func TestStem(t *testing.T) {
	for _, tc := range testCases {
		name := fmt.Sprintf("%v", tc)
		t.Run(name, func(t *testing.T) {
			out := Stem(tc.word)
			require.Equal(t, tc.stem, out, "stem")
		})
	}
}
