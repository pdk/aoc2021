package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		log.Fatalf("program failed: %s\n", err)
	}
	// fmt.Printf("fini\n")
}

func run(args []string, stdout io.Writer) error {

	root := initialUniverse

	player := 0
	for roll := 1; roll < 50; roll++ {

		nonTerminal := root.leaves()
		if len(nonTerminal) == 0 {
			break
		}
		log.Printf("roll %d non terminal %d", roll, len(nonTerminal))
		s := root.getStats()
		log.Printf("roll %d nodeCount %d", roll, s.nodeCount)

		highScore := 0
		for _, uni := range nonTerminal {
			hiScore := uni.play(player)
			if hiScore > highScore {
				highScore = hiScore
			}
		}
		// log.Printf("high score %d", highScore)

		player = otherPlayer(player)
	}

	s := root.getStats()

	log.Printf("player1Wins %d", s.player1Wins)
	log.Printf("player1Losses %d", s.player1Losses)
	log.Printf("player2Wins %d", s.player2Wins)
	log.Printf("player2Losses %d", s.player2Losses)
	log.Printf("nodeCount %d", s.nodeCount)
	log.Printf("universeCount %d", s.universeCount)

	return nil
}

type stats struct {
	player1Wins, player1Losses int
	player2Wins, player2Losses int
	nodeCount, universeCount   int
}

func (s stats) plus(more stats) stats {
	return stats{
		player1Wins:   s.player1Wins + more.player1Wins,
		player1Losses: s.player1Losses + more.player1Losses,
		player2Wins:   s.player2Wins + more.player2Wins,
		player2Losses: s.player2Losses + more.player2Losses,
		nodeCount:     s.nodeCount + more.nodeCount,
		universeCount: s.universeCount + more.universeCount,
	}
}

func (u *universe) getStats() stats {

	if u == nil {
		return stats{}
	}

	if u.players[0].score >= 21 {
		return stats{
			player1Wins:   u.universeCount,
			player2Losses: u.universeCount,
			nodeCount:     1,
			universeCount: u.universeCount,
		}
	}

	if u.players[1].score >= 21 {
		return stats{
			player2Wins:   u.universeCount,
			player1Losses: u.universeCount,
			nodeCount:     1,
			universeCount: u.universeCount,
		}
	}

	s := stats{
		nodeCount: 1,
	}

	for i := 0; i < 7; i++ {
		s = s.plus(u.forks[i].getStats())
	}

	return s
}

func (u *universe) leaves() []*universe {

	if u == nil || u.isTerminal() {
		return []*universe{}
	}

	if u.forks[0] == nil {
		// leaf node
		return []*universe{u}
	}

	r := []*universe{}
	for i := 0; i < 7; i++ {
		n := u.forks[i].leaves()
		r = append(r, n...)
	}

	return r
}

var initialUniverse = universe{
	universeCount: 1,
	players: [2]player{
		// example:
		// {position: 4},
		// {position: 8},
		// actual:
		{position: 10},
		{position: 2},
	},
}

type player struct {
	position int
	score    int
}

type universe struct {
	rollVal       int
	rollCount     int
	players       [2]player
	universeCount int
	forks         [7]*universe // each roll spawns 7 universe variations, 3-9
}

func (u *universe) play(player int) int {

	highScore := 0
	if u.isTerminal() {
		log.Fatal("try to play on terminal node")
	}

	for r := 3; r <= 9; r++ {
		i := r - 3
		c := rollCount(r)

		u.forks[i] = &universe{
			rollVal:       r,
			rollCount:     c,
			players:       u.players, // copy array
			universeCount: u.universeCount * c,
		}

		p := &u.forks[i].players[player]
		p.position = addPos(p.position, r)
		p.score += p.position
		if p.score > highScore {
			highScore = p.score
		}
	}

	return highScore
}

func (u *universe) isTerminal() bool {
	return u.players[0].score >= 21 || u.players[1].score >= 21
}

func rollCount(n int) int {
	switch n {
	case 3:
		return 1
	case 4:
		return 3
	case 5:
		return 6
	case 6:
		return 7
	case 7:
		return 6
	case 8:
		return 3
	case 9:
		return 1
	}

	log.Fatalf("illegal rollCount request: %d", n)

	return 0
}

func otherPlayer(n int) int {
	if n == 1 {
		return 0
	}
	return 1
}

func addPos(start, dist int) int {
	newPos := start + dist
	if newPos > 10 {
		return newPos - 10
	}
	return newPos
}
