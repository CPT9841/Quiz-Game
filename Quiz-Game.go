package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", " a csv file in the the format of 'question,answer'")

	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		fmt.Printf("Failed to pen the CSV File: %s\n", *csvFilename)
		os.Exit(1)
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	timer := time.NewTimer(time.Duration(120) * time.Second)

	problems := parseLines(lines)
	correct := 0

	for i, p := range problems {
		select {
		case <-timer.C:
			fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
			return
		default:
			fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
			fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
			var answer string
			fmt.Scanf("%s\n", &answer)
			if answer == p.a {
				correct++
			}

		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))

}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
