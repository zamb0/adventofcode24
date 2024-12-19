package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type Point struct {
	x int
	y int
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readInput(file *os.File) []Point {

	points := make([]Point, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		var x, y int
		_, err := fmt.Sscanf(line, "%d,%d", &x, &y)
		check(err)
		points = append(points, Point{x, y})
	}

	return points

}

func simulation(points []Point, matrix [71][71]int, sim_length int) [71][71]int {

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			for _, point := range points[:sim_length] {
				if point.x == j && point.y == i {
					matrix[i][j] = -1
				}
			}
		}
	}

	return matrix

}

func findShortestPath(matrix [71][71]int, start Point, end Point) []Point {
	path := make([]Point, 0)

	if matrix[start.y][start.x] == -1 || matrix[end.y][end.x] == -1 {
		return path
	}

	visited := make(map[Point]bool)
	distances := make(map[Point]int)
	previous := make(map[Point]Point)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			distances[Point{j, i}] = 1000000
		}
	}

	distances[start] = 0

	for len(visited) < len(matrix)*len(matrix[0]) {
		min_distance := 1000000
		var current Point
		found := false

		for i := 0; i < len(matrix); i++ {
			for j := 0; j < len(matrix[0]); j++ {
				if distances[Point{j, i}] < min_distance && !visited[Point{j, i}] {
					min_distance = distances[Point{j, i}]
					current = Point{j, i}
					found = true
				}
			}
		}

		if !found {
			break
		}

		visited[current] = true

		if current == end {
			break
		}

		neighbours := make([]Point, 0)

		if current.x > 0 {
			neighbours = append(neighbours, Point{current.x - 1, current.y})
		}

		if current.x < len(matrix[0])-1 {
			neighbours = append(neighbours, Point{current.x + 1, current.y})
		}

		if current.y > 0 {
			neighbours = append(neighbours, Point{current.x, current.y - 1})
		}

		if current.y < len(matrix)-1 {
			neighbours = append(neighbours, Point{current.x, current.y + 1})
		}

		for _, neighbour := range neighbours {
			if !visited[neighbour] && matrix[neighbour.y][neighbour.x] != -1 {
				alt := distances[current] + 1
				if alt < distances[neighbour] {
					distances[neighbour] = alt
					previous[neighbour] = current
				}
			}
		}

	}

	current := end

	for current != start {
		if _, ok := previous[current]; !ok {
			return []Point{}
		}
		path = append(path, current)
		current = previous[current]
	}

	return path
}

func simulateTillNoSolution(points []Point, matrix [71][71]int) Point {

	path := findShortestPath(matrix, Point{0, 0}, Point{70, 70})
	i := 1
	for len(path) > 0 {
		//fmt.Printf("%d: %d,%d\n", 1024+i, points[1024+i].x, points[1024+i].y)
		matrix[points[1024+i].y][points[1024+i].x] = -1
		if slices.Contains(path, points[1024+i]) {
			path = findShortestPath(matrix, Point{0, 0}, Point{70, 70})
		}
		//fmt.Println(len(path))
		i++
	}

	return points[1024+i-1]

}

func main() {

	file, err := os.Open("cmd/aoc18/data.txt")
	check(err)

	points := readInput(file)

	matrix := [71][71]int{{0}}
	matrix = simulation(points, matrix, 1024)

	path := findShortestPath(matrix, Point{0, 0}, Point{70, 70})

	last_point := simulateTillNoSolution(points, matrix)
	last_point_s := fmt.Sprintf("%d,%d", last_point.x, last_point.y)

	fmt.Println(len(path), last_point_s)
}
