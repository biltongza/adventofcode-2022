package day3

import (
	"bufio"
	"fmt"
	"os"
)

func Day3() {
	file, err := os.Open("day3/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	priorities := make(map[rune]int)
	for i := 0; i <= 26; i++ {
		priorities[rune(int('a')+i)] = i + 1
	}
	for i := 0; i <= 26; i++ {
		priorities[rune(int('A')+i)] = i + 27
	}

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		lineLength := len(line)
		firstCompartment := line[0 : lineLength/2]
		secondCompartment := line[lineLength/2 : lineLength]
		var item rune
		for _, itemFirstCompartment := range firstCompartment {
			found := false
			for _, itemSecondCompartment := range secondCompartment {
				if itemFirstCompartment == itemSecondCompartment {
					found = true
					item = itemFirstCompartment
					break
				}
			}
			if found {
				break
			}
		}

		total += priorities[item]
	}
	fmt.Fprintln(os.Stdout, "total", total)
}
