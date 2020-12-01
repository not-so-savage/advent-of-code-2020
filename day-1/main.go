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

type pair struct {
	a int
	b int
}

type trio struct {
	a int
	b int
	c int
}

func main() {
	input, err := extractLinesToArray()
	if err != nil {
		log.Fatal(err)
	}

	pair, err := findCorrectPair(input)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Pair is: ", pair.a, pair.b)

	trio, err := findCorrectTrio(input)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Trio is: ", trio.a, trio.b, trio.c)

	answerA := pair.a * pair.b
	answerB := (trio.a * trio.b) * trio.c
	log.Print("Answer A is: ", answerA)
	log.Print("Answer B is: ", answerB)
}

func extractLinesToArray() ([]int, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	inputSlice := []int{}
	reader := bufio.NewReader(file)
	for {
		inputString, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				return nil, err
			}
			break
		}
		inputString = strings.TrimSuffix(inputString, "\n")

		input, err := strconv.Atoi(inputString)
		if err != nil {
			return nil, err
		}

		inputSlice = append(inputSlice, input)
	}

	return inputSlice, nil
}

func findCorrectPair(input []int) (pair, error) {
	for _, num1 := range input {
		for _, num2 := range input {
			if (num1 + num2) == 2020 {
				return pair{num1, num2}, nil
			}
		}
	}
	return pair{}, fmt.Errorf("no pair found")
}

func findCorrectTrio(input []int) (trio, error) {
	for _, num1 := range input {
		for _, num2 := range input {
			for _, num3 := range input {
				if (num1+num2)+num3 == 2020 {
					return trio{num1, num2, num3}, nil
				}
			}
		}
	}
	return trio{}, fmt.Errorf("no pair found")
}
