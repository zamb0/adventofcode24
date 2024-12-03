package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type LoL struct {
	data []int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func derivative(data []int) []int {

	deriv := []int{}

	for i := 1; i < len(data); i++ {
		deriv = append(deriv, int(math.Abs(float64(data[i])-float64(data[i-1]))))
	}

	return deriv
}

func isSortedA(p []int) bool {
	s := sort.SliceIsSorted(p, func(i, j int) bool {
		return p[i] < p[j]
	})

	return s
}

func isSortedD(p []int) bool {
	s := sort.SliceIsSorted(p, func(i, j int) bool {
		return p[i] > p[j]
	})

	return s
}

func minDelta(p []int, min int) bool {
	m := p[0]
	for _, e := range p {
		if e < m {
			m = e
		}
	}

	return m >= min
}

func maxDelta(p []int, max int) bool {
	m := p[0]
	for _, e := range p {
		if e > m {
			m = e
		}
	}

	return m <= max
}

func allTests(d_data []int, data []int) bool {

	return (isSortedA(data) || isSortedD(data)) && minDelta(d_data, 1) && maxDelta(d_data, 3)

}

func s_check(d_data []int, data []int) bool {

	return allTests(d_data, data)
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func sd_check(data []int) bool {

	for d := range data {
		new_data := make([]int, len(data))
		copy(new_data, data)

		new_data = remove(new_data, d)

		fmt.Println(new_data)
		if allTests(derivative(new_data), new_data) {
			return true
		}
	}

	return false
}

func n_safe(data []string) int {

	int_data := []LoL{}
	sum_ok := 0

	for row := range data {
		row_data := strings.Split(data[row], " ")
		int_row_data := LoL{}
		for d := range row_data {
			x, err := strconv.Atoi(row_data[d])
			check(err)
			int_row_data.data = append(int_row_data.data, x)
		}
		int_data = append(int_data, int_row_data)
	}

	for row := range int_data {
		if s_check(derivative(int_data[row].data), int_data[row].data) {
			sum_ok += 1
		} else if sd_check(int_data[row].data) {
			sum_ok += 1
		}
	}

	return sum_ok

}

func main() {
	file, err := os.ReadFile("cmd/aoc2/data.txt")
	check(err)

	data := strings.Split(string(file), "\n")

	fmt.Println(n_safe(data))
}
