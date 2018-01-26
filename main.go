package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type Problem struct {
	Question string
	Solution string
}

func main() {
	fmt.Println("Welcode to the quiz! Press [Enter] to start.")

	problemsFilePath := flag.String("csv", "./problems.csv", "The problems.csv file path.")
	timeLimitInSeconds := flag.Int("limit", 3000, "The time limit for the quiz.")
	flag.Parse()
	problems := getProblemsFromFile(*problemsFilePath)
	timeLimit := time.Duration(*timeLimitInSeconds) * time.Second

	score := quiz(problems, timeLimit)
	fmt.Printf("\nFinal score: %d/%d.", score, len(problems))
}

func quiz(problems []Problem, timeLimit time.Duration) int {
	correctAnswers := 0
	solutionsChan := make(chan bool)
	doneChan := make(chan bool)

	fmt.Scanf("\n")
	go getSolutions(problems, solutionsChan, doneChan)
	quizTimer := time.NewTimer(timeLimit)
	for {
		select {
		case solution := <-solutionsChan:
			if solution {
				correctAnswers++
			}
		case <-doneChan:
			fmt.Println("\nOut of problems!")
			return correctAnswers
		case <-quizTimer.C:
			fmt.Println("\nTime's run out!")
			return correctAnswers
		}
	}

	return correctAnswers
}
func getSolutions(problems []Problem, solutions chan<- bool, done chan<- bool) {
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s\t=\t", i, problem.Question)

		var userSolution string
		fmt.Scanf("%s\n", &userSolution)

		userSolution = strings.TrimSpace(userSolution)
		if userSolution == problem.Solution {
			solutions <- true
		} else {
			solutions <- false
		}
	}
	done <- true
}

func getProblemsFromFile(problemsFilePath string) []Problem {
	var problems []Problem

	problemsFile, err := os.Open(problemsFilePath)
	if err != nil {
		panic(err)
	}
	csvReader := csv.NewReader(problemsFile)
	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		problems = append(problems, Problem{Question: line[0], Solution: line[1]})
	}
	return problems
}
