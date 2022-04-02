package soundslike

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCmp	tests the Cmp function
func TestCmp(t *testing.T) {
	var o Phonic
	o.Load("words.txt")
	pair := []string{
		"nation", "station",
		"leak", "creek",
		"cancan", "toucan",
		"unique", "antique",
	}
	l := len(pair)
	for i := range pair {
		if i%2 == 0 {
			a := o.FindEndingToken(pair[i])
			b := o.FindEndingToken(pair[i+1])
			na := o.FindEndingToken(pair[(i+2)%l])
			assert.NotEqual(t, "", a, "a makes result")
			assert.NotEqual(t, "", b, "b makes result")
			assert.Equalf(t, a, b, "rhyme match found %s %s", pair[i], pair[i+1])
			assert.NotEqual(t, a, na, "rhyme match NOT found")
		}
	}
}
