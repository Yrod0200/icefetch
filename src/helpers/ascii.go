package helpers

import (
	"bufio"
	"fmt"
	"os"
)

func GetLogoSlice(path string) []string {

	var lines []string

	file, err := os.Open(path)

	if err != nil {
		fmt.Printf("Error opening ascii file: %s", err)
		os.Exit(-1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		lines = append(lines, scanner.Text())

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
		}

	}

	return lines
}
