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
	if err := part1(os.Args, os.Stdout); err != nil {
		log.Fatalf("program failed: %s\n", err)
	}
	if err := part2(os.Args, os.Stdout); err != nil {
		log.Fatalf("program failed: %s\n", err)
	}
	fmt.Printf("fini\n")
}

type pairCounts struct {
	first  rune
	pairs  map[pair]int
	counts map[rune]int
}

func newPairCounts(first rune) pairCounts {
	return pairCounts{
		first,
		map[pair]int{},
		map[rune]int{first: 1},
	}
}

func loadPairs(template string) pairCounts {
	pc := newPairCounts(rune(template[0]))

	for i := 0; i < len(template)-1; i++ {
		left, right := rune(template[i]), rune(template[i+1])
		pc.addPairs(pair{left, right}, 1)
	}

	return pc
}

func (pc *pairCounts) addPairs(p pair, count int) {
	pc.pairs[p] += count
	pc.counts[p.right] += count
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

	pc := loadPairs(start)

	for i := 0; i < 40; i++ {
		pc = pc.incrPairCounts()
		log.Printf("counts %s", printable(pc.counts))
	}

	most, least := mostAndLeast(pc.counts)
	log.Printf("most %d, least %d, diff %d", most, least, most-least)

	return nil
}

func (pc pairCounts) incrPairCounts() pairCounts {

	newPC := newPairCounts(pc.first)

	for cur, cnt := range pc.pairs {
		ins, ok := rules[cur]
		if !ok {
			newPC.addPairs(cur, cnt)
			continue
		}

		newLeftPair := pair{cur.left, ins}
		newRightPair := pair{ins, cur.right}

		newPC.addPairs(newLeftPair, cnt)
		newPC.addPairs(newRightPair, cnt)

	}

	return newPC
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
