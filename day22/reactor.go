package main

import (
	"log"
)

func main() {

	rules := data

	// number the rules
	for i := range rules {
		rules[i].ruleNo = i
	}

	totOn := 0
	for i := 0; i < len(rules); i++ {
		v := computeOnVolume(rules[i], i, rules)
		log.Printf("rule %d has volume %d", i, v)
		totOn += v
	}
	log.Printf("total on is %d", totOn)
}

// computeVolume computes the volume of the given rule/cube, but subtracts out
// what was set by previous rules.
func computeOnVolume(rule cube, ruleNo int, rules []cube) int {

	if rule.invalid() || rule.which == off {
		return 0
	}

	for i := ruleNo - 1; i >= 0; i-- {
		if rule.intersects(rules[i]) {
			except := rules[i]
			return computeOnVolumeParts(rule, except, i, rules)
		}
	}

	// Reached bottom. There are no more previous rules which might have already
	// been ON for this cubelet. Now we can go forward looking for OFFs that may
	// have covered this.

	return subtractOffVolume(rule, ruleNo, rules)
}

func subtractOffVolume(rule cube, ruleNo int, rules []cube) int {

	if rule.invalid() {
		return 0
	}

	for i := ruleNo + 1; i < len(rules); i++ {
		if rules[i].which == off && rule.intersects(rules[i]) {
			offCube := rules[i]
			return subtractOffVolumeParts(rule, offCube, i, rules)
		}
	}

	// woot. there are no OFF cubes blocking.
	return rule.volume()
}

func subtractOffVolumeParts(rule, offCube cube, ruleNo int, rules []cube) int {

	// the "subtraction" is NOT including the volume of the intersect returned
	// from the breakup.

	_, outliers := breakUp(rule, offCube)

	totvol := 0
	for _, out := range outliers {
		totvol += subtractOffVolume(out, ruleNo, rules)
	}

	return totvol
}

func computeOnVolumeParts(rule, except cube, ruleNo int, rules []cube) int {

	intersect, outliers := breakUp(rule, except)

	totvol := 0
	for _, out := range outliers {
		totvol += computeOnVolume(out, ruleNo, rules)
	}

	if except.which == off {
		// claim volume that this rule turned ON
		return subtractOffVolume(intersect, ruleNo, rules) + totvol
	}

	return totvol
}

func breakUp(rule, except cube) (intersection cube, outliers []cube) {

	left := cube{
		xRange: lohi{rule.xRange.from, min(rule.xRange.to, except.xRange.from-1)},
		yRange: lohi{rule.yRange.from, rule.yRange.to},
		zRange: lohi{rule.zRange.from, rule.zRange.to},
	}

	right := cube{
		xRange: lohi{max(rule.xRange.from, except.xRange.to+1), rule.xRange.to},
		yRange: lohi{rule.yRange.from, rule.yRange.to},
		zRange: lohi{rule.zRange.from, rule.zRange.to},
	}

	front := cube{
		xRange: lohi{max(rule.xRange.from, except.xRange.from), min(rule.xRange.to, except.xRange.to)},
		yRange: lohi{rule.yRange.from, rule.yRange.to},
		zRange: lohi{rule.zRange.from, min(rule.zRange.to, except.zRange.from-1)},
	}

	back := cube{
		xRange: lohi{max(rule.xRange.from, except.xRange.from), min(rule.xRange.to, except.xRange.to)},
		yRange: lohi{rule.yRange.from, rule.yRange.to},
		zRange: lohi{max(rule.zRange.from, except.zRange.to+1), rule.zRange.to},
	}

	top := cube{
		xRange: lohi{max(rule.xRange.from, except.xRange.from), min(rule.xRange.to, except.xRange.to)},
		yRange: lohi{max(rule.yRange.from, except.yRange.to+1), rule.yRange.to},
		zRange: lohi{max(rule.zRange.from, except.zRange.from), min(rule.zRange.to, except.zRange.to)},
	}

	bottom := cube{
		xRange: lohi{max(rule.xRange.from, except.xRange.from), min(rule.xRange.to, except.xRange.to)},
		yRange: lohi{rule.yRange.from, min(rule.yRange.to, except.yRange.from-1)},
		zRange: lohi{max(rule.zRange.from, except.zRange.from), min(rule.zRange.to, except.zRange.to)},
	}

	intersect := rule.intersection(except)

	return intersect, []cube{
		left, right, front, back, top, bottom,
	}
}

func (c cube) intersection(other cube) cube {

	return cube{
		xRange: lohi{max(c.xRange.from, other.xRange.from), min(c.xRange.to, other.xRange.to)},
		yRange: lohi{max(c.yRange.from, other.yRange.from), min(c.yRange.to, other.yRange.to)},
		zRange: lohi{max(c.zRange.from, other.zRange.from), min(c.zRange.to, other.zRange.to)},
	}
}

type which int

const (
	on which = iota
	off
)

type lohi struct {
	from, to int
}

type cube struct {
	ruleNo int
	which
	xRange, yRange, zRange lohi
}

func (c cube) volume() int {
	return (1 + c.xRange.to - c.xRange.from) *
		(1 + c.yRange.to - c.yRange.from) *
		(1 + c.zRange.to - c.zRange.from)
}

func (c cube) invalid() bool {
	return c.xRange.from > c.xRange.to ||
		c.yRange.from > c.yRange.to ||
		c.zRange.from > c.zRange.to
}

func (c cube) intersects(c2 cube) bool {

	return !c.intersection(c2).invalid()
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
