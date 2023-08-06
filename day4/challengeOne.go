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

	var fullyContainedPairs int32

	for _, pair := range pairs {
		var sectionOneFirst = pair.elfOne.sections[0]
		var sectionOneLast = pair.elfOne.sections[len(pair.elfOne.sections)-1]
		var sectionTwoFirst = pair.elfTwo.sections[0]
		var sectionTwoLast = pair.elfTwo.sections[len(pair.elfTwo.sections)-1]

		if (sectionOneFirst <= sectionTwoFirst && sectionOneLast >= sectionTwoLast) ||
			(sectionTwoFirst <= sectionOneFirst && sectionTwoLast >= sectionOneLast) {
			fullyContainedPairs++
		}
	}

	fmt.Println(fullyContainedPairs)
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
