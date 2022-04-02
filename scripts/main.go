package main

import (
	"flag"
	"fmt"

	"github.com/vorsprung/soundslike"
)

func main() {
	var word = flag.String("word", "superb", "synonym word")
	flag.Parse()
	//d := soundslike.LoadDict("words.txt")

	pho := &soundslike.Phonic{}

	pho.Load("wlist_match12.txt")

	fmt.Println("--soundslike")
	fmt.Println(pho.SoundsLike(*word))
	fmt.Println("--rhyme")
	fmt.Println(pho.RhymeWith(*word))
}
