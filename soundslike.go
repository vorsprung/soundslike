package soundslike

import (
	"bufio"
	"fmt"

	"log"
	"os"
	"sort"

	"github.com/dlclark/metaphone3"
)

// FindEndingPhoneme   this will be the rhymer
func FindEndingPhoneme(s string) string {
	e := &metaphone3.Encoder{}
	l := len(s)
	for i := l; i >= 0; i-- {
		p, _ := e.Encode(s[i:l])
		if p != "" {
			return p
		}

	}
	return ""
}

// LoadDict  encode all words to metaphone word
func LoadDict(path string) map[string][]string {
	rv := make(map[string][]string)
	cnt := make(map[string]int)

	e := &metaphone3.Encoder{}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		code, _ := e.Encode(word)
		rv[code] = append(rv[code], word)
		cnt[code] = cnt[code] + 1
	}
	for _, x := range sortDictByIntKey(cnt) {
		if x.Value > 2 {
			//fmt.Printf("%s %d %v\n", x.Key, x.Value, rv[x.Key])
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return rv

}

// SoundsLike  what's it like???
func SoundsLike(s string) []string {
	e := &metaphone3.Encoder{}
	primary, secondary := e.Encode(s)
	d := LoadDict("words.txt")
	fmt.Println(s)
	fmt.Printf("primary=%s, second=%s\n=====\n", primary, secondary)
	//fmt.Printf("primary values %v \n", d[primary])
	//fmt.Printf("secondary values %v \n", d[secondary])

	return d[primary]
}

type kv struct {
	Key   string
	Value int
}

func sortDictByIntKey(m map[string]int) []kv {
	var ss []kv
	for k, v := range m {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	return ss
}
