package main

import (
	"errors"
	"fmt"
)

type scanner struct {
	name    string
	beacons []pt
	ptMap   map[pt]bool
}

func (s *scanner) ldMap() {
	if s.ptMap != nil {
		return
	}

	s.ptMap = map[pt]bool{}
	for _, p := range s.beacons {
		s.ptMap[p] = true
	}
}

// rotate computes new x,y,z for all beacons.
func (s scanner) rotate(f facing, r rotation) scanner {

	b := []pt{}
	for _, p := range s.beacons {
		b = append(b, p.rotate(f, r))
	}

	return scanner{
		name:    s.name,
		beacons: b,
	}
}

func (s scanner) shift(origin pt) scanner {

	nb := []pt{}

	for _, b := range s.beacons {
		nb = append(nb, b.shift(origin))
	}

	return scanner{
		name:    s.name,
		beacons: nb,
	}
}

func (s scanner) print() {

	fmt.Printf("--- %s ---\n", s.name)
	for _, p := range s.beacons {
		fmt.Printf("%d,%d,%d\n", p.x, p.y, p.z)
	}
	fmt.Printf("\n")

}

// findMatchRotation tries all rotations of b until it matches a.
func findMatchRotation(a, b scanner) (facing, rotation, error) {

	for _, f := range allFacings {
		for _, r := range allRotations {
			rot := b.rotate(f, r)
			if rot.matchAll(a) {
				return f, r, nil
			}
		}
	}

	return forward, zero, errors.New("failed to find rotation to match")
}

func (s *scanner) matchAll(other scanner) bool {
	s.ldMap()

	for _, p := range other.beacons {
		if !s.ptMap[p] {
			return false
		}
	}

	return true
}

func (s *scanner) matches(other scanner) []pt {
	s.ldMap()

	matches := []pt{}

	for _, p := range other.beacons {
		if s.ptMap[p] {
			matches = append(matches, p)
		}
	}

	return matches
}

func (s scanner) findNMatches(n int, other scanner) ([]pt, facing, rotation, pt) {

	for _, f := range allFacings {
		for _, r := range allRotations {

			rotated := other.rotate(f, r)

			for _, sp := range s.beacons {
				for _, op := range rotated.beacons {

					// assume sp & op are same point
					// compute relative shift
					shift := pt{
						x: sp.x - op.x,
						y: sp.y - op.y,
						z: sp.z - op.z,
					}

					// fmt.Printf("sp=%s op=%s shift=%s shifted=%s\n",
					// 	sp.String(), op.String(), shift.String(), op.shift(shift).String())

					chk := rotated.shift(shift)
					matches := s.matches(chk)
					if len(matches) >= n {
						return matches, f, r, shift
					}
				}
			}

		}
	}

	return []pt{}, forward, zero, pt{}
}

func (s scanner) merge(other scanner) scanner {

	n := scanner{
		name:    s.name,
		beacons: []pt{},
		ptMap:   map[pt]bool{},
	}

	for _, x := range s.beacons {
		if !n.ptMap[x] {
			n.ptMap[x] = true
			n.beacons = append(n.beacons, x)
		}
	}

	for _, x := range other.beacons {
		if !n.ptMap[x] {
			n.ptMap[x] = true
			n.beacons = append(n.beacons, x)
		}
	}

	fmt.Printf("merged %s (%d) and %s (%d) => %d\n", s.name, s.size(), other.name, other.size(), n.size())

	return n
}

func (s scanner) size() int {
	return len(s.beacons)
}
