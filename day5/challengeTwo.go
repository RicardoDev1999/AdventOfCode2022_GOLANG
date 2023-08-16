package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// 	       [H]         [S]         [D]
//     [S] [C]         [C]     [Q] [L]
//     [C] [R] [Z]     [R]     [H] [Z]
//     [G] [N] [H] [S] [B]     [R] [F]
// [D] [T] [Q] [F] [Q] [Z]     [Z] [N]
// [Z] [W] [F] [N] [F] [W] [J] [V] [G]
// [T] [R] [B] [C] [L] [P] [F] [L] [H]
// [H] [Q] [P] [L] [G] [V] [Z] [D] [B]
//  1   2   3   4   5   6   7   8   9

type stack struct {
	Crates []rune
}

type move struct {
	Value       int
	Source      int
	Destination int
}

// FROM BOTTOM TO TOP
func initStack() []stack {
	return []stack{
		{Crates: []rune{'H', 'T', 'Z', 'D'}},
		{Crates: []rune{'Q', 'R', 'W', 'T', 'G', 'C', 'S'}},
		{Crates: []rune{'P', 'B', 'F', 'Q', 'N', 'R', 'C', 'H'}},
		{Crates: []rune{'L', 'C', 'N', 'F', 'H', 'Z'}},
		{Crates: []rune{'G', 'L', 'F', 'Q', 'S'}},
		{Crates: []rune{'V', 'P', 'W', 'Z', 'B', 'R', 'C', 'S'}},
		{Crates: []rune{'Z', 'F', 'J'}},
		{Crates: []rune{'D', 'L', 'V', 'Z', 'R', 'H', 'Q'}},
		{Crates: []rune{'B', 'H', 'G', 'N', 'F', 'Z', 'L', 'D'}},
	}
}

func printCrates(stacks []stack, currMove move) {
	var maxCratesArrSize = 0

	// Calculate max crates arr size, with move prediction.
	for i := 0; i < len(stacks); i++ {
		var stack = stacks[i]
		var lengthStackCrates = len(stack.Crates)

		if i == currMove.Destination-1 {
			lengthStackCrates += currMove.Value
		}

		if lengthStackCrates > maxCratesArrSize {
			maxCratesArrSize = lengthStackCrates
		}
	}

	const cRed = "\033[0;31m"
	const cGreen = "\033[32m"
	const cNone = "\033[0m"

	// Print crates with according movements.
	for y := maxCratesArrSize - 1; y >= 0; y-- {
		fmt.Print(" ")
		for x := 0; x < len(stacks); x++ {
			var stack = stacks[x]

			if len(stack.Crates) > y {
				if x == currMove.Source-1 && y >= len(stack.Crates)-currMove.Value {
					//This is the stack that will get crates moved from
					fmt.Fprintf(os.Stdout, "%s[%s]%s ", cRed, string(stack.Crates[y]), cNone)
					continue
				} else {
					fmt.Print("[", string(stack.Crates[y]), "] ")
					continue
				}
			} else if x == currMove.Destination-1 && y <= (len(stack.Crates)-1)+currMove.Value {
				//This is the stack that will get crates moved to
				var index = y - (len(stack.Crates) - 1)
				var sourceStack = stacks[currMove.Source-1]
				var crateToMove = sourceStack.Crates[(len(sourceStack.Crates)-currMove.Value)+(index-1)]
				fmt.Fprintf(os.Stdout, "%s[%v]%s ", cGreen, string(crateToMove), cNone)
				continue
			}

			fmt.Print("[ ] ")
		}
		fmt.Println()
	}

	for i := 0; i < len(stacks); i++ {
		fmt.Printf(" |%v|", i+1)
	}

	fmt.Println()
}

func main() {
	var stacks = initStack()
	var moves = loadMoves()

	for i := 0; i < len(moves); i++ {
		var move = moves[i]

		fmt.Println("Move:", move)
		printCrates(stacks, move)

		var stackToRemove = &stacks[move.Source-1]
		var removeFrom = len(stackToRemove.Crates) - move.Value

		var cratesToMove = stackToRemove.Crates[removeFrom:]
		stackToRemove.Crates = stackToRemove.Crates[:removeFrom]

		var stackToAppend = &stacks[move.Destination-1]
		stackToAppend.Crates = append(stackToAppend.Crates, cratesToMove...)

		var b []byte = make([]byte, 1)
		os.Stdin.Read(b)
	}
}

func loadMoves() []move {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var moves []move

	for scanner.Scan() {
		var line = scanner.Text()

		if len(line) > 0 {
			// Create a new template and parse the template string
			var toMove, from, to int
			fmt.Sscanf(line, "move %d from %d to %d", &toMove, &from, &to)
			moves = append(moves, move{Value: toMove, Source: from, Destination: to})
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return moves
}
