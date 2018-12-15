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

	nums := []int{}

	s := bufio.NewScanner(f)
	for s.Scan() {
		var n int
		_, err := fmt.Sscanf(s.Text(), "%d", &n)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, n)
	}

	if err = s.Err(); err != nil {
		log.Fatal(err)
	}

	sum := 0
	seen := map[int]bool{}
	for {
		for _, n := range nums {
			sum += n

			if seen[sum] == true {
				fmt.Println(sum)
				return
			}
			seen[sum] = true
		}
	}

}
