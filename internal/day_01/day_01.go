package day01

import (
	"aoc/aoc_2025/pkg/utils"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func Solve() {
	header := strings.Repeat("=", 80)
	fmt.Println(header)
	fmt.Println("Running Aoc day 01")
	fmt.Println(header)
	fmt.Println("Solving part 1...")
	solvePart1()
	fmt.Println()
	fmt.Println("Solving part 2...")
	solvePart2()
}
func solvePart1() {
	position := 50
	pPosition := &position
	password := 0

	rotations := utils.ReadInputFile("01")
	for _, rot := range rotations {
		val := parseRotation(rot)
		udpatePosition(val, pPosition)
		if position == 0 {
			password += 1
		}
	}

	fmt.Printf("The password is: %d\n", password)

}

func solvePart2() {
	position := 50
	password := 0
	rotations := utils.ReadInputFile("01")

	for _, rot := range rotations {
		val := parseRotation(rot)
		nextPos := position + val
		zeroCross := 0
		if nextPos > position {
			end := math.Floor(float64(nextPos) / 100.0)
			start := math.Floor(float64(position) / 100.0)
			zeroCross = int(end - start)
		} else {
			start := math.Floor(float64(position-1) / 100.0)
			end := math.Floor(float64(nextPos-1) / 100.0)
			zeroCross = int(start - end)

		}
		password += zeroCross
		position = nextPos
	}
	fmt.Printf("The password is: %d\n", password)
}

func parseRotation(rotation string) int {
	value, err := strconv.Atoi(rotation[1:])
	if err != nil {
		log.Fatal(err)
	}

	if strings.ContainsAny(rotation, "L") {
		value = value * -1
	}
	return value
}

func udpatePosition(modifier int, pos *int) {
	newPos := (*pos + modifier) % 100
	*pos = newPos
	if newPos < 0 {
		*pos = newPos + 100
	}
}
