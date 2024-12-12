package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cache = make(map[string]int)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func iterate(stones []int, n int) []int {

	for i := 0; i < n; i++ {
		output := []int{}

		for _, stone := range stones {

			if stone == 0 {
				output = append(output, 1)
				continue
			}

			stringStone := fmt.Sprintf("%d", stone)
			if len(stringStone)%2 == 0 {
				intStone1, err := strconv.Atoi(stringStone[:len(stringStone)/2])
				check(err)
				intStone2, err := strconv.Atoi(stringStone[len(stringStone)/2:])
				check(err)
				output = append(output, intStone1)
				output = append(output, intStone2)
			} else {
				output = append(output, stone*2024)
			}

		}

		stones = output

	}

	return stones
}

func iterateCache(stones []int, n int) int {

	num := 0

	for _, stone := range stones {
		num += count(stone, n)
	}

	return num
}

func count(stone, steps int) int {

	key := fmt.Sprintf("%d-%d", stone, steps)

	if val, found := cache[key]; found {
		return val
	}

	if steps == 0 {
		cache[key] = 1
		return 1
	}

	if stone == 0 {
		cache[key] = count(1, steps-1)
		return cache[key]
	}

	strStone := strconv.Itoa(stone)
	length := len(strStone)

	if length%2 == 0 {
		left, right := strStone[:length/2], strStone[length/2:]
		leftStone, _ := strconv.Atoi(left)
		rightStone, _ := strconv.Atoi(right)
		cache[key] = count(leftStone, steps-1) + count(rightStone, steps-1)
		return cache[key]
	}

	cache[key] = count(stone*2024, steps-1)
	return cache[key]
}

func toInt(data []string) []int {

	output := []int{}
	for _, d := range data {
		d, err := strconv.Atoi(d)
		check(err)
		output = append(output, d)
	}
	return output
}

func main() {

	file, err := os.ReadFile("cmd/aoc11/data.txt")
	check(err)

	data := strings.Fields(string(file))

	fmt.Println(toInt(data))
	output := iterate(toInt(data), 25)

	fmt.Println(len(output), iterateCache(toInt(data), 75))

}
