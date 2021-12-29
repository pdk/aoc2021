package main

import (
	"log"
)

func main() {

	sliceAndDice(data)

}

func sliceAndDice(rules []cube) {

	// number the rules
	for i := range rules {
		rules[i].ruleNo = i
	}

	// printCubes(rules)

	log.Printf("slice and dice %d cubes, total volume %d, %d", len(rules), sumOnVolumes(rules), sumOffVolumes(rules))

	xBlades := distinct(xPlains(rules))
	yBlades := distinct(yPlains(rules))
	zBlades := distinct(zPlains(rules))

	// printSlices(rules[0], "rule 0", xBlades, yBlades, zBlades)
	// printSlices(rules[1], "rule 1", xBlades, yBlades, zBlades)
	// printSlices(rules[2], "rule 2", xBlades, yBlades, zBlades)

	log.Printf("xBlades %d", len(xBlades))
	log.Printf("yBlades %d", len(yBlades))
	log.Printf("zBlades %d", len(zBlades))

	sliced := xSlice(xBlades, rules)
	log.Printf("after x slicing %d cubes, total volume %d, %d", len(sliced), sumOnVolumes(sliced), sumOffVolumes(sliced))

	sliced = ySlice(yBlades, sliced)
	log.Printf("after y slicing %d cubes, total volume %d, %d", len(sliced), sumOnVolumes(sliced), sumOffVolumes(sliced))

	sliced = zSlice(zBlades, sliced)
	log.Printf("after z slicing %d cubes, total volume %d, %d", len(sliced), sumOnVolumes(sliced), sumOffVolumes(sliced))

	m := map[cubeDims][]cube{}

	for i, c := range sliced {
		if i%1000 == 0 {
			log.Printf("mapping %d", i)
		}
		e := m[c.cubeDims()]
		m[c.cubeDims()] = append(e, c)
	}

	log.Printf("distinct cube dims %d", len(m))

	// make sure none of the sliced cubes intersect
	// allCubeDims := []cube{}
	// for _, v := range m {
	// 	allCubeDims = append(allCubeDims, v[0])
	// }
	// for i, a := range allCubeDims {
	// 	for j, b := range allCubeDims {
	// 		if i == j {
	// 			continue
	// 		}
	// 		if a.intersects(b) {
	// 			log.Printf("found intersection %v %v", a, b)
	// 		}
	// 	}
	// }

	totOn, totOff, tot := 0, 0, 0
	for k := range m {
		_, on, off, v := volume(m[k])

		// log.Printf("volume for %v (%d): %d %d %d", c[0], rno, on, off, v)

		totOn += on
		totOff += off
		tot += v
	}

	log.Printf("totOn %d, totOff %d, tot %d", totOn, totOff, tot)
}

func sumOnVolumes(cubes []cube) int {
	s := 0
	for _, c := range cubes {
		if c.which == on {
			s += c.volume()
		}
	}
	return s
}

func sumOffVolumes(cubes []cube) int {
	s := 0
	for _, c := range cubes {
		if c.which == off {
			s += c.volume()
		}
	}
	return s
}

// return max rule number, on volume, off volume, volume
func volume(cubes []cube) (int, int, int, int) {
	maxRule := -1
	w := on
	v := cubes[0].volume()

	for _, c := range cubes {
		if c.ruleNo > maxRule {
			w = c.which
			maxRule = c.ruleNo
		}
	}

	if w == on {
		return maxRule, v, 0, v
	}
	return maxRule, 0, v, v
}

type which int

const (
	on which = iota
	off
)

type lohi struct {
	from, to int
}

// cubeDims used for index of map
type cubeDims struct {
	xRange, yRange, zRange lohi
}

type cube struct {
	ruleNo int
	which
	xRange, yRange, zRange lohi
}

func (c cube) cubeDims() cubeDims {
	return cubeDims{
		xRange: c.xRange,
		yRange: c.yRange,
		zRange: c.zRange,
	}
}

func (c cube) volume() int {
	return (1 + c.xRange.to - c.xRange.from) *
		(1 + c.yRange.to - c.yRange.from) *
		(1 + c.zRange.to - c.zRange.from)
}

func distinct(input []float64) []float64 {
	plains := []float64{}
	m := map[float64]bool{}

	for _, p := range input {
		if !m[p] {
			m[p] = true
			plains = append(plains, p)
		}
	}

	return plains
}

func xPlains(cubes []cube) []float64 {
	plains := []float64{}

	for _, e := range cubes {
		p1 := float64(e.xRange.from) - 0.5
		p2 := float64(e.xRange.to) + 0.5
		plains = append(plains, p1, p2)
	}

	return plains
}

func yPlains(cubes []cube) []float64 {
	plains := []float64{}

	for _, e := range cubes {
		p1 := float64(e.yRange.from) - 0.5
		p2 := float64(e.yRange.to) + 0.5
		plains = append(plains, p1, p2)
	}

	return plains
}

func zPlains(cubes []cube) []float64 {
	plains := []float64{}

	for _, e := range cubes {
		p1 := float64(e.zRange.from) - 0.5
		p2 := float64(e.zRange.to) + 0.5
		plains = append(plains, p1, p2)
	}

	return plains
}

