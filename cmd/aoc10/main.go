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

func trailSearching(grid [][]rune, i, j int, pathLength int, visited map[string]bool) int {

	directions := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	score := 0

	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return 0
	}

	if visited[fmt.Sprintf("%d,%d", i, j)] {
		return 0
	}

	if pathLength == 9 && !visited[fmt.Sprintf("%d,%d", i, j)] {

		visited[fmt.Sprintf("%d,%d", i, j)] = true
		return 1
	}

	for _, direction := range directions {

		if i+direction[0] >= 0 && i+direction[0] < len(grid) && j+direction[1] >= 0 && j+direction[1] < len(grid[0]) {

			if string(grid[i+direction[0]][j+direction[1]]) == fmt.Sprintf("%d", pathLength+1) {

				score += trailSearching(grid, i+direction[0], j+direction[1], pathLength+1, visited)

			}
		}
	}

	return score
}

func trailSearching2(grid [][]rune, i, j int, pathLength int) int {

	directions := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	score := 0

	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return 0
	}

	if pathLength == 9 {

		return 1
	}

	for _, direction := range directions {

		if i+direction[0] >= 0 && i+direction[0] < len(grid) && j+direction[1] >= 0 && j+direction[1] < len(grid[0]) {

			if string(grid[i+direction[0]][j+direction[1]]) == fmt.Sprintf("%d", pathLength+1) {

				score += trailSearching2(grid, i+direction[0], j+direction[1], pathLength+1)

			}
		}
	}

	return score
}

func startsSearch(grid [][]rune) (int, int) {

	starts := findStart(grid)
	score := 0
	score2 := 0

	for _, start := range starts {
		i, j := start[0], start[1]
		score += trailSearching(grid, i, j, 0, map[string]bool{})
		score2 += trailSearching2(grid, i, j, 0)
	}

	return score, score2
}

func findStart(grid [][]rune) [][2]int {

	starts := [][2]int{}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '0' {
				starts = append(starts, [2]int{i, j})
			}
		}
	}

	return starts
}

func main() {

	grid, err := readMatrixFromFile("cmd/aoc10/data.txt")
	check(err)

	score, score2 := startsSearch(grid)
	fmt.Println(score, score2)

}
