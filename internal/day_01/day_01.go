package day01

import (
	"aoc/aoc_2025/pkg/utils"
	"fmt"
	"log"
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
