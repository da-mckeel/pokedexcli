package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		text := scanner.Text()
		firstWord := cleanInput(text)[0]
		fmt.Printf("Your command was: %s\n", firstWord)
	}
}
