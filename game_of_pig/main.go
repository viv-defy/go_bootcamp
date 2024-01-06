package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func findMin(a, b int) int {
	if a < b {
		return a
	}

	return b
}

type player struct {
	id       int
	strategy int
	score    int
}

func (p *player) turn() {
	var turnScore int
	toScore := findMin(p.strategy, 100-p.score)
	for turnScore < toScore {
		roll := rand.Intn(5) + 1
		if roll == 1 {
			return
		} else {
			turnScore += roll
		}
	}
	p.score += turnScore
}

func (p *player) reset() {
	p.score = 0
}

func game(p1, p2 player) player {
	t := 1
	for {
		if t == 1 {
			p1.turn()
			if p1.score >= 100 {
				return p1
			}
		} else {
			p2.turn()
			if p2.score >= 100 {
				return p2
			}
		}
		t *= -1
	}
}

func simulateGames(p1Hold, p2Hold int) {
	p1 := player{id: 1, strategy: p1Hold}
	p2 := player{id: 2, strategy: p2Hold}
	var scoreBoard [2]int
	for i := 0; i < 10; i++ {
		winner := game(p1, p2)
		scoreBoard[winner.id-1]++
		p1.reset()
		p2.reset()
	}
	fmt.Printf("Holding at  %v vs Holding at  %v: wins: %v/10 (%.1f%%), losses: %v/10 (%.1f%%)\n",
		p1.strategy, p2.strategy,
		scoreBoard[0], float64(scoreBoard[0])/10.0*100.0,
		scoreBoard[1], float64(scoreBoard[1])/10.0*100.0)
}

func fixStartegies(p1Start, p1End, p2Start, p2End int) {
	for i := p1Start; i <= p1End; i++ {
		for j := p2Start; j <= p2End; j++ {
			if i == j {
				continue
			}
			simulateGames(i, j)
		}
	}
}

func story1() {
	p1Hold, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid argument 1")
		return
	}

	p2Hold, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid argument 2")
		return
	}

	fixStartegies(p1Hold, p1Hold, p2Hold, p2Hold)
}

func story2() {
	p1Hold, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid argument 1")
		return
	}

	p2Range := strings.Split(os.Args[2], "-")
	p2Start, err := strconv.Atoi(p2Range[0])
	if err != nil {
		fmt.Println("Invalid argument 2")
		return
	}
	p2End, err := strconv.Atoi(p2Range[1])
	if err != nil {
		fmt.Println("Invalid argument 2")
		return
	}
	if p2Start > p2End {
		fmt.Println("Invalid argument 2")
		return
	}

	fixStartegies(p1Hold, p1Hold, p2Start, p2End)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./<main> <player1 strategy> <player2 strategy>")
		return
	}

	story1()
	story2()
}
