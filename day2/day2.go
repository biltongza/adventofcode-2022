package day2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Day2() {
	scores := make(map[string]int)
	scores["A"] = 1
	scores["B"] = 2
	scores["C"] = 3
	scores["X"] = 1
	scores["Y"] = 2
	scores["Z"] = 3
	file, err := os.Open("day2/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	totalMyScore := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, " ")
		roundOpponentScore := 0
		roundMyScore := 0
		opponentMove := parts[0]
		myTargetMove := parts[1]

		opponentScore := scores[opponentMove]

		opponentResult, myResult, myMove := determineScores(opponentScore, myTargetMove)

		roundOpponentScore += opponentResult
		roundMyScore += myResult + myMove
		totalMyScore += roundMyScore

	}
	fmt.Fprintln(os.Stdout, "my total score: ", totalMyScore)
}

func determineScores(opponentMove int, myTargetMove string) (opponentScore int, myScore int, myMove int) {
	if myTargetMove == "X" {
		opponentScore = 6
		myScore = 0
		if opponentMove == 1 {
			myMove = 3
		} else if opponentMove == 2 {
			myMove = 1
		} else if opponentMove == 3 {
			myMove = 2
		}
	} else if myTargetMove == "Y" {
		opponentScore = 3
		myScore = 3
		myMove = opponentMove
	} else if myTargetMove == "Z" {
		opponentScore = 0
		myScore = 6
		if opponentMove == 1 {
			myMove = 2
		} else if opponentMove == 2 {
			myMove = 3
		} else if opponentMove == 3 {
			myMove = 1
		}
	}
	return
}
