package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var words []string
	s := bufio.NewScanner(f)
	for s.Scan() {
		words = append(words, s.Text())
	}

	if err = s.Err(); err != nil {
		log.Fatal(err)
	}

	for i, word := range words {
		for _, word2 := range words[i:] {
			diffnum := 0
			// var a, b byte
			var j int
			for r := range word {
				if word[r] != word2[r] {
					diffnum++
					// a, b = word[r], word2[r]
					j = r
				}
				if diffnum > 1 {
					continue
				}
			}
			if diffnum == 1 {
				// fmt.Println(string(a), string(b))

				fmt.Println(word, j, string(word[j]))
				// break
			}
		}
	}
}
