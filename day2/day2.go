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

	totalOpponentScore := 0
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
		myMove := parts[1]

		opponentScore := scores[opponentMove]
		roundOpponentScore += opponentScore
		myScore := scores[myMove]
		roundMyScore += myScore

		opponentResult, myResult := determineWin(opponentScore, myScore)

		roundOpponentScore += opponentResult
		roundMyScore += myResult
		totalOpponentScore += roundOpponentScore
		totalMyScore += roundMyScore

		fmt.Fprintln(os.Stdout, opponentMove, myMove, opponentScore, myScore, opponentResult, myResult, roundOpponentScore, roundMyScore)
	}
	fmt.Fprintln(os.Stdout, "opponent total score:", totalOpponentScore)
	fmt.Fprintln(os.Stdout, "my total score: ", totalMyScore)
}

func determineWin(player1 int, player2 int) (player1Score int, player2Score int) {
	if (player1 == 1 && player2 == 3) || player1-player2 == 1 {
		player1Score = 6
		player2Score = 0
	} else if (player1 == 3 && player2 == 1) || player2-player1 == 1 {
		player1Score = 0
		player2Score = 6
	} else {
		player1Score = 3
		player2Score = 3
	}
	return
}
