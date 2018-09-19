package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main() {
	words := map[string]int{}

	var input string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}

	wordList := strings.Split(input, " ")

	for _, word := range wordList {
		if _, ok := words[word]; !ok {
			words[word] = 1
		} else {
			fmt.Println("no")
			return
		}
	}

	fmt.Println("yes")
}
