package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		log.Fatalf("program failed: %s\n", err)
	}

	// fmt.Printf("fini\n")
}

func run(args []string, stdout io.Writer) error {

	// for _, n := range []int{10, 11, 12} {
	// 	fmt.Printf("%d => %s\n", n, split(n).String())
	// }

	// for _, e := range data {
	// 	fmt.Printf("%s\n", e.String())
	// }

	// for _, t := range []struct{ before, after *node }{

	// 	{before: np(np(np(np(np(9, 8), 1), 2), 3), 4), after: np(np(np(np(0, 9), 2), 3), 4)},
	// 	{before: np(7, np(6, np(5, np(4, np(3, 2))))), after: np(7, np(6, np(5, np(7, 0))))},
	// 	{before: np(np(6, np(5, np(4, np(3, 2)))), 1), after: np(np(6, np(5, np(7, 0))), 3)},
	// 	{before: np(np(3, np(2, np(1, np(7, 3)))), np(6, np(5, np(4, np(3, 2))))), after: np(np(3, np(2, np(8, 0))), np(9, np(5, np(4, np(3, 2)))))},
	// 	{before: np(np(3, np(2, np(8, 0))), np(9, np(5, np(4, np(3, 2))))), after: np(np(3, np(2, np(8, 0))), np(9, np(5, np(7, 0))))},
	// } {

	// 	log.Printf("input  %s\n", t.before)
	// 	r, _ := explode(t.before, 0, nil, nil)
	// 	log.Printf("got    %s\n", r)
	// 	log.Printf("expect %s\n", t.after)
	// 	log.Printf("---------------\n")
	// }

	// [[[[4,3],4],4],[7,[[8,4],9]]] + [1,1]
	// add(
	// 	np(np(np(np(4, 3), 4), 4), np(7, np(np(8, 4), 9))),
	// 	np(1, 1))

	// s := data[0]
	// // log.Printf("s = %s", s.String())
	// for i := 1; i < len(data); i++ {
	// 	s = add(s, data[i])
	// 	// log.Printf("s = %s", s.String())
	// }

	// log.Printf("final mag is %d", mag(s))

	maxMag := 0

	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {

			a := copy(data[i])
			b := copy(data[j])

			newMag := mag(add(a, b))
			if newMag > maxMag {
				maxMag = newMag
				log.Printf("new max mag %d", maxMag)
			}

			a = copy(data[i])
			b = copy(data[j])

			newMag = mag(add(b, a))
			if newMag > maxMag {
				maxMag = newMag
				log.Printf("new max mag %d", maxMag)
			}

		}
	}

	log.Printf("max mag is %d", maxMag)

	return nil
}

func copy(n *node) *node {
	if n.isVal() {
		return &node{val: n.val}
	}

	return &node{
		left:  copy(n.left),
		right: copy(n.right),
	}
}

func mag(n *node) int {

	if n == nil {
		return 0
	}

	if n.isVal() {
		return n.val
	}

	return 3*mag(n.left) + 2*mag(n.right)
}

func add(a, b *node) *node {

	// log.Printf("%s + %s", a, b)

	ans := reduce(np(a, b))

	// log.Printf("= %s", ans)

	return ans
}

func reduce(n *node) *node {

	// log.Printf("reduce %s", n.String())

	for {
		var changed bool

		n, changed = explode(n, 0, nil, nil)
		if changed {
			// log.Printf("exploded %s", n.String())
			continue
		}

		n, changed = splitLargerThan10(n)
		if changed {
			// log.Printf("split %s", n.String())
			continue
		}

		break
	}

	// log.Printf("complete %s", n.String())
	return n
}

func splitLargerThan10(n *node) (*node, bool) {
	var changed bool

	if n == nil {
		return nil, false
	}

	if n.isVal() && n.val >= 10 {
		return split(n), true
	}

	n.left, changed = splitLargerThan10(n.left)
	if changed {
		return n, true
	}

	n.right, changed = splitLargerThan10(n.right)
	return n, changed
}

func findLeftmostLeaf(n *node) *node {
	// log.Printf("findLeftmostLeaf %s", n.String())
	if n == nil {
		return nil
	}
	if n.isVal() {
		return n
	}
	return findLeftmostLeaf(n.left)
}

func findRightmostLeaf(n *node) *node {
	// log.Printf("findRightmostLeaf %s", n.String())
	if n == nil {
		return nil
	}
	if n.isVal() {
		return n
	}
	return findRightmostLeaf(n.right)
}

func explode(n *node, depth int, parentLeft, parentRight *node) (*node, bool) {
	var changed bool

	if n == nil {
		return nil, false
	}

	// dig left
	n.left, changed = explode(n.left, depth+1, parentLeft, n)
	if changed {
		return n, true
	}

	if depth >= 4 && n.hasTwoVals() {
		// log.Printf("depth=%2d exploding %s", depth, n.String())

		if parentLeft != nil {
			leftLeaf := findRightmostLeaf(parentLeft.left)
			if leftLeaf != nil {
				leftLeaf.val += n.left.val
			}
		}
		if parentRight != nil {
			rightLeaf := findLeftmostLeaf(parentRight.right)
			if rightLeaf != nil {
				rightLeaf.val += n.right.val
			}
		}
		return &node{
			val: 0,
		}, true
	}

	// dig right
	n.right, changed = explode(n.right, depth+1, n, parentRight)
	return n, changed
}

func split(n *node) *node {
	return &node{
		left:  &node{val: n.val / 2},
		right: &node{val: n.val/2 + n.val%2},
	}
}

func coalesce(vals ...*node) *node {
	for _, p := range vals {
		if p != nil {
			return p
		}
	}
	return nil
}

type node struct {
	left, right *node
	val         int
}

func (n *node) isVal() bool {
	if n == nil {
		return false
	}
	return n.left == nil && n.right == nil
}

func (n *node) String() string {
	if n == nil {
		return ""
	}
	if n.isVal() {
		return strconv.Itoa(n.val)
	}
	return fmt.Sprintf("[%s,%s]", n.left.String(), n.right.String())
}

func (n *node) hasTwoVals() bool {
	return n != nil && n.left.isVal() && n.right.isVal()
}

func np(a, b interface{}) *node {

	alit, alitok := a.(int)
	blit, blitok := b.(int)

	switch {
	case alitok && blitok:
		return &node{
			left:  &node{val: alit},
			right: &node{val: blit},
		}
	case alitok:
		return &node{
			left:  &node{val: alit},
			right: b.(*node),
		}
	case blitok:
		return &node{
			left:  a.(*node),
			right: &node{val: blit},
		}
	}

	return &node{
		left:  a.(*node),
		right: b.(*node),
	}
}
