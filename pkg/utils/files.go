package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func ReadInputFile(day string) []string {
	fileName := fmt.Sprintf("day_%s.txt", day)
	filePath := filepath.Join("inputs", fileName)
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}
