package main

import (
	"fmt"
	"testing"
)

func TestCombineTestData(t *testing.T) {

	merged, _, ok := mergeAll(testData[0], testData[1:])

	if !ok {
		t.Errorf("failed to merge all. expected ok=true, got ok=false")
	}
	if merged.size() != 79 {
		t.Errorf("failed to merge all. expected 79 beacons, got %d", merged.size())
	}
}

func TestMaxDist(t *testing.T) {

	points := []pt{
		{1105, -1205, 1229},
		{-92, -2380, -20},
	}

	d := maxDist(points)

	if d != 3621 {
		t.Errorf("expected distance %d, got %d", 3621, d)
	}
}

func TestMatchZeroAndOne(t *testing.T) {

	matching, f, r, p := testData[0].findNMatches(12, testData[1])

	if len(matching) == 0 {
		t.Error("failed to find match")
		return
	}

	exp := pt{68, -1246, -43}
	if p != exp {
		t.Errorf("expected shift to be %s, but got %s", exp.String(), p.String())
	}

	fmt.Printf("found matches at %s, %s, %s\n", f.String(), r.String(), p.String())

	scanner{
		name:    "matches",
		beacons: matching,
	}.print()
}

func TestMatchFourAndOne(t *testing.T) {

	scanner1 := testData[1].rotate(backward, zero).shift(pt{68, -1246, -43})
	scanner4 := testData[4]

	matching, f, r, p := scanner1.findNMatches(12, scanner4)

	if len(matching) == 0 {
		t.Error("failed to find match")
		return
	}

	fmt.Printf("found matches at %s, %s, %s\n", f.String(), r.String(), p.String())

	scanner{
		name:    "matches",
		beacons: matching,
	}.print()
}

func TestRotate(t *testing.T) {

	start := scanner{
		name: "scanner 0",
		beacons: []pt{
			{-1, -1, 1},
			{-2, -2, 2},
			{-3, -3, 3},
			{-2, -3, 1},
			{5, 6, -4},
			{8, 0, 7},
		},
	}

	variations := []scanner{
		{
			name: "scanner 0",
			beacons: []pt{
				{1, -1, 1},
				{2, -2, 2},
				{3, -3, 3},
				{2, -1, 3},
				{-5, 4, -6},
				{-8, -7, 0},
			},
		},
		{
			name: "scanner 0",
			beacons: []pt{
				{-1, -1, -1},
				{-2, -2, -2},
				{-3, -3, -3},
				{-1, -3, -2},
				{4, 6, 5},
				{-7, 0, 8},
			},
		},
		{
			name: "scanner 0",
			beacons: []pt{
				{1, 1, -1},
				{2, 2, -2},
				{3, 3, -3},
				{1, 3, -2},
				{-4, -6, 5},
				{7, 0, 8},
			},
		},
		{
			name: "scanner 0",
			beacons: []pt{
				{1, 1, 1},
				{2, 2, 2},
				{3, 3, 3},
				{3, 1, 2},
				{-6, -4, -5},
				{0, 7, -8},
			},
		},
	}

	for i, v := range variations {
		f, r, err := findMatchRotation(start, v)
		if err != nil {
			t.Errorf("failed to find matching rotation for variationi %d", i)
			continue
		}

		fmt.Printf("found rotation %s, %s to match variation %d\n", f.String(), r.String(), i)
	}
}
