package day4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day4() {
	file, err := os.Open("day4/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, ",")
		first := parts[0]
		second := parts[1]
		firstStart, firstEnd := getRange(first)
		secondStart, secondEnd := getRange(second)

		containedInFirst := true
		for i := firstStart; i <= firstEnd; i++ {
			containedInFirst = containedInFirst && i >= secondStart && i <= secondEnd
			if !containedInFirst {
				break
			}
		}

		containedInSecond := true
		for i := secondStart; i <= secondEnd; i++ {
			containedInSecond = containedInSecond && i >= firstStart && i <= firstEnd
			if !containedInSecond {
				break
			}
		}

		if containedInFirst || containedInSecond {
			count++
		}

	}
	fmt.Fprintln(os.Stdout, "num of overlaps: ", count)
}

func getRange(input string) (start int, end int) {
	parts := strings.Split(input, "-")
	start, _ = strconv.Atoi(parts[0])
	end, _ = strconv.Atoi(parts[1])
	return
}
