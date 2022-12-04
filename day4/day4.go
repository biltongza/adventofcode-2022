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

		overlap := false
		for i := firstStart; i <= firstEnd; i++ {
			for j := secondStart; j <= secondEnd; j++ {
				if i == j {
					overlap = true
					break
				}
			}
			if overlap {
				break
			}
		}

		if overlap {
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
