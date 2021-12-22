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
	moves     []pt
}

var (
	exploredPoints = map[pt]bool{
		{0, 0}: true,
	}
	exploredPaths = []path{
		{
			moves: []pt{
				{0, 0},
			},
		},
	}
	dirs = []dir{north, south, east, west}
)

func run(args []string, stdout io.Writer) error {

	for i := 0; i < 10000; i++ {
		path, move, risk := findNextShortestPath()
		exploredPaths = append(exploredPaths, newPath(path, move, risk))
		exploredPoints[move] = true

		if move.x == 99 && move.y == 99 {
			break
		}
	}

	logPaths()

	return nil
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

func newPath(fromPath int, move pt, risk int) path {

	// make a copy of the slice so we don't go messing up existing paths.
	moves := append([]pt{}, exploredPaths[fromPath].moves...)
	moves = append(moves, move)

	return path{
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
		lastMove := p.moves[len(p.moves)-1]
		for _, d := range dirs {
			candidatePt, candidateRisk := getNextRisk(lastMove, d)
			if candidateRisk >= 0 && p.totalRisk+candidateRisk < leastRisk {
				nextPath = i
				leastRisk = p.totalRisk + candidateRisk
				nextPoint = candidatePt
			}
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
	if exploredPoints[pt{x, y}] || x < 0 || x >= len(risks[0]) || y < 0 || y >= len(risks) {
		return -1
	}

	c := risks[y][x]

	return int(c - '0')
}
