package day6

import (
	"bufio"
	"fmt"
	"os"
)

func Day6() {
	file, err := os.Open("day6/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		startOfMessage := detectUniqueSequence(line, 14)
		fmt.Fprintf(os.Stdout, "detected start of message at %d\n", startOfMessage)
	}
}

func detectUniqueSequence(str string, sequenceLength int) int {
	length := len(str)
	num := -1
	for i := 0; i <= length; i++ {
		buf := str[i : i+sequenceLength]
		skip := false
		for index, char := range buf {
			for j := 0; j <= sequenceLength-1; j++ {
				if j <= index {
					continue
				}
				otherChar := rune(buf[j])
				if char == otherChar {
					skip = true
					break
				}
			}
			if skip {
				break
			}
		}
		if skip {
			continue
		} else {
			num = i + sequenceLength
			break
		}
	}
	return num
}
