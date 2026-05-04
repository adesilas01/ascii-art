package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . [string] [bannerfile] | cat -e")
		return
	}

	if os.Args[1] == "" {
		return
	}

	words := strings.Split(os.Args[1], "\\n")

	font := "standard.txt"
	if len(os.Args) == 3 {
		font = os.Args[2]
	}

	file, err := os.ReadFile(font)
	if err != nil {
		fmt.Println("unable to read file!")
		return
	}

	content := strings.Split(string(file), "\n")
	if font == "thinkertoy.txt" {
		content = strings.Split(string(file), "\r\n")
	}

	for i, word := range words {
		if words[len(words)-1] == "" && i == len(words)-1 && words[i-1] == "" {
			break
		}

		if word == "" {
			fmt.Println()
			continue
		}

		for i := 0; i < 8; i++ {
			for _, char := range word {
				fmt.Printf("%s", content[FetchIndex(char)+i])
			}
			fmt.Println()
		}
	}
	// Reverse()
}

func FetchIndex(char rune) int {
	return 9*(int(char)-32) + 1
}
