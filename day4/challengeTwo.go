package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type elfPair struct {
	elfOne assignent
	elfTwo assignent
}

type assignent struct {
	sections []int
}

var pairs []elfPair

func main() {
	loadAssignments()

	var totalOverlap uint32

	// Look through each pair
	// For each section in elfOne
	// Look through each section in elfTwo
	// If there is a match, increment totalOverlap and break out of the loop

	for _, pair := range pairs {
		for _, sectionA := range pair.elfOne.sections {
			var overlapped = false
			for _, sectionB := range pair.elfTwo.sections {
				if sectionA == sectionB {
					overlapped = true
				}
			}
			if overlapped {
				totalOverlap++
				break
			}
		}
	}

	fmt.Println(totalOverlap)
}

func loadAssignments() {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		var line = scanner.Text()
		if len(line) > 0 {
			//split line by ","
			//split each section by "-"
			//convert each section to int

			var parts = strings.Split(line, ",")
			var sectionsPartOne = strings.Split(parts[0], "-")
			var sectionsPartTwo = strings.Split(parts[1], "-")

			var sectionsOneFirstNumber, _ = strconv.Atoi(sectionsPartOne[0])
			var sectionsOneLastNumber, _ = strconv.Atoi(sectionsPartOne[len(sectionsPartOne)-1])

			var sectionsTwoFirstNumber, _ = strconv.Atoi(sectionsPartTwo[0])
			var sectionsTwoLastNumber, _ = strconv.Atoi(sectionsPartTwo[len(sectionsPartTwo)-1])

			var sectionOneRange []int
			for i := sectionsOneFirstNumber; i <= sectionsOneLastNumber; i++ {
				sectionOneRange = append(sectionOneRange, i)
			}

			var sectionTwoRange []int
			for i := sectionsTwoFirstNumber; i <= sectionsTwoLastNumber; i++ {
				sectionTwoRange = append(sectionTwoRange, i)
			}

			var pair = elfPair{
				elfOne: assignent{sections: sectionOneRange},
				elfTwo: assignent{sections: sectionTwoRange},
			}

			pairs = append(pairs, pair)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
