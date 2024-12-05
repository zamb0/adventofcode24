package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func patternSearch1(grid [][]rune) int {

	pCount := 0

	for row := range grid {
		for col := range grid[0] {
			if grid[row][col] != 'X' { //skip if not starting with X
				continue
			}
			for _, dr := range []int{-1, 0, 1} {
				for _, dc := range []int{-1, 0, 1} {
					if dr == 0 && dc == 0 {
						continue
					}

					dimr := row + 3*dr
					dimc := col + 3*dc

					if !(dimr >= 0 && dimr < len(grid) && dimc >= 0 && dimc < len(grid[0])) {
						continue
					}

					if grid[row+dr][col+dc] == 'M' && grid[row+2*dr][col+2*dc] == 'A' && grid[row+3*dr][col+3*dc] == 'S' {
						pCount += 1
					}
				}
			}

		}
	}

	return pCount
}

func patternSearch2(grid [][]rune) int {

	pCount := 0

	for row := 1; row < len(grid)-1; row++ {
		for col := 1; col < len(grid[0])-1; col++ {

			if grid[row][col] != 'A' { //skip if not starting with A
				continue
			}

			corners := []rune{grid[row-1][col-1], grid[row-1][col+1], grid[row+1][col+1], grid[row+1][col-1]}

			if corners[0] == 'M' && corners[1] == 'M' && corners[2] == 'S' && corners[3] == 'S' {
				pCount += 1
			}

			if corners[0] == 'S' && corners[1] == 'S' && corners[2] == 'M' && corners[3] == 'M' {
				pCount += 1
			}

			if corners[0] == 'M' && corners[1] == 'S' && corners[2] == 'S' && corners[3] == 'M' {
				pCount += 1
			}

			if corners[0] == 'S' && corners[1] == 'M' && corners[2] == 'M' && corners[3] == 'S' {
				pCount += 1
			}

		}
	}

	return pCount
}

func readMatrixFromFile(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	check(err)

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return grid, nil
}

func main() {

	matrix, err := readMatrixFromFile("cmd/aoc4/data.txt")
	check(err)

	result1 := patternSearch1(matrix)
	result2 := patternSearch2(matrix)

	fmt.Println(result1, result2)

}
