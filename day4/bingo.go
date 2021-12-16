package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

type position struct {
	value  int
	called bool
}

type board struct {
	bingo     bool
	score     int
	positions [5][5]position
}

func newBoard(vals ...int) board {
	b := board{}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			b.positions[i][j].value = vals[(5*j)+i]
		}
	}

	return b
}

func (b board) isBingoRow(row int) bool {
	for i := 0; i < 5; i++ {
		if !b.positions[i][row].called {
			return false
		}
	}

	return true
}

func (b board) isBingoCol(col int) bool {
	for i := 0; i < 5; i++ {
		if !b.positions[col][i].called {
			return false
		}
	}

	return true
}

func (b board) isBingo(col, row int) bool {
	return b.isBingoCol(col) || b.isBingoRow(row)
}

func (b board) computeScore(lastCall int) int {

	sumUncalled := 0

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b.positions[i][j].called {
				sumUncalled += b.positions[i][j].value
			}

		}
	}

	return lastCall * sumUncalled
}

func (b board) setCalled(value int) board {

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.positions[i][j].value == value {
				b.positions[i][j].called = true
				if b.isBingo(i, j) {
					b.bingo = true
					b.score = b.computeScore(value)
				}
			}
		}
	}

	return b
}

func run(args []string, stdout io.Writer) error {

	for c, value := range moves {
		for i := 0; i < len(boards); i++ {
			boards[i] = boards[i].setCalled(value)
			if boards[i].bingo {
				log.Printf("bingo! move %d, board %d, score %d", c+1, i, boards[i].score)
				return nil
			}
		}
	}

	return nil
}
