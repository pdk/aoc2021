package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(args []string, stdout io.Writer) error {

	m := getMap()

	sumRisks := 0

	for i := 0; i < len(m); i++ {
		// log.Printf("%v", m[i])
		r := m[i]
		for j := 0; j < len(r); j++ {
			if lowerThan(m[i][j], adjacent(m, i, j)) {
				log.Printf("found low point %d, %d: risk level %d", i, j, m[i][j]+1)
				sumRisks += m[i][j] + 1
			}
		}
	}

	log.Printf("sum of risks %d", sumRisks)

	return nil
}

// lowerThan returns true if the value is less than all the (provided) adjacent
// values.
func lowerThan(v int, ad []int) bool {
	for _, x := range ad {
		if v >= x {
			return false
		}
	}
	return true
}

// adjacent returns the heights of the adjacent locations
func adjacent(m [][]int, i, j int) []int {

	r := []int{}

	if i > 0 {
		r = append(r, m[i-1][j])
	}
	if j > 0 {
		r = append(r, m[i][j-1])
	}
	if j+1 < len(m[i]) {
		r = append(r, m[i][j+1])
	}
	if i+1 < len(m) {
		r = append(r, m[i+1][j])
	}

	return r
}

func getMap() [][]int {

	m := [][]int{}

	for _, line := range mapStrings {
		mapRow := []int{}
		nums := strings.Split(line, "")
		for _, n := range nums {
			v, _ := strconv.Atoi(n)
			mapRow = append(mapRow, v)
		}
		m = append(m, mapRow)
	}

	return m
}
