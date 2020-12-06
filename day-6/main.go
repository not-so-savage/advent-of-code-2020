package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileString := string(file)

	groups := strings.Split(fileString, "\n\n")

	groupUniqueAnsCount := 0
	groupAllAnsweredCount := 0

	for _, group := range groups {
		answerSets := strings.Split(group, "\n")
		uniqueAnswers := []string{}

		for _, answers := range answerSets {
			for _, answer := range answers {
				if checkIfUniqueAnswer(uniqueAnswers, string(answer)) {
					uniqueAnswers = append(uniqueAnswers, string(answer))
					groupUniqueAnsCount++
				}
			}
		}

		for _, uniqueAns := range uniqueAnswers {
			if checkAnsIsInEachSet(uniqueAns, answerSets) {
				groupAllAnsweredCount++
			}
		}
	}

	log.Println("Part 1: ", groupUniqueAnsCount)
	log.Println("Part 2: ", groupAllAnsweredCount)
}

func checkAnsIsInEachSet(ans string, sets []string) bool {
	for _, set := range sets {
		if !checkAnswerIsInSet(ans, set) {
			return false
		}
	}
	return true
}

func checkAnswerIsInSet(queryAns string, answers string) bool {
	for _, answer := range answers {
		if string(answer) == queryAns {
			return true
		}
	}
	return false
}

func checkIfUniqueAnswer(uniqueAnswers []string, ans string) bool {
	for i := range uniqueAnswers {
		if uniqueAnswers[i] == ans {
			return false
		}
	}
	return true
}
