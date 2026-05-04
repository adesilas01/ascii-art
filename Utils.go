package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func GetTerminalWidth() int {
	cmd := exec.Command("tput", "cols")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()

	if err != nil {
		fmt.Println("Error getting terminal width", err)
		os.Exit(2)
	}

	width := strings.TrimSpace(string(out))

	num, err := strconv.Atoi(width)

	if err != nil {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println("Example: go run . --align=right something standard")
		os.Exit(3)
	}

	return num
}

func Calculator(char rune) int {
	return 9*(int(char)-32) + 1
}

func GetArtWidth(word string, banner []string) int {

	var index, artwidth int
	for _, char := range word {
		index = int(Calculator(char))
		artwidth += len(banner[index])
	}
	return artwidth
}

func NumSpace(side string, wordwidth, Termwidth int) string {
	var padding string
	var spaces int
	if side == "left" || side == "right" || side == "center" || side == "justify" {

		if side == "left" {
			spaces = 0
			padding = strings.Repeat(" ", spaces)
		}

		if side == "right" {
			spaces = Termwidth - wordwidth
			padding = strings.Repeat(" ", spaces)

		}
		if side == "center" {
			spaces = (Termwidth - wordwidth) / 2
			padding = strings.Repeat(" ", spaces)
		}

	}
	return padding
}

func multipleWords(input string, banner []string) ([]string, int) {

	var inspace int
	var totalwidth int
	splitbyspace := strings.Split(input, " ")
	for _, line := range splitbyspace {
		totalwidth += GetArtWidth(line, banner)
	}
	inspace = len(splitbyspace) - 1
	totalGapSpace := GetTerminalWidth() - totalwidth
	gapSize := totalGapSpace / inspace
	return splitbyspace, gapSize
}
