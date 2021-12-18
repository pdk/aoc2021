package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// if err := part1(os.Args, os.Stdout); err != nil {
	// 	fmt.Fprintf(os.Stderr, "%s\n", err)
	// 	os.Exit(1)
	// }
	if err := part2(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

type point struct {
	x, y int
}

type location struct {
	basin  point
	height int
}

var basins [100][100]location

func loadBasins() {
	for i, row := range getMap() {
		for j, height := range row {
			basins[i][j].height = height
			// use -1 as marker that basin for this location has not been identified
			basins[i][j].basin.x = -1
			basins[i][j].basin.y = -1
		}
	}
}

func findAndSetBasins() {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if basins[i][j].basin.x >= 0 || // basin already set
				basins[i][j].height == 9 { // too high, not in any basin
				continue
			}

			included := findAllFrom(point{i, j}, map[point]bool{})
			nwPoint := mostNorthWestOf(included)
			for p := range included {
				basins[p.x][p.y].basin = nwPoint
			}
		}
	}
}

func mostNorthWestOf(included map[point]bool) point {
	p := point{1000, 1000} // 0,0 is most NW point. everything is NW of 1000,1000
	for i := range included {
		if i.y < p.y || (i.y == p.y && i.x < p.x) {
			p = i
		}
	}
	return p
}

func findAllFrom(p point, included map[point]bool) map[point]bool {

	if included[p] ||
		p.x < 0 || p.x >= 100 ||
		p.y < 0 || p.y >= 100 ||
		basins[p.x][p.y].height >= 9 {

		return included
	}

	included[p] = true

	included = findAllFrom(point{p.x - 1, p.y}, included)
	included = findAllFrom(point{p.x + 1, p.y}, included)
	included = findAllFrom(point{p.x, p.y - 1}, included)
	included = findAllFrom(point{p.x, p.y + 1}, included)

	return included
}

func printBasins() {

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if basins[i][j].basin.x == -1 {
				// fmt.Printf("      ")
				fmt.Printf(" ")
				continue
			}
			// fmt.Printf("%2d,%2d ", basins[i][j].basin.x, basins[i][j].basin.y)
			fmt.Printf("%c", (basins[i][j].basin.x+basins[i][j].basin.y)%26+'a')
		}
		fmt.Print("\n")
	}
}

func countBasinSizes() []int {

	counts := map[point]int{}

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			counts[basins[i][j].basin]++
		}
	}

	delete(counts, point{-1, -1})

	sizes := []int{}
	for _, size := range counts {
		sizes = append(sizes, size)
	}

	sort.Ints(sizes)

	return sizes
}

func part2(args []string, stdout io.Writer) error {
	loadBasins()
	findAndSetBasins()
	printBasins()

	basinSizes := countBasinSizes()
	log.Printf("basin sizes %v", basinSizes)

	s1 := basinSizes[len(basinSizes)-1]
	s2 := basinSizes[len(basinSizes)-2]
	s3 := basinSizes[len(basinSizes)-3]

	log.Printf("%d x %d x %d = %d", s1, s2, s3, s1*s2*s3)

	return nil
}

func part1(args []string, stdout io.Writer) error {

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
