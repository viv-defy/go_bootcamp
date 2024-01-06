package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./<main> <player1 hold> <player2 hold>")
		return
	}

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

	fmt.Printf("you have entered %v and %v\n", p1Hold, p2Hold)
}
