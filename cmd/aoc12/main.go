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

func mapRegions(matrix [][]rune) map[string][][]int {

	regions := make(map[string][][]int)
	mappedRegions := make(map[string]bool)

	for i, row := range matrix {
		for j, cell := range row {

			if _, ok := mappedRegions[fmt.Sprintf("%d,%d", i, j)]; !ok {
				regions[string(cell)] = append(regions[string(cell)], []int{i, j})
				mappedRegions[fmt.Sprintf("%d,%d", i, j)] = true
			}
		}
	}

	return regions

}

func findSeparatedRegions(regionName rune, region [][]int, matrix [][]rune) map[string][][]int {

	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	separatedRegions := make(map[string][][]int)
	visited := make(map[string]bool)

	var dfs func(int, int, int)
	regionCount := 0

	dfs = func(x, y, count int) {
		stack := [][]int{{x, y}}
		for len(stack) > 0 {
			cell := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			cx, cy := cell[0], cell[1]
			key := fmt.Sprintf("%d,%d", cx, cy)
			if visited[key] {
				continue
			}
			visited[key] = true
			separatedRegions[fmt.Sprintf("%s_%d", string(regionName), count)] = append(separatedRegions[fmt.Sprintf("%s_%d", string(regionName), count)], []int{cx, cy})
			for _, direction := range directions {
				nx, ny := cx+direction[0], cy+direction[1]
				if nx >= 0 && nx < len(matrix) && ny >= 0 && ny < len(matrix[0]) && matrix[nx][ny] == regionName && !visited[fmt.Sprintf("%d,%d", nx, ny)] {
					stack = append(stack, []int{nx, ny})
				}
			}
		}
	}

	for _, cell := range region {
		x, y := cell[0], cell[1]
		if !visited[fmt.Sprintf("%d,%d", x, y)] {
			regionCount++
			dfs(x, y, regionCount)
		}
	}

	return separatedRegions
}

func regionArea(region [][]int) int {

	return len(region)

}

func regionPerimeter(regionName rune, region [][]int, matrix [][]rune) int {

	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	perimeter := 0

	for _, cell := range region {
		for _, direction := range directions {
			dx := cell[0] + direction[0]
			dy := cell[1] + direction[1]

			if dx < 0 || dx >= len(matrix) || dy < 0 || dy >= len(matrix[0]) {
				perimeter++
				continue
			}

			if matrix[dx][dy] != regionName {
				perimeter++
			}

		}
	}

	return perimeter
}

func regionCorners(region [][]int) int {

	cornerCandidates := make(map[[2]float64]bool)

	for _, cell := range region {
		r, c := float64(cell[0]), float64(cell[1])
		corners := [][2]float64{
			{r - 0.5, c - 0.5},
			{r + 0.5, c - 0.5},
			{r + 0.5, c + 0.5},
			{r - 0.5, c + 0.5},
		}
		for _, corner := range corners {
			cornerCandidates[corner] = true
		}
	}

	corners := 0

	for corner := range cornerCandidates {
		cr, cc := corner[0], corner[1]

		config := []bool{
			contains(region, [2]int{int(cr - 0.5), int(cc - 0.5)}),
			contains(region, [2]int{int(cr + 0.5), int(cc - 0.5)}),
			contains(region, [2]int{int(cr + 0.5), int(cc + 0.5)}),
			contains(region, [2]int{int(cr - 0.5), int(cc + 0.5)}),
		}

		number := 0
		for _, v := range config {
			if v {
				number++
			}
		}

		if number == 1 {
			corners++
		} else if number == 2 {
			if (config[0] && !config[1] && config[2] && !config[3]) || (!config[0] && config[1] && !config[2] && config[3]) {
				corners += 2
			}
		} else if number == 3 {
			corners++
		}
	}

	return corners
}

func contains(region [][]int, cell [2]int) bool {
	for _, c := range region {
		if c[0] == cell[0] && c[1] == cell[1] {
			return true
		}
	}
	return false
}

func totPrice(regions map[string][][]int, matrix [][]rune) int {

	price := 0

	for key, region := range regions {
		price += regionArea(region) * regionPerimeter(rune(key[0]), region, matrix)
	}

	return price
}

func totBulkPrice(regions map[string][][]int) int {

	price := 0

	for _, region := range regions {
		price += regionArea(region) * regionCorners(region)
	}

	return price
}

func main() {

	file, err := os.Open("cmd/aoc12/data.txt")
	check(err)

	matrix := readMatrix(file)

	regions := mapRegions(matrix)

	separatedRegions := make(map[string][][]int)

	for key1, region := range regions {
		newRegions := findSeparatedRegions(rune(key1[0]), region, matrix)
		for key2, region := range newRegions {
			separatedRegions[key2] = region
		}
	}

	price := totPrice(separatedRegions, matrix)
	priceBulk := totBulkPrice(separatedRegions)
	fmt.Println(price, priceBulk)

}
