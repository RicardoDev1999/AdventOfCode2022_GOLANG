package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type elf struct {
	calories int
}

var elfs []elf

func main() {
	loadElfs()
	orderByAscending()

	challengeOne()
	challengeTwo()
}

func loadElfs() {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var sumCalories int = 0

	for scanner.Scan() {
		var line = scanner.Text()
		if len(line) > 0 {
			calories, _ := strconv.Atoi(line)
			sumCalories += calories
			continue
		}
		elfs = append(elfs, elf{calories: sumCalories})
		sumCalories = 0
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func orderByAscending() {
	for i := 0; i < len(elfs); i++ {
		for j := 0; j < len(elfs); j++ {
			if elfs[j].calories > elfs[i].calories {
				var temp elf = elfs[i]
				elfs[i] = elfs[j]
				elfs[j] = temp
			}
		}
	}
}

func challengeOne() {
	fmt.Println("Challenge One")

	fmt.Println(elfs[len(elfs)-1].calories)

	fmt.Println("--------------")
}

func challengeTwo() {
	fmt.Println("Challenge Two")

	var first = elfs[len(elfs)-1].calories
	var second = elfs[len(elfs)-2].calories
	var third = elfs[len(elfs)-3].calories
	fmt.Println(first + second + third)

	fmt.Println("--------------")
}
