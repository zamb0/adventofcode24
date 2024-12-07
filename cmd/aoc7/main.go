package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Eqne struct {
	left  int
	right []int
}

func readFile(filename string) ([]Eqne, error) {
	file, err := os.Open(filename)
	check(err)

	equations := []Eqne{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()
		right := []int{}

		split := strings.Split(line, ":")

		left, err := strconv.Atoi(split[0])
		check(err)

		right_s := strings.Fields(split[1])

		for _, r := range right_s {
			right_i, err := strconv.Atoi(r)
			check(err)
			right = append(right, right_i)
		}

		equations = append(equations, Eqne{left: left, right: right})

		if err := scanner.Err(); err != nil {
			return nil, err
		}
	}

	return equations, nil

}

func recursiveCheck1(equation Eqne, current int, index int) bool {

	if index == len(equation.right) {
		return current == equation.left
	}

	if recursiveCheck1(equation, current+equation.right[index], index+1) {
		return true
	}

	if recursiveCheck1(equation, current*equation.right[index], index+1) {
		return true
	}

	return false

}

func recursiveCheck2(equation Eqne, current int, index int) bool {

	if index == len(equation.right) {
		return current == equation.left
	}

	if recursiveCheck2(equation, current+equation.right[index], index+1) {
		return true
	}

	if recursiveCheck2(equation, current*equation.right[index], index+1) {
		return true
	}

	concat, err := strconv.Atoi(fmt.Sprintf("%d%d", current, equation.right[index]))
	check(err)

	return recursiveCheck2(equation, concat, index+1)

}

func checkOperators(equation Eqne, part int) bool {

	if len(equation.right) == 0 {
		return false
	}

	if part == 1 {

		return recursiveCheck1(equation, equation.right[0], 1)

	} else {

		return recursiveCheck2(equation, equation.right[0], 1)

	}
}

func main() {

	equations, err := readFile("cmd/aoc7/data.txt")
	check(err)

	sum1 := 0
	sum2 := 0

	for _, eq := range equations {
		if checkOperators(eq, 1) {
			sum1 += eq.left
		}

		if checkOperators(eq, 2) {
			sum2 += eq.left
		}
	}

	fmt.Println(sum1, sum2)
}
