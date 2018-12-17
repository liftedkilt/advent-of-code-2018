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
	polymer := ""

	s := bufio.NewScanner(f)
	for s.Scan() {
		polymer = s.Text()
	}
	if err = s.Err(); err != nil {
		log.Fatal(err)
	}

	shortest, letter := len(polymer), ""
	for i := 'A'; i <= 'Z'; i++ {
		trimmed := strings.Replace(polymer, string(i), "", -1)
		trimmed = strings.Replace(trimmed, string(i+32), "", -1)

		t := []byte(trimmed)

		if l := reactPolymer(t); l < shortest {
			shortest, letter = l, string(i)
		}
	}
	fmt.Println(shortest, letter)
}

func reactPolymer(bytes []byte) int {
	for {
		length := len(bytes)
		for i := range bytes {
			if i < len(bytes)-1 {
				if bytes[i] == bytes[i+1]+32 || bytes[i] == bytes[i+1]-32 {
					bytes = append(bytes[:i], bytes[i+2:]...)
				}
			}
		}
		if length == len(bytes) {
			return length
		}
	}
}
