package day8

import (
	"bufio"
	"fmt"
	"os"
)

func Day8() {
	file, err := os.Open("day8/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	treeMap := make([][]int, 0)
	visibleFromOutside := make([][]bool, 0)
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		treeMap = append(treeMap, make([]int, len(line)))
		visibleFromOutside = append(visibleFromOutside, make([]bool, len(line)))
		for col, char := range line {
			treeMap[row][col] = int(char) - 48
		}

		row++
	}
	total := 0
	heightOfMap := len(treeMap)
	widthOfMap := len(treeMap[0])

	for indexTopDown := 0; indexTopDown < heightOfMap; indexTopDown++ {

		highestLeftRight := treeMap[indexTopDown][0]
		highestRightLeft := treeMap[indexTopDown][widthOfMap-1]
		visibleFromOutside[indexTopDown][0] = true
		visibleFromOutside[indexTopDown][widthOfMap-1] = true

		for indexLeftRight := 0; indexLeftRight < widthOfMap; indexLeftRight++ {
			indexRightLeft := widthOfMap - 1 - indexLeftRight

			// left to right
			if treeMap[indexTopDown][indexLeftRight] > highestLeftRight {
				visibleFromOutside[indexTopDown][indexLeftRight] = true
				highestLeftRight = treeMap[indexTopDown][indexLeftRight]
			}

			// right to left
			if treeMap[indexTopDown][indexRightLeft] > highestRightLeft {
				visibleFromOutside[indexTopDown][indexRightLeft] = true
				highestRightLeft = treeMap[indexTopDown][indexRightLeft]
			}
		}
	}

	for indexLeftRight := 0; indexLeftRight < widthOfMap; indexLeftRight++ {
		highestTopDown := treeMap[0][indexLeftRight]
		highestBottomUp := treeMap[heightOfMap-1][indexLeftRight]
		visibleFromOutside[0][indexLeftRight] = true
		visibleFromOutside[heightOfMap-1][indexLeftRight] = true
		for indexTopDown := 0; indexTopDown < heightOfMap; indexTopDown++ {
			indexBottomUp := heightOfMap - 1 - indexTopDown

			// top down
			if treeMap[indexTopDown][indexLeftRight] > highestTopDown {
				visibleFromOutside[indexTopDown][indexLeftRight] = true
				highestTopDown = treeMap[indexTopDown][indexLeftRight]
			}

			if treeMap[indexBottomUp][indexLeftRight] > highestBottomUp {
				visibleFromOutside[indexBottomUp][indexLeftRight] = true
				highestBottomUp = treeMap[indexBottomUp][indexLeftRight]
			}
		}
	}

	for _, row := range visibleFromOutside {
		for _, cell := range row {
			if cell {
				total++
			}
		}
	}

	fmt.Fprintf(os.Stdout, "total number of trees visible from outside: %d\n", total)

}
