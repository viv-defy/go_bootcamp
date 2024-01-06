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

func simulateGames(p1Hold, p2Hold int) int {
	p1 := player{id: 1, strategy: p1Hold}
	p2 := player{id: 2, strategy: p2Hold}
	var wins int
	for i := 0; i < 10; i++ {
		if winner := game(p1, p2); winner.id == 1 {
			wins += 1
		}
		p1.reset()
		p2.reset()
	}
	return wins
}

func fixStartegies(p1Start int, p1End int, p2Start int, p2End int, printPerStrategy bool) {
	var wins, winsAti int
	for i := p1Start; i <= p1End; i++ {
		for j := p2Start; j <= p2End; j++ {
			if i == j {
				continue
			}
			wins = simulateGames(i, j)
			if printPerStrategy {
				fmt.Printf("Holding at  %v vs Holding at  %v: wins: %v/10 (%.1f%%), losses: %v/10 (%.1f%%)\n",
					i, j,
					wins, float64(wins)/10.0*100.0,
					10-wins, float64(wins)/10.0*100.0)
			} else {
				winsAti += wins
			}
		}
		if !printPerStrategy {
			fmt.Printf("Result: Wins, losses staying at k = %v: %v/990 (%.1f%%), %v/990 (%.1f%%)\n",
				i, winsAti, float64(winsAti)/990.0*100.0, 990-winsAti, float64(990-winsAti)/990.0*100.0)
		}
		winsAti = 0
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

	fixStartegies(p1Hold, p1Hold, p2Hold, p2Hold, true)
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

	fixStartegies(p1Hold, p1Hold, p2Start, p2End, true)
}

func story3() {
	p1Range := strings.Split(os.Args[1], "-")
	p1Start, err := strconv.Atoi(p1Range[0])
	if err != nil {
		fmt.Println("Invalid argument 1")
		return
	}
	p1End, err := strconv.Atoi(p1Range[1])
	if err != nil {
		fmt.Println("Invalid argument 1")
		return
	}
	if p1Start > p1End {
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

	fixStartegies(p1Start, p1End, p2Start, p2End, false)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./<main> <player1 strategy> <player2 strategy>")
		return
	}

	story1()
	story2()
	story3()
}
