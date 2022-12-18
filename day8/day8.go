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
	scenicScores := make([][]int, 0)
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		treeMap = append(treeMap, make([]int, len(line)))
		scenicScores = append(scenicScores, make([]int, len(line)))
		for col, char := range line {
			treeMap[row][col] = int(char) - 48
		}

		row++
	}

	heightOfMap := len(treeMap)
	widthOfMap := len(treeMap[0])

	for indexTopDown := 0; indexTopDown < heightOfMap; indexTopDown++ {
		for indexLeftRight := 0; indexLeftRight < widthOfMap; indexLeftRight++ {
			scenicScores[indexTopDown][indexLeftRight] = calculateScore(treeMap, indexLeftRight, indexTopDown)
		}
	}

	highestScenicScore := 0
	for _, row := range scenicScores {
		for _, cell := range row {
			if cell > highestScenicScore {
				highestScenicScore = cell
			}
		}
	}

	fmt.Fprintf(os.Stdout, "highestScenicScore: %d\n", highestScenicScore)

}

func calculateScore(treeMap [][]int, x int, y int) int {
	heightOfMap := len(treeMap)
	widthOfMap := len(treeMap[0])

	scoreUp := 0
	scoreDown := 0
	scoreLeft := 0
	scoreRight := 0

	// look up
	for yCopy := y - 1; yCopy >= 0; yCopy-- {
		if treeMap[yCopy][x] < treeMap[y][x] {
			scoreUp++
		} else if treeMap[yCopy][x] >= treeMap[y][x] {
			scoreUp++
			break
		}
	}
	// look down
	for yCopy := y + 1; yCopy < heightOfMap; yCopy++ {
		if treeMap[yCopy][x] < treeMap[y][x] {
			scoreDown++
		} else if treeMap[yCopy][x] >= treeMap[y][x] {
			scoreDown++
			break
		}
	}

	// look left
	for xCopy := x - 1; xCopy >= 0; xCopy-- {
		if treeMap[y][xCopy] < treeMap[y][x] {
			scoreLeft++
		} else if treeMap[y][xCopy] >= treeMap[y][x] {
			scoreLeft++
			break
		}
	}

	// look right
	for xCopy := x + 1; xCopy < widthOfMap; xCopy++ {
		if treeMap[y][xCopy] < treeMap[y][x] {
			scoreRight++
		} else if treeMap[y][xCopy] >= treeMap[y][x] {
			scoreRight++
			break
		}
	}
	return scoreUp * scoreDown * scoreLeft * scoreRight
}
