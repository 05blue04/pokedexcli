package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CleanInput(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		input := CleanInput(scanner.Text())

		if len(input) == 0 {
			continue // skip empty lines
		}

		// fmt.Print("Pokedex > ")
		fmt.Println("Your command was:", input[0])

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
}
