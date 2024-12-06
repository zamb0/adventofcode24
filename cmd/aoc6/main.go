package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func navigationInMatrix(matrixr [][]rune, loopSearch bool) ([]string, bool) {

	//search for the starting point ^
	//go in the direction of the arrow until you reach # or the end of the matrix
	//if you reach #, turn right
	//if you reach the end of the matrix, stop
	//count distinct positions visited

	memory := make([]string, 0)
	directions := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	direction := 0
	start := [2]int{0, 0}
	countChangeDirection := 0

	matrix := deepCopyMatrix(matrixr)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {

			if matrix[i][j] == '^' {
				start = [2]int{i, j}
				direction = 0
				matrix[i][j] = '.'
				break
			}

			if matrix[i][j] == '>' {
				start = [2]int{i, j}
				direction = 1
				matrix[i][j] = '.'
				break
			}

			if matrix[i][j] == 'v' {
				start = [2]int{i, j}
				direction = 2
				matrix[i][j] = '.'
				break
			}

			if matrix[i][j] == '<' {
				start = [2]int{i, j}
				direction = 3
				matrix[i][j] = '.'
				break
			}

		}
	}

	for {

		if start[0]+directions[direction][0] >= len(matrix) || start[1]+directions[direction][1] >= len(matrix[0]) || start[0]+directions[direction][0] < 0 || start[1]+directions[direction][1] < 0 {
			break
		}

		if loopSearch && countChangeDirection > 4 {

			if slices.Contains(memory, fmt.Sprintf("%d,%d", start[0]+directions[direction][0], start[1]+directions[direction][1])) {
				return memory, true
			}
		}

		if matrix[start[0]+directions[direction][0]][start[1]+directions[direction][1]] == '#' || matrix[start[0]+directions[direction][0]][start[1]+directions[direction][1]] == 'O' {
			direction = (direction + 1) % 4
			countChangeDirection++

		} else if matrix[start[0]+directions[direction][0]][start[1]+directions[direction][1]] == '.' {
			start = [2]int{start[0] + directions[direction][0], start[1] + directions[direction][1]}
			if !slices.Contains(memory, fmt.Sprintf("%d,%d", start[0], start[1])) {
				memory = append(memory, fmt.Sprintf("%d,%d", start[0], start[1]))
				countChangeDirection = 0
			}
		}

	}

	return memory, false
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

func deepCopyMatrix(matrix [][]rune) [][]rune {
	copyMatrix := make([][]rune, len(matrix))
	for i := range matrix {
		copyMatrix[i] = make([]rune, len(matrix[i]))
		copy(copyMatrix[i], matrix[i])
	}
	return copyMatrix
}

func makeObastacle(matrix [][]rune, memory []string) int {

	loops := 0

	for key := range memory {

		newMatrix := deepCopyMatrix(matrix)

		mkey := [2]int{0, 0}
		fmt.Sscanf(memory[key], "%d,%d", &mkey[0], &mkey[1])

		if newMatrix[mkey[0]][mkey[1]] == '.' {
			newMatrix[mkey[0]][mkey[1]] = 'O'
		} else {
			continue
		}

		if _, l := navigationInMatrix(newMatrix, true); l {
			loops++
		}

	}

	return loops

}

func main() {

	matrix, err := readMatrixFromFile("cmd/aoc6/data.txt")
	check(err)

	memory, _ := navigationInMatrix(matrix, false)

	loops := makeObastacle(matrix, memory)

	fmt.Println(len(memory), loops)

}
