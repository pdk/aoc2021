package main

import (
	"log"
	"testing"
)

func TestXYZSplitting(t *testing.T) {

	a := cube{
		xRange: lohi{1, 5},
		yRange: lohi{1, 5},
		zRange: lohi{1, 5},
	}

	b := cube{
		xRange: lohi{3, 7},
		yRange: lohi{3, 7},
		zRange: lohi{3, 7},
	}

	if !a.intersects(b) {
		t.Errorf("expected a, b to intersect, but no %v %v", a, b)
	}

	aSplit, bSplit := xSplit(ySplit(zSplit([]cube{a}, []cube{b})))

	if len(aSplit) != 8 {
		t.Errorf("expected 8 sub-cubes, got %d", len(aSplit))
	}

	if len(bSplit) != 8 {
		t.Errorf("expected 8 sub-cubes, got %d", len(bSplit))
	}

	// find the matching/intersecting cubes
	interCount := 0
	equalCount := 0
	for _, a := range aSplit {
		for _, b := range bSplit {
			if a.intersects(b) {
				interCount++
				log.Printf("found a/b intersection %v %v", a, b)
			}
			if a.equals(b) {
				equalCount++
			}
		}
	}

	if interCount != 1 {
		t.Errorf("expected exactly 1 intersection, found %d", interCount)
	}
	if equalCount != 1 {
		t.Errorf("expected exactly 1 equal match, found %d", equalCount)
	}

}

func TestRule26Rule34Foo(t *testing.T) {

	rule26 := testData[26]
	rule34 := testData[34]

	t.Errorf("rules 26 intersects 34 %t", rule26.intersects(rule34.cube))
	t.Errorf("rules 34 intersects 26 %t", rule34.intersects(rule26.cube))

	t.Errorf("rules 26 contains 34 %t", rule26.contains(rule34.cube))
	t.Errorf("rules 34 contains 26 %t", rule34.contains(rule26.cube))

	_, intersection1, _ := splitterate(rule26.cube, rule34.cube)
	_, intersection2, _ := splitterate(rule34.cube, rule26.cube)

	t.Errorf("rule26/34 intersection is %#v", intersection1)

	t.Errorf("intersection 1 == 2: %t", intersection1 == intersection2)

	t.Errorf("rule26 contains intersection: %t", rule26.contains(intersection1))
	t.Errorf("rule34 contains intersection: %t", rule34.contains(intersection1))
}

