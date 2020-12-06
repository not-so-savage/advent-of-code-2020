package main

import (
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

func main() {
	fileBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(fileBytes)
	boardingPasses := strings.Split(input, "\n")

	totalRows := 128
	totalColumns := 8

	seatIDs := make([]int, len(boardingPasses))

	for _, pass := range boardingPasses {
		rows := make([]int, totalRows)
		for i := range rows {
			rows[i] = i
		}

		columns := make([]int, totalColumns)
		for i := range columns {
			columns[i] = i
		}

		for _, char := range pass {
			switch char {
			case 'F':
				rows = rows[:len(rows)/2]
			case 'B':
				rows = rows[len(rows)/2:]
			case 'L':
				columns = columns[:len(columns)/2]
			case 'R':
				columns = columns[len(columns)/2:]
			}
		}

		seatIDs = append(seatIDs, rows[0]*8+columns[0])
	}

	sort.Ints(seatIDs)
	log.Print("Largest ID: ", findLargestElement(seatIDs))

	for i := range seatIDs {
		if seatIDs[i] == 0 || seatIDs[i-1] == 0 {
			continue
		}

		if (seatIDs[i] - seatIDs[i-1]) != 1 {
			log.Println("Your seat ID is: ", seatIDs[i]-1)
		}

	}
}

func findLargestElement(arr []int) int {
	max := arr[0]
	for _, elem := range arr {
		if elem > max {
			max = elem
		}
	}
	return max
}
