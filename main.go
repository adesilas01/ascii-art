package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	// 1. Check for empty arguments
	if len(args) == 0 {
		printUsage()
		return
	}

	alignValue := "left" // Default alignment
	var remainingArgs []string

	// 2. Parse Arguments
	for _, arg := range args {
		if strings.HasPrefix(arg, "--") {
			// Check if flag starts with the correct prefix
			if !strings.HasPrefix(arg, "--align=") {
				printUsage()
				return
			}

			// Extract value after "--align="
			splitFlag := strings.SplitN(arg, "=", 2)
			if len(splitFlag) < 2 || splitFlag[1] == "" {
				printUsage()
				return
			}
			alignValue = splitFlag[1]
		} else {
			remainingArgs = append(remainingArgs, arg)
		}
	}

	// 3. Validate alignment options
	validAligns := map[string]bool{"left": true, "right": true, "center": true, "justify": true}
	if !validAligns[alignValue] {
		printUsage()
		return
	}

	// 4. Assign Input and Font
	if len(remainingArgs) < 1 {
		printUsage()
		return
	}

	userInput := remainingArgs[0]
	font := "standard.txt"

	if len(remainingArgs) >= 2 {
		font = remainingArgs[1]
		if !strings.HasSuffix(font, ".txt") {
			font += ".txt"
		}
	}

	// 5. File Processing
	bannerdata, err := os.ReadFile(font)
	if err != nil {
		// If the file is missing, we exit.
		// Note: Usually, file errors are distinct from usage errors.
		fmt.Printf("Error: banner file '%s' not found\n", font)
		os.Exit(1)
	}

	standardizedContent := strings.ReplaceAll(string(bannerdata), "\r\n", "\n")
	content := strings.Split(standardizedContent, "\n")

	// Handle literal \n in string
	userline := strings.Split(userInput, "\\n")
	twidth := GetTerminalWidth()

	// 6. Output Logic
	// If it's not justify, or it's justify but only one word, use standard padding
	if alignValue != "justify" || !strings.Contains(userInput, " ") {
		for _, line := range userline {
			if line == "" {
				fmt.Println()
				continue
			}
			wwidth := GetArtWidth(line, content)
			pad := NumSpace(alignValue, wwidth, twidth)
			for i := 0; i < 8; i++ {
				fmt.Print(pad)
				for _, char := range line {
					fmt.Printf("%s", content[Calculator(char)+i])
				}
				fmt.Println()
			}
		}
	} else {
		// Justify Logic for multiple words
		for _, line := range userline {
			if line == "" {
				fmt.Println()
				continue
			}
			words, gapSize := multipleWords(line, content)
			for i := 0; i < 8; i++ {
				for j, word := range words {
					for _, char := range word {
						fmt.Printf("%s", content[Calculator(char)+i])
					}
					// Add spaces between words, but not after the last word
					if j < len(words)-1 {
						fmt.Print(strings.Repeat(" ", gapSize))
					}
				}
				fmt.Println()
			}
		}
	}
}

func printUsage() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
	fmt.Println()
	fmt.Println("Example: go run . --align=right something standard")
}
