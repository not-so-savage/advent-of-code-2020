package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type policy struct {
	char string
	num1 int
	num2 int
}

type counts struct {
	validNumOfOccurrencesCount int
	validCharPositioningCount  int
}

func main() {
	inputArr, err := extractInputToArray("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	validPasswordsCounts, err := countValidPasswords(inputArr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Valid passwords count (occurrences policy): ", validPasswordsCounts.validNumOfOccurrencesCount)
	fmt.Println("Valid passwords count (positioning policy): ", validPasswordsCounts.validCharPositioningCount)
}

func countValidPasswords(passwordsArr []string) (counts, error) {
	var validNumOfOccurrencesCount int
	var validCharPositioningCount int

	for _, line := range passwordsArr {
		policy, err := extractPolicy(line)
		if err != nil {
			log.Fatal(err)
		}
		password, err := extractPassword(line)
		if err != nil {
			log.Fatal(err)
		}

		isValidNumOfOccurrences := checkNumOfOccurrencesValidity(policy, password)
		if isValidNumOfOccurrences {
			validNumOfOccurrencesCount++
		}

		isValidCharPositioning := checkCharPositioningValidity(policy, password)
		if isValidCharPositioning {
			validCharPositioningCount++
		}
	}
	return counts{validNumOfOccurrencesCount, validCharPositioningCount}, nil
}

func checkNumOfOccurrencesValidity(policy policy, password string) bool {
	actualOccurrences := strings.Count(password, policy.char)
	if actualOccurrences < policy.num1 || actualOccurrences > policy.num2 {
		return false
	}
	return true
}

func checkCharPositioningValidity(policy policy, password string) bool {
	if !strings.Contains(password, policy.char) {
		return false
	}

	for i, char := range password {
		if string(char) == policy.char && (i+1) == policy.num1 {
			for j, char := range password {
				if j != i {
					if string(char) == policy.char && (j+1) == policy.num2 {
						return false
					}
				}
			}
			return true
		}
		if string(char) == policy.char && (i+1) == policy.num2 {
			return true
		}
	}
	return false
}

func extractPolicy(dbEntry string) (policy, error) {
	policyFragment := strings.Split(dbEntry, ":")[0]

	num1, err := strconv.Atoi(strings.Split(policyFragment, "-")[0])
	if err != nil {
		log.Fatal(err)
	}

	num2, err := strconv.Atoi(strings.Split(strings.TrimPrefix(policyFragment, strconv.Itoa(num1)+"-"), " ")[0])
	if err != nil {
		log.Fatal(err)
	}

	char := strings.Split(policyFragment, " ")[1]

	return policy{
		char: char,
		num1: num1,
		num2: num2,
	}, nil
}

func extractPassword(dbEntry string) (string, error) {
	password := strings.TrimSpace(strings.Split(dbEntry, ":")[1])
	if len(password) < 1 {
		return "", fmt.Errorf("no password present in entry: %s", dbEntry)
	}
	return password, nil
}

func extractInputToArray(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	resultArr := []string{}

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				return nil, err
			}
			break
		}

		resultArr = append(resultArr, line)
	}

	return resultArr, nil
}
