package main

import (
	"io"
	"log"
)

func main() {
	// if err := run(os.Args, os.Stdout); err != nil {
	// 	log.Fatalf("program failed: %s\n", err)
	// }

	sliceAndDice(testData)

	// fmt.Printf("fini\n")
}

func sliceAndDice(rules []cmd) {

	// number the rules
	for i := range rules {
		rules[i].ruleNo = i
	}

	xBlades := distinct(xPlains(cubesOf(rules)))
	yBlades := distinct(yPlains(cubesOf(rules)))
	zBlades := distinct(zPlains(cubesOf(rules)))

	log.Printf("xBlades %d", len(xBlades))
	log.Printf("yBlades %d", len(yBlades))
	log.Printf("zBlades %d", len(zBlades))

	sliced := xSlice(cubesOf(rules), xBlades)
	log.Printf("after xSlice %d", len(sliced))

	sliced = ySlice(sliced, yBlades)
	log.Printf("after ySlice %d", len(sliced))

	sliced = zSlice(sliced, yBlades)
	log.Printf("after zSlice %d", len(sliced))

	m := map[cubeDims][]cube{}

	for _, c := range sliced {
		e := m[c.cubeDims()]
		m[c.cubeDims()] = append(e, c)
	}

	log.Printf("distinct cube dims %d", len(m))
}

func cubesOf(rules []cmd) []cube {
	cubes := []cube{}

	for _, r := range rules {
		cubes = append(cubes, r.cube)
	}

	return cubes
}

func run(args []string, stdout io.Writer) error {

	space := []cube{}
	rules := testData

	for i, rule := range rules {

		// if i != 26 &&
		// 	i != 34 &&
		// 	i != 23 &&
		// 	i != 10 &&
		// 	i != 14 &&
		// 	i != 17 &&
		// 	i != 21 &&
		// 	i != 37 {
		// 	continue
		// }

		rule.cube.ruleNo = i
		space = applyRule(space, rule)

		stopIfIntersection(space, i, rules)

		// log.Printf("after rule %d space has %d cuboids", i, len(space))
	}

	totalOn := 0
	for _, c := range space {
		// log.Printf("cuboid %d from rule %d volume %d", i, c.ruleNo, c.volume())
		totalOn += c.volume()
	}

	log.Printf("total on is %d", totalOn)

	return nil
}

func stopIfIntersection(space []cube, ruleNo int, rules []cmd) {
	for i, a := range space {
		for j, b := range space {
			if i == j {
				continue
			}
			if a.intersects(b) {

				log.Printf("last rule %d %v", ruleNo, rules[ruleNo])
				log.Printf("intersection: (%d) %v (%d) %v", i, a, j, b)

				_, inter, _ := splitterate(a, b)
				log.Printf("intersection is %v", inter)

				log.Fatal("abort")
			}
		}
	}
}

var rule2634intersect = cube{xRange: lohi{from: -111166, to: -40997}, yRange: lohi{from: -46592, to: 2688}, zRange: lohi{from: 5609, to: 50954}}

func chkAdd(adding ...cube) {
	for _, c := range adding {
		if c.intersects(rule2634intersect) {
			log.Printf("adding cube that intersects rule 26/34 intersection %v", c)
		}
	}
}