func TestRule26Rule34(t *testing.T) {
	// 2021/12/28 21:52:16 rule 26 {1 {0 {-120100 -32970} {-46592 27473} {-11695 61039}}}
	// 2021/12/28 21:52:16 rule 34 {1 {0 {-111166 -40997} {-71714 2688} {5609 50954}}}

	// intersection is {-111166 -40997} {-46592 2688} {5609 50954}

	errSample := []cube{
		{ruleNo: 26, xRange: lohi{-120100, -98157}, yRange: lohi{-46592, -19988}, zRange: lohi{-8455, 10803}},
		{ruleNo: 26, xRange: lohi{-98156, -95823}, yRange: lohi{-46592, -19988}, zRange: lohi{-8455, 10803}},
		{ruleNo: 26, xRange: lohi{-95822, -77140}, yRange: lohi{-46592, -19988}, zRange: lohi{-80, 10803}},
		{ruleNo: 26, xRange: lohi{-120100, -98157}, yRange: lohi{-19987, -19627}, zRange: lohi{-8455, 10803}},
		{ruleNo: 26, xRange: lohi{-98156, -95823}, yRange: lohi{-19987, -19627}, zRange: lohi{-8455, 10803}},
		{ruleNo: 26, xRange: lohi{-95822, -77140}, yRange: lohi{-19987, -19627}, zRange: lohi{-80, 10803}},
		{ruleNo: 26, xRange: lohi{-120100, -95823}, yRange: lohi{-46592, -19988}, zRange: lohi{10804, 39014}},
		{ruleNo: 26, xRange: lohi{-95822, -77140}, yRange: lohi{-46592, -19988}, zRange: lohi{10804, 39014}},
		{ruleNo: 26, xRange: lohi{-120100, -95823}, yRange: lohi{-19987, -19627}, zRange: lohi{10804, 39014}},
		{ruleNo: 26, xRange: lohi{-120100, -95823}, yRange: lohi{-19987, -19627}, zRange: lohi{10804, 39014}},
		{ruleNo: 26, xRange: lohi{-120100, -95823}, yRange: lohi{-19987, -19627}, zRange: lohi{10804, 39014}},
		{ruleNo: 26, xRange: lohi{-120100, -95823}, yRange: lohi{-19987, -19627}, zRange: lohi{10804, 39014}},
		{ruleNo: 26, xRange: lohi{-120100, -95823}, yRange: lohi{-46592, -19988}, zRange: lohi{39015, 61039}},
		{ruleNo: 26, xRange: lohi{-95822, -77140}, yRange: lohi{-46592, -19988}, zRange: lohi{39015, 59318}},
		{ruleNo: 26, xRange: lohi{-120100, -95823}, yRange: lohi{-19987, -19627}, zRange: lohi{39015, 61039}},
		{ruleNo: 26, xRange: lohi{-120100, -95823}, yRange: lohi{-19987, -19627}, zRange: lohi{39015, 61039}},
		{ruleNo: 34, xRange: lohi{-111166, -95823}, yRange: lohi{-71714, -19988}, zRange: lohi{5609, 10803}},
		{ruleNo: 34, xRange: lohi{-111166, -95823}, yRange: lohi{-71714, -19988}, zRange: lohi{5609, 10803}},
		{ruleNo: 34, xRange: lohi{-95822, -77140}, yRange: lohi{-71714, -19988}, zRange: lohi{5609, 10803}},
		{ruleNo: 34, xRange: lohi{-111166, -95823}, yRange: lohi{-19987, 2688}, zRange: lohi{5609, 10803}},
		{ruleNo: 34, xRange: lohi{-111166, -95823}, yRange: lohi{-19987, 2688}, zRange: lohi{5609, 10803}},
		{ruleNo: 34, xRange: lohi{-95822, -77140}, yRange: lohi{-19987, -18797}, zRange: lohi{5609, 10803}},
		{ruleNo: 34, xRange: lohi{-111166, -95823}, yRange: lohi{-71714, -19988}, zRange: lohi{10804, 50954}},
		{ruleNo: 34, xRange: lohi{-111166, -95823}, yRange: lohi{-71714, -19988}, zRange: lohi{10804, 50954}},
		{ruleNo: 34, xRange: lohi{-95822, -77140}, yRange: lohi{-71714, -19988}, zRange: lohi{10804, 50954}},
		{ruleNo: 34, xRange: lohi{-95822, -77140}, yRange: lohi{-71714, -19988}, zRange: lohi{10804, 50954}},
		{ruleNo: 34, xRange: lohi{-111166, -101087}, yRange: lohi{-19987, -7089}, zRange: lohi{10804, 33934}},
		{ruleNo: 34, xRange: lohi{-101086, -95823}, yRange: lohi{-19987, -7089}, zRange: lohi{10804, 33934}},
		{ruleNo: 34, xRange: lohi{-111166, -101087}, yRange: lohi{-19987, -7089}, zRange: lohi{33935, 50954}},
		{ruleNo: 34, xRange: lohi{-111166, -101087}, yRange: lohi{-19987, -7089}, zRange: lohi{33935, 50954}},
		{ruleNo: 34, xRange: lohi{-101086, -95823}, yRange: lohi{-19987, -7089}, zRange: lohi{33935, 50954}},
		{ruleNo: 34, xRange: lohi{-101086, -95823}, yRange: lohi{-19987, -7089}, zRange: lohi{33935, 50954}},
	}

	// 2021/12/28 21:52:16 rule 26 {1 {0 {-120100 -32970} {-46592 27473} {-11695 61039}}}
	// 2021/12/28 21:52:16 rule 34 {1 {0 {-111166 -40997} {-71714 2688} {5609 50954}}}

	rule26 := cube{
		xRange: lohi{-120100, -32970},
		yRange: lohi{-46592, 27473},
		zRange: lohi{-11695, 61039},
	}

	rule34 := cube{
		xRange: lohi{-111166, -40997},
		yRange: lohi{-71714, 2688},
		zRange: lohi{5609, 50954},
	}

	for i, s := range errSample {
		if s.ruleNo == 26 {
			if rule26.contains(s) {
				t.Errorf("rule 26 DOES contain %d %v", i, s)
			} else {
				t.Errorf("rule 26 does NOT contain %d %v", i, s)
			}
		}
		if s.ruleNo == 34 {
			if rule34.contains(s) {
				t.Errorf("rule 34 DOES contain %d %v", i, s)
			} else {
				t.Errorf("rule 34 does NOT contain %d %v", i, s)
			}
		}
	}

	interCubelet := cube{
		xRange: lohi{-111166, -40997},
		yRange: lohi{-46592, 2688},
		zRange: lohi{5609, 50954},
	}

	for i, s := range errSample {
		if interCubelet.intersects(s) {
			t.Errorf("found intersection with %d (%t, %t): %v", i, interCubelet.contains(s), s.contains(interCubelet), s)
		}
	}

	a := cube{
		xRange: lohi{-120100, -32970},
		yRange: lohi{-46592, 27473},
		zRange: lohi{-11695, 61039},
	}

	b := cube{
		xRange: lohi{-111166, -40997},
		yRange: lohi{-71714, 2688},
		zRange: lohi{5609, 50954},
	}

	log.Printf("a/b intersect() = %t", a.intersects(b))

	aSplit, bSplit := xSplit(ySplit(zSplit([]cube{a}, []cube{b})))

	log.Printf("aSplit len %d", len(aSplit))
	log.Printf("bSplit len %d", len(bSplit))

	for _, a := range aSplit {
		for _, b := range bSplit {
			if a.intersects(b) {
				t.Errorf("aSplit/bSplit intersect %v %v", a, b)
			}
		}
	}

	// for _, a := range aSplit {
	// 	log.Printf("a %v", a)
	// }

	// for _, b := range bSplit {
	// 	log.Printf("b %v", b)
	// }

	t.Errorf("fail")
}