func xSlice(blades []float64, cubes []cube) []cube {
	nCubes := []cube{}

	for _, c := range cubes {
		nCubes = append(nCubes, c.xSplit(blades...)...)
	}

	return nCubes
}

func ySlice(blades []float64, cubes []cube) []cube {
	nCubes := []cube{}

	for _, c := range cubes {
		nCubes = append(nCubes, c.ySplit(blades...)...)
	}

	return nCubes
}

func zSlice(blades []float64, cubes []cube) []cube {
	nCubes := []cube{}

	for _, c := range cubes {
		nCubes = append(nCubes, c.zSplit(blades...)...)
	}

	return nCubes
}

func (c cube) xSplit(vals ...float64) []cube {

	for _, val := range vals {
		if float64(c.xRange.from) < val && val < float64(c.xRange.to) {
			c1 := c // copy original
			c1.xRange.to = int(val - 0.5)
			c2 := c // another copy
			c2.xRange.from = int(val + 0.5)

			// log.Printf("split on x %f between %d - %d, creating %d - %d and %d - %d",
			// 	val, c.xRange.from, c.xRange.to,
			// 	c1.xRange.from, c1.xRange.to,
			// 	c2.xRange.from, c2.xRange.to)

			return append(c1.xSplit(vals...), c2.xSplit(vals...)...)
		}
	}

	return []cube{c}
}

func (c cube) ySplit(vals ...float64) []cube {

	for _, val := range vals {
		if float64(c.yRange.from) < val && val < float64(c.yRange.to) {
			c1 := c // copy original
			c1.yRange.to = int(val - 0.5)
			c2 := c // another copy
			c2.yRange.from = int(val + 0.5)

			return append(c1.ySplit(vals...), c2.ySplit(vals...)...)
		}
	}

	return []cube{c}
}

func (c cube) zSplit(vals ...float64) []cube {

	for _, val := range vals {
		if float64(c.zRange.from) < val && val < float64(c.zRange.to) {
			c1 := c // copz original
			c1.zRange.to = int(val - 0.5)
			c2 := c // another copz
			c2.zRange.from = int(val + 0.5)

			return append(c1.zSplit(vals...), c2.zSplit(vals...)...)
		}
	}

	return []cube{c}
}

// type pt struct {
// 	x, y, z int
// }

// // between returns true if b is between a and b.
// func between(a, b, c int) bool {
// 	return a <= b && b <= c
// }

// // within return true if the point is within the cube.
// func (p pt) within(c cube) bool {

// 	return between(c.xRange.from, p.x, c.xRange.to) &&
// 		between(c.yRange.from, p.y, c.yRange.to) &&
// 		between(c.zRange.from, p.z, c.zRange.to)
// }

// func (c cube) intersects(c2 cube) bool {

// 	for _, p := range c2.corners() {
// 		if p.within(c) {
// 			return true
// 		}
// 	}

// 	for _, p := range c.corners() {
// 		if p.within(c2) {
// 			return true
// 		}
// 	}

// 	return false
// }

// func (c cube) corners() []pt {
// 	return []pt{
// 		// front
// 		{c.xRange.from, c.yRange.from, c.zRange.from},
// 		{c.xRange.from, c.yRange.to, c.zRange.from},
// 		{c.xRange.to, c.yRange.from, c.zRange.from},
// 		{c.xRange.to, c.yRange.to, c.zRange.from},
// 		// back
// 		{c.xRange.from, c.yRange.from, c.zRange.to},
// 		{c.xRange.from, c.yRange.to, c.zRange.to},
// 		{c.xRange.to, c.yRange.from, c.zRange.to},
// 		{c.xRange.to, c.yRange.to, c.zRange.to},
// 	}
// }

// func printCubes(cubes []cube) {

// 	g := newPGrid()

// 	for i, c := range cubes {
// 		for x := c.xRange.from; x <= c.xRange.to; x++ {
// 			for y := c.yRange.from; y <= c.yRange.to; y++ {
// 				g.put(x, y, i)
// 			}
// 		}
// 	}

// 	g.print()
// }

// type pGrid [8][8]int

// func newPGrid() pGrid {
// 	x := [8]int{-1, -1, -1, -1, -1, -1, -1, -1}
// 	return pGrid{x, x, x, x, x, x, x, x}
// }

// func (g pGrid) print() {
// 	for y := 7; y >= 0; y-- {
// 		for x := 0; x < 8; x++ {
// 			if g[x][y] > -1 {
// 				fmt.Printf(" %2d", g[x][y])
// 			} else {
// 				fmt.Printf(" ..")
// 			}
// 		}
// 		fmt.Printf("\n")
// 	}
// 	fmt.Printf("- - - - - - - - - - - - - \n")
// }

// func (g *pGrid) put(x, y, v int) {
// 	g[x][y] = v
// }

// func printSlices(rule cube, name string, xBlades, yBlades, zBlades []float64) {
// 	sliced := zSlice(zBlades, ySlice(yBlades, xSlice(xBlades, []cube{rule})))
// 	fmt.Printf("%s sliced into %d (%d blades): \n", name, len(sliced), len(xBlades)+len(yBlades)+len(zBlades))
// 	printCubes(sliced)
// }