func applyRule(cubeSpace []cube, rule cmd) []cube {

	// log.Printf("applying rule %d to %d incoming cubes", rule.ruleNo, len(cubeSpace))

	ruleCopy := rule

	newSpace := []cube{}

	for _, each := range cubeSpace {

		// check if the incoming cube intersects with any previous cubes
		conflict, conflictFound := each.findIntersect(newSpace)
		if conflictFound {
			log.Printf("rule %d found conflict while processing incoming cubes %v %v", rule.ruleNo, each, conflict)
			log.Printf("source rule %d is %v", conflict.ruleNo, testData[conflict.ruleNo])
		}

		if !rule.intersects(each) {
			if conflictFound {
				log.Printf("and the conflict does NOT interect the current rule")
			}
			newSpace = append(newSpace, each)
			continue
		}

		nonIntersects, _, _ := splitterate(each, rule.cube)

		if each.ruleNo == 26 {
			log.Printf("rule %d creating %d cubelets from r26 %v", rule.ruleNo, len(nonIntersects), each)
		}

		// log.Printf("processing rule %d, intersects rule %d, producing %d cubelets", rule.ruleNo, each.ruleNo, len(nonIntersects))

		// make sure all newly produced cubelets are clean
		for _, cubelet := range nonIntersects {
			if !each.contains(cubelet) {
				log.Printf("cubelet %v is NOT contained within %v", cubelet, each)
			}
			if rule.cube.intersects(cubelet) {
				log.Printf("cubelet %v intersects generator rule %v", cubelet, rule)
			}
			// log.Printf("cubelet %v is contained within %v: %t", cubelet, each, each.contains(cubelet))
			for _, prev := range newSpace {
				if cubelet.intersects(prev) {
					log.Printf("rule %d fresh cubelet intersects with previous cube %v %v", rule.ruleNo, cubelet, prev)
					// log.Printf("previous cube intersects with current cube? %t", prev.intersects(each))
				}
			}
		}

		chkAdd(nonIntersects...)
		newSpace = append(newSpace, nonIntersects...)
	}

	if rule != ruleCopy {
		log.Printf("whoa! rule changed!!!")
	}

	if rule.which == on {

		for _, prev := range newSpace {
			if rule.cube.intersects(prev) {
				log.Printf("want to add new rule cube, but intersects with prev %v %v", rule, prev)
			}
		}

		chkAdd(rule.cube)
		newSpace = append(newSpace, rule.cube)
	}

	return newSpace
}

type cubeDims struct {
	xRange, yRange, zRange lohi
}

type cube struct {
	ruleNo                 int
	xRange, yRange, zRange lohi
}

func (c cube) cubeDims() cubeDims {
	return cubeDims{
		xRange: c.xRange,
		yRange: c.yRange,
		zRange: c.zRange,
	}
}

func (c cube) findIntersect(others []cube) (cube, bool) {
	for _, other := range others {
		if c.intersects(other) {
			return other, true
		}
	}

	return cube{}, false
}

func (c cube) equals(c2 cube) bool {
	return c.xRange == c2.xRange && c.yRange == c2.yRange && c.zRange == c2.zRange
}

func (c cube) volume() int {
	return (1 + c.xRange.to - c.xRange.from) *
		(1 + c.yRange.to - c.yRange.from) *
		(1 + c.zRange.to - c.zRange.from)
}

type pt struct {
	x, y, z int
}

// splitterate returns
// 1. the cubelets that do not intersect (from a)
// 2. the intersection cubelet
// 3. the cubelets that do not intersect (from b)
func splitterate(a, b cube) ([]cube, cube, []cube) {
	if !a.intersects(b) {
		log.Fatalf("splitterate called on non-intersecting cubes %v %v", a, b)
	}

	aSplit, bSplit := xSplit(ySplit(zSplit([]cube{a}, []cube{b})))

	var equalCube cube
	eCount := 0
	for _, a := range aSplit {
		for _, b := range bSplit {
			if a.equals(b) {
				equalCube = a
				eCount++
			}
		}
	}

	if eCount != 1 {
		log.Fatalf("should have gotten 1 equal cubelet, but got %d", eCount)
	}

	return without(equalCube, aSplit), equalCube, without(equalCube, bSplit)
}

