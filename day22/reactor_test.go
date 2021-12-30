package main

import (
	"testing"
)

func TestIntersect(t *testing.T) {

	a := cube{which: on, xRange: lohi{1, 10}, yRange: lohi{6, 6}, zRange: lohi{1, 1}}
	b := cube{which: off, xRange: lohi{4, 4}, yRange: lohi{1, 10}, zRange: lohi{1, 1}}

	if !a.intersects(b) {
		t.Errorf("expected a,b to intersect, but didn't")
	}
}

func TestDataOrdering(t *testing.T) {

	for _, each := range testData {
		checkOrder(t, each.xRange)
		checkOrder(t, each.yRange)
		checkOrder(t, each.zRange)
	}

	for _, each := range data {
		checkOrder(t, each.xRange)
		checkOrder(t, each.yRange)
		checkOrder(t, each.zRange)
	}
}

func checkOrder(t *testing.T, r lohi) {

	if r.from >= r.to {
		t.Errorf("hilo %d,%d is out of order", r.from, r.to)
	}
}

func TestNumbers(t *testing.T) {

	nums := []int{}

	for _, each := range data {
		nums = append(nums, maxInt(
			each.xRange.from, each.xRange.to,
			each.yRange.from, each.yRange.to,
			each.zRange.from, each.zRange.to))
	}

	m := maxInt(nums...)
	if m != 98790 {
		t.Errorf("max is %d", m)
	}
}

func maxInt(ints ...int) int {
	m := 0
	for _, i := range ints {
		i = abs(i)
		if i > m {
			m = i
		}
	}
	return m
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
