package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type Problem struct {
	Question string
	Solution string
}

func main() {
	fmt.Println("Welcode to the quiz!")

	problemsFilePath := flag.String("f", "./problems.csv", "The problems.csv file path.")
	flag.Parse()
	problems := getProblemsFromFile(*problemsFilePath)

	correctAnswers := 0
	stdinReader := bufio.NewReader(os.Stdin)
	for index, problem := range problems {
		fmt.Printf("Problem #%d:\t%s\t=\t", index, problem.Question)
		userSolution, err := stdinReader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		userSolution = strings.TrimSpace(userSolution)
		if userSolution == problem.Solution {
			correctAnswers = correctAnswers + 1
			fmt.Println("Good")
		} else {
			fmt.Println("Mistake")
		}
	}

	fmt.Printf("\nFinal score: %d/%d.", correctAnswers, len(problems))
}

func getProblemsFromFile(problemsFilePath string) []Problem {
	var problems []Problem

	problemsFile, err := os.Open(problemsFilePath)
	if err != nil {
		panic(err)
	}
	csvReader := csv.NewReader(bufio.NewReader(problemsFile))
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
