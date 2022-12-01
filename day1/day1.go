package day1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func Day1() {
	file, err := os.Open("day1/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	elves := make([]int, 0)
	currentElf := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			num, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			currentElf = currentElf + num
		} else {
			elves = append(elves, currentElf)
			currentElf = 0
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(elves)))

	combined := elves[0] + elves[1] + elves[2]

	fmt.Fprintln(os.Stdout, "combined top 3: ", combined)

}
