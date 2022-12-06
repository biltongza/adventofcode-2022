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
		length := len(line)
		for i := 0; i <= length; i++ {
			buf := line[i : i+4]
			skip := false
			for index, char := range buf {
				for j := 0; j <= 3; j++ {
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
				fmt.Fprintf(os.Stdout, "found start of packet at index %d\n", i+4)
				break
			}
		}

	}
}
