package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// A / X = ROCK = 1 POINT
// B / Y = PAPER = 2 POINT
// C / Z = SCISSORS = 3 POINT

type play struct {
	enemyPlay  int
	yourPlay   int
	enemyScore int
	yourScore  int
}

const (
	SHOULD_LOSE = "X"
	SHOULD_DRAW = "Y"
	SHOULD_WIN  = "Z"
)

const (
	ROCK     = 1
	PAPER    = 2
	SCISSORS = 3
)

const (
	LOST = 0
	DRAW = 3
	WIN  = 6
)

var plays []play

func main() {
	loadPlays()
	loadScores()

	challengeTwo()
}

func loadPlays() {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	enemyPlays := map[string]int{"A": ROCK, "B": PAPER, "C": SCISSORS}

	for scanner.Scan() {
		var line = scanner.Text()
		if len(line) > 0 {
			var enemyPlay = enemyPlays[string(line[0])]
			var yourPlay = ROCK

			var should = string(line[2])

			if should == SHOULD_WIN {
				if enemyPlay == SCISSORS {
					yourPlay = ROCK
				}
				if enemyPlay == ROCK {
					yourPlay = PAPER
				}
				if enemyPlay == PAPER {
					yourPlay = SCISSORS
				}
			}
			if should == SHOULD_DRAW {
				yourPlay = enemyPlay
			}
			if should == SHOULD_LOSE {
				if enemyPlay == ROCK {
					yourPlay = SCISSORS
				}
				if enemyPlay == PAPER {
					yourPlay = ROCK
				}
				if enemyPlay == SCISSORS {
					yourPlay = PAPER
				}
			}

			plays = append(plays, play{enemyPlay: enemyPlay, yourPlay: yourPlay})
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func loadScores() {
	for i := 0; i < len(plays); i++ {
		var result = LOST

		if plays[i].yourPlay == plays[i].enemyPlay {
			result = DRAW
		}

		if (plays[i].yourPlay == ROCK && plays[i].enemyPlay == SCISSORS) ||
			(plays[i].yourPlay == PAPER && plays[i].enemyPlay == ROCK) ||
			(plays[i].yourPlay == SCISSORS && plays[i].enemyPlay == PAPER) {
			result = WIN
		}

		if result == LOST {
			plays[i].yourScore = LOST + plays[i].yourPlay
			plays[i].enemyScore = WIN + plays[i].enemyPlay
		}
		if result == WIN {
			plays[i].yourScore = WIN + plays[i].yourPlay
			plays[i].enemyScore = LOST + plays[i].enemyPlay
		}
		if result == DRAW {
			plays[i].yourScore = DRAW + plays[i].yourPlay
			plays[i].enemyScore = DRAW + plays[i].enemyPlay
		}
	}
}

func challengeTwo() {
	var score int = 0
	for i := 0; i < len(plays); i++ {
		fmt.Println(plays[i])
		score += plays[i].yourScore
	}
	fmt.Println(score)
}
