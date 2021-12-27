package main

import "fmt"

type pt struct {
	x, y, z int
}

func (p pt) String() string {
	return fmt.Sprintf("(%d,%d,%d)", p.x, p.y, p.z)
}

// shift computes a new x,y,z such that origin will be the new 0,0,0.
func (p pt) shift(origin pt) pt {

	p.x += origin.x
	p.y += origin.y
	p.z += origin.z

	return p
}

// rotate computes a new x,y,z for a new facing and rotation.
func (p pt) rotate(f facing, r rotation) pt {

	switch f {
	case forward:
		// no change
	case backward:
		// rotate 180 around y
		p.x, p.z = -p.x, -p.z
	case north:
		// rotate 90 around x
		p.y, p.z = -p.z, p.y
	case south:
		// rotate -90 around x
		p.y, p.z = p.z, -p.y
	case east:
		// rotate 90 around y
		p.x, p.z = -p.z, p.x
	case west:
		// rotate -90 around y
		p.x, p.z = p.z, -p.x
	}

	switch r {
	case zero:
		// no change
	case oneQuarter:
		p.x, p.y = p.y, -p.x
	case twoQuarter:
		p.x, p.y = -p.x, -p.y
	case threeQuarter:
		p.x, p.y = -p.y, p.x
	}

	return p
}
