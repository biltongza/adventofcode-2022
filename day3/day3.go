package day3

import (
	"fmt"
	"os"
	"strings"
)

func Day3() {
	file, err := os.ReadFile("day3/input.txt")
	if err != nil {
		panic(err)
	}
	fileAsString := string(file)
	lines := strings.SplitAfter(fileAsString, "\n")

	priorities := make(map[rune]int)
	for i := 0; i <= 26; i++ {
		priorities[rune(int('a')+i)] = i + 1
	}
	for i := 0; i <= 26; i++ {
		priorities[rune(int('A')+i)] = i + 27
	}

	total := 0
	for lineIndex := 0; lineIndex <= len(lines); {
		firstElf := lines[lineIndex]
		if firstElf == "" {
			break
		}
		secondElf := lines[lineIndex+1]
		thirdElf := lines[lineIndex+2]

		var item rune
		for _, itemFirstElf := range firstElf {
			found := false
			for _, itemSecondElf := range secondElf {
				for _, itemThirdElf := range thirdElf {
					if itemFirstElf == itemSecondElf && itemFirstElf == itemThirdElf {
						found = true
						item = itemFirstElf
						break
					}
				}
				if found {
					break
				}
			}
			if found {
				break
			}
		}

		total += priorities[item]
		lineIndex += 3
	}
	fmt.Fprintln(os.Stdout, "total", total)
}
