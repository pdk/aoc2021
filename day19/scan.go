package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		log.Fatalf("program failed: %s\n", err)
	}
	fmt.Printf("fini\n")
}

func run(args []string, stdout io.Writer) error {

	merged, scanners, ok := mergeAll(data[0], data[1:])

	log.Printf("merged is %t, size %d, max manhattan distance is %d", ok, merged.size(), maxDist(scanners))

	return nil
}

func maxDist(pts []pt) int {
	d := 0
	for i := 0; i < len(pts); i++ {
		for j := i + 1; j < len(pts); j++ {
			nd := dist(pts[i], pts[j])
			if nd > d {
				log.Printf("dist %s %s => %d", pts[i], pts[j], nd)
				d = nd
			}
		}
	}
	return d
}

func dist(p1, p2 pt) int {
	return abs(p1.x-p2.x) + abs(p1.y-p2.y) + abs(p1.z-p2.z)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func mergeAll(base scanner, toMerge []scanner) (scanner, []pt, bool) {

	scannerPts := []pt{}

	c := 0
	for len(toMerge) > 0 && c < 100 {
		c++

		next := toMerge[0]

		matching, f, r, p := base.findNMatches(12, next)
		if len(matching) != 0 {
			scannerPts = append(scannerPts, p)
			next = next.rotate(f, r).shift(p)
			base = base.merge(next)
			toMerge = toMerge[1:]
		} else {
			// put this one back on the end
			toMerge = append(toMerge[1:], next)
		}
	}

	return base, scannerPts, len(toMerge) == 0
}
