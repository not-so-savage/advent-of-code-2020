package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var input, _ = ioutil.ReadFile("input.txt")
var inputMap = string(input)
var rows = strings.Split(inputMap, "\n")
var width = len(rows[0])

func main() {
	partATrees := traverseMap(3, 1)
	fmt.Println(partATrees)

	partBAnswer := traverseMap(1, 1) * traverseMap(3, 1) * traverseMap(5, 1) * traverseMap(7, 1) * traverseMap(1, 2)
	fmt.Println(partBAnswer)
}

func traverseMap(xStep int, yStep int) int {
	treeCount := 0

	x := 0
	y := 0

	for y < len(rows) {
		if rows[y][x] == '#' {
			treeCount++
		}

		x = x + xStep
		y = y + yStep

		if x >= width {
			x = x - width
		}
	}

	return treeCount
}
