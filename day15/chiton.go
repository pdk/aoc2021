package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		log.Fatalf("program failed: %s\n", err)
	}
	fmt.Printf("fini\n")
}

type dir int

const (
	north dir = iota
	south
	east
	west
)

type pt struct {
	x, y int
}

type path struct {
	totalRisk int
	deadEnd   bool
	moves     []pt
}

var (
	exploredPoints = map[pt]bool{
		{0, 0}: true,
	}
	exploredPaths = []*path{
		{
			moves: []pt{
				{0, 0},
			},
		},
	}
	dirs = []dir{north, south, east, west}
)

func run(args []string, stdout io.Writer) error {

	for i := 0; i < 1000000; i++ {
		path, move, risk := findNextShortestPath()
		exploredPaths = append(exploredPaths, newPath(path, move, risk))
		exploredPoints[move] = true

		if move.x == 499 && move.y == 499 {
			break
		}

		if i%1000 == 0 {
			log.Printf("%d paths, %d dead ends", len(exploredPaths), deadEndCount())
		}
	}

	logPaths()

	return nil
}

func deadEndCount() int {
	count := 0
	for _, e := range exploredPaths {
		if e.deadEnd {
			count++
		}
	}
	return count
}
func logPaths() {
	for _, p := range exploredPaths[len(exploredPaths)-100:] {
		log.Printf("risk %4d %v", p.totalRisk, movesString(p.moves))
	}
	log.Printf("explored %d paths", len(exploredPaths))
}

func movesString(moves []pt) string {

	r := []string{}
	for _, m := range moves {
		r = append(r, fmt.Sprintf("(%d,%d)", m.x, m.y))
	}
	return strings.Join(r, "->")
}

func newPath(fromPath int, move pt, risk int) *path {

	// make a copy of the slice so we don't go messing up existing paths.
	moves := append([]pt{}, exploredPaths[fromPath].moves...)
	moves = append(moves, move)

	return &path{
		totalRisk: risk,
		moves:     moves,
	}
}

// findNextShortestPath considers all potential next moves, identifies which
// route (returned as int) to extend, and which next point to move to from
// there.
func findNextShortestPath() (int, pt, int) {

	nextPath := 0
	nextPoint := pt{}
	leastRisk := math.MaxInt

	for i, p := range exploredPaths {
		if p.deadEnd {
			continue
		}
		lastMove := p.moves[len(p.moves)-1]
		noMoveCount := 0
		for _, d := range dirs {
			candidatePt, candidateRisk := getNextRisk(lastMove, d)
			if candidateRisk == -1 {
				noMoveCount++
				continue
			}
			if p.totalRisk+candidateRisk < leastRisk {
				nextPath = i
				leastRisk = p.totalRisk + candidateRisk
				nextPoint = candidatePt
			}
		}
		if noMoveCount == 4 {
			p.deadEnd = true
		}
	}

	return nextPath, nextPoint, leastRisk
}

func getNextRisk(p pt, d dir) (pt, int) {
	switch d {
	case north:
		return pt{p.x, p.y - 1}, getRisk(p.x, p.y-1)
	case south:
		return pt{p.x, p.y + 1}, getRisk(p.x, p.y+1)
	case east:
		return pt{p.x + 1, p.y}, getRisk(p.x+1, p.y)
	case west:
		return pt{p.x - 1, p.y}, getRisk(p.x-1, p.y)
	}

	return pt{}, -1
}

func getRisk(x, y int) int {
	if exploredPoints[pt{x, y}] || x < 0 || x >= 500 || y < 0 || y >= 500 {
		return -1
	}

	r := int(risks[y%100][x%100] - '0')

	for ; x > 100; x -= 100 { // so hack
		switch r {
		case 9:
			r = 1
		default:
			r++
		}
	}

	for ; y > 100; y -= 100 {
		switch r {
		case 9:
			r = 1
		default:
			r++
		}
	}

	return r
}
