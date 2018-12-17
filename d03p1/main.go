package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

type MinCount struct {
	Minute, Count int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var dates []string
	s := bufio.NewScanner(f)
	for s.Scan() {
		dates = append(dates, s.Text())
	}

	if err = s.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Slice(dates, func(i, j int) bool {
		return toDate(dates[i]).Before(toDate(dates[j]))

	})

	minutes := map[int]map[int]int{}
	guard := 0
	for i := range dates {
		// fmt.Println(string(d[19]))
		switch dates[i][19] {
		case 'G':
			var a, b, c, d, e int
			fmt.Sscanf(dates[i], "[%d-%d-%d %d:%d] Guard #%d begins shift", &a, &b, &c, &d, &e, &guard)
		case 'f':
			sleep := toDate(dates[i]).Minute()
			awake := toDate(dates[i+1]).Minute()
			for i := sleep; i < awake; i++ {
				// minutes[guard] = toDate(dates[i]).Minute
				if minutes[guard] == nil {
					minutes[guard] = make(map[int]int)
				}
				minutes[guard][i]++
			}
		}
	}

	for guard, mins := range minutes {
		var sum, g, gv int
		for minute, value := range mins {
			sum += value
			if value > gv {
				g = minute
				gv = value
			}
		}
		fmt.Printf("Guard #%d slept %d minutes, sleeping for %d mins at minute #%d\n", guard, sum, gv, g)
	}
}

func toDate(event string) time.Time {
	if event[0] != '[' {
		log.Fatalf("Event is not a valid formatted event: %v", event[0])
	}

	d := event[1:17]

	format := "2006-01-02 15:04"
	date, err := time.Parse(format, d)
	if err != nil {
		log.Fatalf("Cannot convert to date: %v", err)
	}

	return date
}
