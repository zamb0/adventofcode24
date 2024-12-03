package main

import (
	"fmt"
	"math"
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

func sort_asc(a []string) []string {
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})

	return a
}

func sum_diff(a []string, b []string) int {

	sum := 0

	for i := range a {
		x, err := strconv.Atoi(a[i])
		check(err)
		y, err := strconv.Atoi(b[i])
		check(err)
		sum += int(math.Abs(float64(x) - float64(y)))
	}

	return sum

}

func sum_sim(a []string, b []string) int {

	aa := map[string]int{}
	bb := map[string]int{}
	sum := 0

	for i := range a {
		_, exista := aa[a[i]]

		if !exista {
			aa[a[i]] = 1
		} else {
			aa[a[i]] += 1
		}

		_, existb := bb[b[i]]
		if !existb {
			bb[b[i]] = 1
		} else {
			bb[b[i]] += 1
		}
	}

	for i, j := range aa {
		b, exist := bb[i]
		if exist {
			x, err := strconv.Atoi(i)
			check(err)
			sum += x * b * j
		}

	}

	return sum
}

func main() {
	file, err := os.ReadFile("cmd/aoc1/data.txt")
	check(err)

	data := strings.Fields(string(file))

	a := []string{}
	b := []string{}

	for i, d := range data {
		if i%2 == 0 {
			a = append(a, d)
		} else {
			b = append(b, d)
		}
	}

	a = sort_asc(a)
	b = sort_asc(b)

	//Question 1
	fmt.Println(sum_diff(a, b))
	//Question 2
	fmt.Println(sum_sim(a, b))

}
