package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input2.txt")
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

	// fmt.Println(polymer)

	bytes := []byte(polymer)

	reactPolymer(bytes)
}

func reactPolymer(bytes []byte) {
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
			fmt.Println(length)
			break
		}
	}
}