func without(x cube, them []cube) []cube {
	ecnt := 0
	t := []cube{}
	for _, p := range them {
		if p.equals(x) {
			ecnt++
			continue
		}
		t = append(t, p)
	}

	if ecnt != 1 {
		log.Fatalf("failed to weed out intersection cube")
	}

	return t
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

func xSlice(cubes []cube, blades []float64) []cube {
	nCubes := []cube{}

	for _, c := range cubes {
		nCubes = append(nCubes, c.xSplit(blades...)...)
	}

	return nCubes
}

func ySlice(cubes []cube, blades []float64) []cube {
	nCubes := []cube{}

	for _, c := range cubes {
		nCubes = append(nCubes, c.ySplit(blades...)...)
	}

	return nCubes
}

func zSlice(cubes []cube, blades []float64) []cube {
	nCubes := []cube{}

	for _, c := range cubes {
		nCubes = append(nCubes, c.zSplit(blades...)...)
	}

	return nCubes
}

func xSplit(a []cube, b []cube) ([]cube, []cube) {
	plains := append(xPlains(a), xPlains(b)...)
	return xSlice(a, plains), xSlice(b, plains)
}

func ySplit(a []cube, b []cube) ([]cube, []cube) {
	plains := append(yPlains(a), yPlains(b)...)
	return ySlice(a, plains), ySlice(b, plains)
}

func zSplit(a []cube, b []cube) ([]cube, []cube) {
	plains := append(zPlains(a), zPlains(b)...)
	return zSlice(a, plains), zSlice(b, plains)
}

func (c cube) xSplit(vals ...float64) []cube {

	for _, val := range vals {
		if float64(c.xRange.from) < val && val < float64(c.xRange.to) {
			c1 := c // copy original
			c1.xRange.to = int(val - 0.5)
			c2 := c // another copy
			c2.xRange.from = int(val + 0.5)

			return append([]cube{c1}, c2.xSplit(vals...)...)
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

			return append([]cube{c1}, c2.ySplit(vals...)...)
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

			return append([]cube{c1}, c2.zSplit(vals...)...)
		}
	}

	return []cube{c}
}

// contains returns true if this completely contains the other.
func (c cube) contains(inner cube) bool {

	outer := c

	return inner.xRange.between(outer.xRange) &&
		inner.yRange.between(outer.yRange) &&
		inner.zRange.between(outer.zRange)
}

// within return true if the point is within the cube.
func (p pt) within(c cube) bool {

	return between(c.xRange.from, p.x, c.xRange.to) &&
		between(c.yRange.from, p.y, c.yRange.to) &&
		between(c.zRange.from, p.z, c.zRange.to)
}

func (c cube) intersects(c2 cube) bool {

	for _, p := range c2.corners() {
		if p.within(c) {
			return true
		}
	}

	for _, p := range c.corners() {
		if p.within(c2) {
			return true
		}
	}

	return false
}

func (c cube) corners() []pt {
	return []pt{
		// front
		{c.xRange.from, c.yRange.from, c.zRange.from},
		{c.xRange.from, c.yRange.to, c.zRange.from},
		{c.xRange.to, c.yRange.from, c.zRange.from},
		{c.xRange.to, c.yRange.to, c.zRange.from},
		// back
		{c.xRange.from, c.yRange.from, c.zRange.to},
		{c.xRange.from, c.yRange.to, c.zRange.to},
		{c.xRange.to, c.yRange.from, c.zRange.to},
		{c.xRange.to, c.yRange.to, c.zRange.to},
	}
}

func isOn(x, y, z int) bool {
	for i := len(data) - 1; i >= 0; i-- {
		e := data[i].effect(x, y, z)
		if e != nada {
			return e == on
		}
	}
	return false
}

type which int

const (
	nada which = iota
	on
	off
)

type lohi struct {
	from, to int
}

// between returns true if both the other from/to are between this from/to.
func (x lohi) between(o lohi) bool {
	return between(o.from, x.from, o.to) && between(o.from, x.to, o.to)
}

// between returns true if b is between a and b.
func between(a, b, c int) bool {
	return a <= b && b <= c
}

type cmd struct {
	which
	// xRange, yRange, zRange lohi
	cube
}

func (c cmd) effect(x, y, z int) which {
	if x >= c.xRange.from && x <= c.xRange.to &&
		y >= c.yRange.from && y <= c.yRange.to &&
		z >= c.zRange.from && z <= c.zRange.to {

		return c.which
	}

	return nada
}
