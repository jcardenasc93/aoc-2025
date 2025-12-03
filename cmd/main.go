package main

import (
	"aoc/aoc_2025/internal/day_01"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: Missing day")
		return
	}

	dayStr := os.Args[1]

	switch dayStr {
	case "1", "01":
		day01.Solve()
	default:
		fmt.Println("Not a valid day")
	}
}
