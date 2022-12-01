package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Day1() {
	file, err := os.Open("day1/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	//maxIndex := 0
	maxElf := 0
	currentElf := 0
	currentIndex := 1
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			num, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			currentElf = currentElf + num
		} else {
			if currentElf > maxElf {
				//maxIndex = currentIndex
				maxElf = currentElf
			}
			currentElf = 0
			currentIndex++
		}
	}

	fmt.Fprintln(os.Stdout, "Highest calories: ", maxElf)

}
