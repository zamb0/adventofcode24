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

func readMatrix(file *os.File) [][]rune {
	scanner := bufio.NewScanner(file)
	matrix := make([][]rune, 0)

	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, []rune(line))
	}

	file.Close()

	return matrix
}

func mapAntennas(matrix [][]rune) map[string][][]int {
	antennas := make(map[string][][]int)

	for i := range matrix {
		for j := range matrix[i] {
			if string(matrix[i][j]) != "." {
				antennas[string(matrix[i][j])] = append(antennas[string(matrix[i][j])], []int{i, j})
			}
		}
	}

	return antennas
}

func findAntinodes(antennas map[string][][]int, l int, h int) map[string]bool {
	antinodes := make(map[string]bool)

	for key := range antennas {
		for i := 0; i < len(antennas[key]); i++ {
			for j := i + 1; j < len(antennas[key]); j++ {

				a1y, a1x := antennas[key][i][0], antennas[key][i][1]
				a2y, a2x := antennas[key][j][0], antennas[key][j][1]

				dy, dx := a2y-a1y, a2x-a1x

				an1y, an1x := a1y-dy, a1x-dx
				an2y, an2x := a2y+dy, a2x+dx

				if an1y >= 0 && an1y < h && an1x >= 0 && an1x < l && !antinodes[fmt.Sprintf("%d,%d", an1y, an1x)] {
					antinodes[fmt.Sprintf("%d,%d", an1y, an1x)] = true
				}

				if an2y >= 0 && an2y < h && an2x >= 0 && an2x < l && !antinodes[fmt.Sprintf("%d,%d", an2y, an2x)] {
					antinodes[fmt.Sprintf("%d,%d", an2y, an2x)] = true
				}
			}
		}
	}

	return antinodes
}

func findMultiAntinodes(antennas map[string][][]int, l int, h int) map[string]bool {
	antinodes := make(map[string]bool)

	for key := range antennas {
		for i := 0; i < len(antennas[key]); i++ {
			antinodes[fmt.Sprintf("%d,%d", antennas[key][i][0], antennas[key][i][1])] = true
		}
	}

	for key := range antennas {
		for i := 0; i < len(antennas[key]); i++ {
			for j := i + 1; j < len(antennas[key]); j++ {

				a1y, a1x := antennas[key][i][0], antennas[key][i][1]
				a2y, a2x := antennas[key][j][0], antennas[key][j][1]

				dy, dx := a2y-a1y, a2x-a1x

				for a1y-dy >= 0 && a1y-dy < h && a1x-dx >= 0 && a1x-dx < l {

					an1y, an1x := a1y-dy, a1x-dx

					if !antinodes[fmt.Sprintf("%d,%d", an1y, an1x)] {
						antinodes[fmt.Sprintf("%d,%d", an1y, an1x)] = true
					}

					a1y, a1x = an1y, an1x

				}

				for a2y+dy >= 0 && a2y+dy < h && a2x+dx >= 0 && a2x+dx < l {

					an2y, an2x := a2y+dy, a2x+dx

					if !antinodes[fmt.Sprintf("%d,%d", an2y, an2x)] {
						antinodes[fmt.Sprintf("%d,%d", an2y, an2x)] = true
					}

					a2y, a2x = an2y, an2x

				}
			}
		}
	}

	return antinodes
}

func main() {

	file, err := os.Open("cmd/aoc8/data.txt")
	check(err)

	matrix := readMatrix(file)

	antennas := mapAntennas(matrix)

	antinodes := findAntinodes(antennas, len(matrix[0]), len(matrix))

	multiAntinodes := findMultiAntinodes(antennas, len(matrix[0]), len(matrix))

	fmt.Println(len(antinodes), len(multiAntinodes))

}
