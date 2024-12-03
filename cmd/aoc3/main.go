package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func do_mult(str string) int {
	pattern := regexp.MustCompile(`\d{1,3},\d{1,3}`)
	ab := strings.Split(pattern.FindAllString(str, -1)[0], ",")
	a, err := strconv.Atoi(ab[0])
	check(err)
	b, err := strconv.Atoi(ab[1])
	check(err)

	return a * b
}

func main() {
	file, err := os.ReadFile("cmd/aoc3/data.txt")
	check(err)

	data := string(file)

	pattern := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	pattern_new := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)

	mult := pattern.FindAllString(data, -1)
	mult_new := pattern_new.FindAllString(data, -1)

	sum := 0
	sum_new := 0

	for m := range mult {
		sum += do_mult(mult[m])
	}

	run := true
	for m := range mult_new {

		if mult_new[m] == "do()" {
			run = true
			continue
		}

		if mult_new[m] == "don't()" {
			run = false
			continue
		}

		if run {
			sum_new += do_mult(mult_new[m])
		}

	}

	fmt.Println(sum, sum_new)
}