func TestXSplitting(t *testing.T) {

	a := cube{
		xRange: lohi{0, 5},
		yRange: lohi{0, 1},
		zRange: lohi{0, 1},
	}

	b := cube{
		xRange: lohi{3, 7},
		yRange: lohi{0, 1},
		zRange: lohi{0, 1},
	}

	if !a.intersects(b) {
		t.Errorf("expected a & b to intersect, but nope")
	}

	aSplit, bSplit := xSplit([]cube{a}, []cube{b})

	// for _, a := range aSplit {
	// 	log.Printf("a %d - %d", a.xRange.from, a.xRange.to)
	// }

	// for _, b := range bSplit {
	// 	log.Printf("b %d - %d", b.xRange.from, b.xRange.to)
	// }

	if len(aSplit) != 2 {
		t.Errorf("expected 2 cubes in aSplit, but got %d", len(aSplit))
	}

	if len(bSplit) != 2 {
		t.Errorf("expected 2 cubes in bSplit, but got %d", len(bSplit))
	}

	if aSplit[1] != bSplit[0] {
		t.Errorf("expected aSplit[1] == bSplit[0], but got %v %v", aSplit[1], bSplit[0])
	}

	r1, r2, r3 := aSplit[0], aSplit[1], bSplit[1]
	if r1.intersects(r2) || r1.intersects(r3) || r2.intersects(r3) {
		t.Errorf("unexpected intersection %v %v %v", r1, r2, r3)
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

	// t.Errorf("plus/minus is %f/%f", float64(m)-0.5, float64(m)+0.5)
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
