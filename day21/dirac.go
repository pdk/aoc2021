package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		log.Fatalf("program failed: %s\n", err)
	}
	fmt.Printf("fini\n")
}

func run(args []string, stdout io.Writer) error {

	d := die{}

	rollCount := 0
	playerN := 0

	for rollCount < 1000 {

		r1, r2, r3 := d.roll(), d.roll(), d.roll()
		rollVal := r1 + r2 + r3
		rollCount += 3

		players[playerN].move(rollVal)
		log.Printf("roll count %d player %d at %d with score %d",
			rollCount, playerN, players[playerN].pos, players[playerN].score)

		if players[playerN].score >= 1000 {
			break
		}

		playerN = otherPlayer(playerN)
	}

	loser := otherPlayer(playerN)
	log.Printf("after %d rolls, the loser has %d points, puzzle input is %d",
		rollCount, players[loser].score, rollCount*players[loser].score)

	return nil
}

func otherPlayer(n int) int {
	if n == 1 {
		return 0
	}
	return 1
}

type player struct {
	pos   int
	score int
}

var players = []player{
	{pos: 10},
	{pos: 2},
	// {pos: 4},
	// {pos: 8},
}

func (p *player) move(n int) {

	p.pos += n
	if p.pos > 10 {
		p.pos = p.pos % 10
	}
	if p.pos == 0 {
		p.pos = 10
	}

	p.score += p.pos
}

type die struct {
	lastRoll int
}

func (d *die) roll() int {
	d.lastRoll++
	if d.lastRoll > 100 {
		d.lastRoll = 1
	}
	return d.lastRoll
}
