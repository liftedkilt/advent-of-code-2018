package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point struct {
	X, Y int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	s := bufio.NewScanner(f)

	claims := []int{}
	graph := map[Point][]int{}
	for s.Scan() {
		var id, x, y, w, h int
		_, err := fmt.Sscanf(s.Text(), "#%d @ %d,%d: %dx%d", &id, &x, &y, &w, &h)
		if err != nil {
			log.Fatal(err)
		}

		claims = append(claims, id)

		for i := y; i < (y + h); i++ {
			for j := x; j < (x + w); j++ {
				point := Point{
					X: j,
					Y: i,
				}
				graph[point] = append(graph[point], id)
			}
		}
	}

	if err = s.Err(); err != nil {
		log.Fatal(err)
	}

	contention := map[int]bool{}
	for _, v := range graph {
		if len(v) > 1 {
			for _, claim := range v {
				contention[claim] = true
			}
		}
	}

	for _, claim := range claims {
		if !contention[claim] {
			fmt.Println(claim)
		}
	}

}
