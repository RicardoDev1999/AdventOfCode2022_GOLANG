package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type rucksack struct {
	compartmentOne string
	compartmentTwo string
}

var rucksacks []rucksack

func main() {
	loadRucksacks()

	var totalPoints int

	for i := 0; i < len(rucksacks); i++ {
		var rs = rucksacks[i]
		totalPoints += compareCompartments(rs)
	}

	fmt.Println(totalPoints)
}

func getPriority(itemType rune) int {
	if 'a' <= itemType && itemType <= 'z' {
		// Lowercase letters have priorities 1 through 26
		return int(itemType - 'a' + 1)
	} else if 'A' <= itemType && itemType <= 'Z' {
		// Uppercase letters have priorities 27 through 52
		return int(itemType - 'A' + 27)
	}
	// Return 0 for non-letter characters or unexpected input
	return 0
}

func loadRucksacks() {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		var line = scanner.Text()
		if len(line) > 0 {

			var half = len(line) / 2
			var compartmentOne = line[:half]
			var compartmentTwo = line[half:]

			var rs = rucksack{compartmentOne, compartmentTwo}
			rucksacks = append(rucksacks, rs)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func compareCompartments(rs rucksack) int {
	var points int
	var eqLetters string

	for i := 0; i < len(rs.compartmentOne); i++ {
		for j := 0; j < len(rs.compartmentTwo); j++ {
			if rs.compartmentOne[i] == rs.compartmentTwo[j] {
				var letter = string(rs.compartmentOne[i])
				if !strings.Contains(eqLetters, letter) {
					eqLetters += letter
				}
			}
		}
	}

	itemTypesArray := strings.Split(eqLetters, "")

	for _, itemType := range itemTypesArray {
		points += getPriority([]rune(itemType)[0])
	}

	return points
}
