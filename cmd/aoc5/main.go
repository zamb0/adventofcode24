package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(filename string) ([][]int, [][]int, error) {
	file, err := os.Open(filename)
	check(err)

	var rules [][]int
	var pages [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		split := strings.Split(line, "|")

		rule1, err := strconv.Atoi(split[0])
		check(err)
		rule2, err := strconv.Atoi(split[1])
		check(err)

		rules = append(rules, []int{rule1, rule2})

	}

	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, ",")

		var page []int

		for _, s := range split {
			num, err := strconv.Atoi(s)
			check(err)
			page = append(page, num)
		}

		pages = append(pages, page)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}
	return rules, pages, nil
}

func checkRule(cache map[string]bool, page []int) bool {

	for i := 0; i < len(page); i++ {
		for j := i + 1; j < len(page); j++ {

			if !cache[fmt.Sprintf("%d,%d", page[i], page[j])] {
				return false
			}

		}
	}

	return true

}

func orderPage(cache map[string]bool, page []int) int {

	sort.SliceStable(page, func(i, j int) bool {
		return cache[fmt.Sprintf("%d,%d", page[i], page[j])]
	})

	return page[int(len(page)/2)]

}

func main() {

	rules, pages, err := readFile("cmd/aoc5/data.txt")
	check(err)

	cache := make(map[string]bool)

	for _, r := range rules {

		cache[fmt.Sprintf("%d,%d", r[0], r[1])] = true
		cache[fmt.Sprintf("%d,%d", r[1], r[0])] = false

	}

	countCorrect := 0
	countOrdered := 0
	for _, page := range pages {
		if checkRule(cache, page) {
			countCorrect += page[int(len(page)/2)]
		} else {
			countOrdered += orderPage(cache, page)
		}
	}

	fmt.Println(countCorrect, countOrdered)
}
