package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
)

func main() {
	// if err := part1(os.Args, os.Stdout); err != nil {
	// 	log.Fatalf("program failed: %s\n", err)
	// }
	if err := part2(os.Args, os.Stdout); err != nil {
		log.Fatalf("program failed: %s\n", err)
	}
	fmt.Printf("fini\n")
}

type poly struct {
	element rune
	next    *poly
}

func newPoly(template string) *poly {
	var r, c *poly

	for _, e := range template {
		if c == nil {
			c = new(poly)
		} else {
			c.next = new(poly)
			c = c.next
		}
		if r == nil {
			r = c
		}

		c.element = e
	}

	return r
}

func part1(args []string, stdout io.Writer) error {

	t := start

	for i := 0; i < 10; i++ {
		t = insert(t)
		log.Printf("counts %s", printable(counts(t)))
	}

	most, least := mostAndLeast(counts(t))
	log.Printf("most %d, least %d, diff %d", most, least, most-least)

	return nil
}

func part2(args []string, stdout io.Writer) error {

	t := newPoly(start)

	for i := 0; i < 40; i++ {
		t.insert()
		log.Printf("counts %s", printable(t.counts()))
	}

	most, least := mostAndLeast(t.counts())
	log.Printf("most %d, least %d, diff %d", most, least, most-least)

	return nil
}

func printable(counts map[rune]int) string {

	keys := []rune{}
	for k := range counts {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	s := ""

	for _, k := range keys {
		s += fmt.Sprintf("%c: %6d, ", k, counts[k])
	}
	return s
}

func mostAndLeast(counts map[rune]int) (int, int) {
	most, least := 0, math.MaxInt

	for _, v := range counts {
		if v > most {
			most = v
		}
		if v < least {
			least = v
		}
	}

	return most, least
}

func counts(template string) map[rune]int {
	r := map[rune]int{}

	for _, each := range template {
		r[each]++
	}

	return r
}

func (p *poly) counts() map[rune]int {
	r := map[rune]int{}

	for c := p; c != nil; c = c.next {
		r[c.element]++
	}

	return r
}

func insert(template string) string {
	r := ""

	var last rune
	for _, each := range template {

		ins, ok := rules[pair{last, each}]
		if ok {
			r += string(ins)
		}
		r += string(each)
		last = each
	}

	return r
}

func (p *poly) insert() {

	for c := p; c.next != nil; c = c.next {
		k := pair{c.element, c.next.element}
		ins, ok := rules[k]
		if ok {
			c.next = &poly{ins, c.next}
			c = c.next
		}
	}
}
