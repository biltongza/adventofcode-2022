package day5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day5() {
	file, err := os.Open("day5/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	stacks := make(map[int]string)

	scanMode := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if scanMode == 1 {
				break
			}
			scanMode = 1
			continue
		}

		if scanMode == 0 {
			index := 1
			for i := 0; i < 35; i += 4 {
				part := line[i:min(i+4, len(line))]
				part = strings.ReplaceAll(strings.ReplaceAll(part, "[", ""), "]", "")
				part = strings.TrimSpace(part)
				// this is gross
				_, err = strconv.Atoi(part)
				if err != nil {
					stacks[index] = part + stacks[index]
				}
				index++
			}
		} else {
			parts := strings.Split(line, " ")
			countToMove, _ := strconv.Atoi(parts[1])
			src, _ := strconv.Atoi(parts[3])
			dst, _ := strconv.Atoi(parts[5])
			srcLen := len(stacks[src])
			crates := stacks[src][srcLen-countToMove:]
			crates = Reverse(crates)
			stacks[src] = stacks[src][:srcLen-countToMove]
			stacks[dst] = stacks[dst] + crates
		}
	}
	var lastCrates string
	for i := 1; i <= 9; i++ {
		stack := stacks[i]
		lastCrates = lastCrates + stack[len(stack)-1:]
	}
	fmt.Fprintln(os.Stdout, lastCrates)
}

func min(a int, b int) (result int) {
	if a > b {
		result = b
	} else if a < b {
		result = a
	} else {
		result = a
	}
	return
}

// https://stackoverflow.com/a/4965535/1492861
func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
