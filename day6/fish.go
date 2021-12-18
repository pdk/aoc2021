package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if err := part1(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	if err := part2(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func part2(args []string, stdout io.Writer) error {

	// have a map of cycle days to fish count
	pops := map[int]int64{}

	for _, f := range startPool() {
		pops[f.cycleDays]++
	}

	for i := 0; i < 256; i++ {
		gen := pops[0]
		pops[0] = pops[1]
		pops[1] = pops[2]
		pops[2] = pops[3]
		pops[3] = pops[4]
		pops[4] = pops[5]
		pops[5] = pops[6]
		pops[6] = pops[7] + gen
		pops[7] = pops[8]
		pops[8] = gen

		var t int64
		for _, c := range pops {
			t += c
		}
		log.Printf("day %d: count = %d, pops = %#v", i, t, pops)
	}

	return nil
}

func part1(args []string, stdout io.Writer) error {

	p := startPool()

	for i := 0; i < 80; i++ {
		p = nextDay(p)
		log.Printf("day %d: pool has %d fish", i, len(p))
	}

	return nil
}

type fish struct {
	cycleDays int
}

func (f fish) nextDay() []fish {
	if f.cycleDays == 0 {
		return []fish{
			fish{6},
			fish{8},
		}
	}

	f.cycleDays--

	return []fish{f}
}

func nextDay(pool []fish) []fish {
	newPool := []fish{}
	for _, f := range pool {
		newPool = append(newPool, f.nextDay()...)
	}
	return newPool
}

func startPool() []fish {
	pool := []fish{}

	for _, d := range []int{
		5, 1, 2, 1, 5, 3, 1, 1, 1, 1, 1, 2, 5, 4, 1, 1, 1, 1, 2, 1, 2, 1, 1, 1, 1,
		1, 2, 1, 5, 1, 1, 1, 3, 1, 1, 1, 3, 1, 1, 3, 1, 1, 4, 3, 1, 1, 4, 1, 1, 1,
		1, 2, 1, 1, 1, 5, 1, 1, 5, 1, 1, 1, 4, 4, 2, 5, 1, 1, 5, 1, 1, 2, 2, 1, 2,
		1, 1, 5, 3, 1, 2, 1, 1, 3, 1, 4, 3, 3, 1, 1, 3, 1, 5, 1, 1, 3, 1, 1, 4, 4,
		1, 1, 1, 5, 1, 1, 1, 4, 4, 1, 3, 1, 4, 1, 1, 4, 5, 1, 1, 1, 4, 3, 1, 4, 1,
		1, 4, 4, 3, 5, 1, 2, 2, 1, 2, 2, 1, 1, 1, 2, 1, 1, 1, 4, 1, 1, 3, 1, 1, 2,
		1, 4, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 1, 1, 5, 5, 1, 1, 1, 5, 1, 1, 1, 1, 5,
		1, 3, 2, 1, 1, 5, 2, 3, 1, 2, 2, 2, 5, 1, 1, 3, 1, 1, 1, 5, 1, 4, 1, 1, 1,
		3, 2, 1, 3, 3, 1, 3, 1, 1, 1, 1, 1, 1, 1, 2, 3, 1, 5, 1, 4, 1, 3, 5, 1, 1,
		1, 2, 2, 1, 1, 1, 1, 5, 4, 1, 1, 3, 1, 2, 4, 2, 1, 1, 3, 5, 1, 1, 1, 3, 1,
		1, 1, 5, 1, 1, 1, 1, 1, 3, 1, 1, 1, 4, 1, 1, 1, 1, 2, 2, 1, 1, 1, 1, 5, 3,
		1, 2, 3, 4, 1, 1, 5, 1, 2, 4, 2, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 4, 1, 5,
	} {
		pool = append(pool, fish{d})
	}

	return pool
}
