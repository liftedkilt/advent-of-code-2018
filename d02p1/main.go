package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var twos, threes int

	s := bufio.NewScanner(f)
	for s.Scan() {
		w := s.Text()

		var countstwo, countsthree bool

		for _, r := range w {
			switch strings.Count(w, string(r)) {
			case 2:
				countstwo = true
			case 3:
				countsthree = true
			}
		}

		if countstwo {
			twos++
		}

		if countsthree {
			threes++
		}
	}

	fmt.Println(twos * threes)

	if err = s.Err(); err != nil {
		log.Fatal(err)
	}
}
