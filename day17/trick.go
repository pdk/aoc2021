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

	// t := 0
	// for i := 0; i < 100; i++ {
	// 	log.Printf("i=%d t=%d", i, t)
	// 	t += i
	// 	if t > 171 {
	// 		break
	// 	}
	// }

	// target area: x=137..171, y=-98..-73
	target := newArea(137, 171, -98, -73)
	hitCount := 0
	tries := 0

	for tryY := 100; tryY > -100; tryY-- {
		for tryX := 0; tryX <= 200; tryX++ {
			tries++

			p := pos{0, 0}
			v := velocity{tryX, tryY}

			for i := 0; i < 200000; i++ {

				p, v = step(p, v)

				if isBeyond(p, target) {
					break
				}

				if isWithin(p, target) {
					hitCount++
					log.Printf("hit!!!  i=%6d p=%3d,%3d v=%3d,%3d", i, p.x, p.y, v.x, v.y)
					break
				}
			}
		}
	}

	log.Printf("hit count %d (tries %d)", hitCount, tries)

	return nil
}

type area struct {
	nw, se pos
}

func newArea(x1, x2, y1, y2 int) area {
	return area{
		nw: pos{
			min(x1, x2),
			max(y1, y2),
		},
		se: pos{
			max(x1, x2),
			min(y1, y2),
		},
	}
}

type pos struct {
	x, y int
}

type velocity struct {
	x, y int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func isWithin(p pos, a area) bool {

	return p.x >= a.nw.x &&
		p.y <= a.nw.y &&
		p.x <= a.se.x &&
		p.y >= a.se.y

}

func isBeyond(p pos, a area) bool {
	return p.y < a.se.y || p.x > a.se.x
}

// step computes a new position and velocity
func step(p pos, v velocity) (pos, velocity) {

	p.x += v.x
	p.y += v.y

	v.x += drag(v.x)
	v.y-- // gravity

	return p, v
}

func drag(x int) int {
	if x > 0 {
		return -1
	}
	if x < 0 {
		return 1
	}
	return 0
}
