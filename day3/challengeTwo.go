package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type rucksack struct {
	line string
}

var rucksacks []rucksack

func main() {
	loadRucksacks()

	var points int

	for i := 0; i < len(rucksacks); i += 3 {
		var fst = rucksacks[i]
		var scd = rucksacks[i+1]
		var thr = rucksacks[i+2]

		var fsscd = compareRucksacks(fst, scd)
		var scdth = compareRucksacks(scd, thr)
		var result = compareRucksacks(rucksack{fsscd}, rucksack{scdth})

		itemTypesArray := strings.Split(result, "")

		for _, itemType := range itemTypesArray {
			points += getPriority([]rune(itemType)[0])
		}
	}

	fmt.Println(points)
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
			var rs = rucksack{line}
			rucksacks = append(rucksacks, rs)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func compareRucksacks(this rucksack, next rucksack) string {
	var eqLetters string

	for i := 0; i < len(this.line); i++ {
		for j := 0; j < len(next.line); j++ {
			if this.line[i] == next.line[j] {
				var letter = string(this.line[i])
				if !strings.Contains(eqLetters, letter) {
					eqLetters += letter
				}
			}
		}
	}

	return eqLetters
}
