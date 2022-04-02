package soundslike

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/dlclark/metaphone3"
)

type setup struct {
	lookup map[string][]string
	fname  string
}

// Phonic   datatype for methods
type Phonic struct {
	Setup []setup
	enc   *metaphone3.Encoder
}

// Load  setup a new Phonic with a file contents
func (e *Phonic) Load(fnames ...string) {
	e.enc = &metaphone3.Encoder{}
	e.Setup = make([]setup, len(fnames))
	for ix, v := range fnames {
		e.Setup[ix].lookup = e.LoadDict(v)
	}
}

// LoadDict   parse a specific dict and index by encoding
func (e Phonic) LoadDict(path string) map[string][]string {
	rv := make(map[string][]string)

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// file has one word per line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// get the word, find it's metaphone encoding
		word := scanner.Text()
		code, _ := e.enc.Encode(word)
		// rv code is a dict keyed by metaphone code
		// values are lists of matching words
		rv[code] = append(rv[code], word)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return rv

}

//SoundsLike   finds matches for whole word
func (e Phonic) SoundsLike(s string) []string {
	p, _ := e.enc.Encode(s)
	fmt.Printf("%s encodes to %s\n", s, p)
	fmt.Printf("lookup table has %d keys\n", len(e.Setup[0].lookup))
	return e.Setup[0].lookup[p]
}

//RhymeWith  with find matches on endings, ie "nearly" "dearly"
func (e Phonic) RhymeWith(s string) []string {
	rt := e.FindEndingToken(s)
	if rt != "" {
		return e.Setup[0].lookup[rt]
	}
	return []string{}
}

// FindEndingToken  tokenises just the ending
func (e Phonic) FindEndingToken(s string) string {
	const adjust = 3
	l := len(s)
	for i := l - adjust; i >= 0; i-- {
		p, _ := e.enc.Encode(s[i:l])
		if p != "" {
			return p
		}

	}
	return ""
}
