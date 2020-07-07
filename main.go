package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()
	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open CSV file: %s\n", *csvFileName))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the CSV file")
	}
	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problem.que)
		ansCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			ansCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d.\n", correct, len(problems))
			return
		case answer := <-ansCh:
			if answer == problem.ans {
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
			que: line[0],
			ans: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	que string
	ans string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
