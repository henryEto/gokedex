package main

import (
	"bufio"
	"fmt"
	"os"
)

const prompt = "Pokedex > "

func main() {
	dexScanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(prompt)
		if !dexScanner.Scan() {
			break
		}
		command := cleanInput(dexScanner.Text())
		fmt.Println("Your command was:", command[0])
	}
}
