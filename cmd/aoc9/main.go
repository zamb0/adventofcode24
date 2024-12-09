package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fileRead(file []byte) []int {

	memory := []int{}
	fid := 0

	for i, char := range file {
		x, err := strconv.Atoi(string(char))
		check(err)

		if i%2 == 0 {
			for j := 0; j < x; j++ {
				memory = append(memory, fid)
			}
			fid++
		} else {
			for j := 0; j < x; j++ {
				memory = append(memory, -1)
			}
		}
	}

	return memory

}

func fileRead2(file []byte) (map[int][]int, [][]int, error) {

	files := map[int][]int{}
	free := [][]int{}
	fid := 0
	pos := 0

	for i, char := range file {
		x, err := strconv.Atoi(string(char))
		check(err)

		if i%2 == 0 {
			if x == 0 {
				return nil, nil, errors.New("find an 'empty' file (0), this is not allowed")
			}
			files[fid] = []int{pos, x}
			fid++
		} else {
			if x != 0 {
				free = append(free, []int{pos, x})
			}
		}
		pos += x
	}

	return files, free, nil
}

func fillMemory(memory []int) []int {

	free := free(memory)

	for _, i := range free {
		for len(memory) > 0 && memory[len(memory)-1] == -1 {
			memory = memory[:len(memory)-1]
		}
		if len(memory) <= i {
			break
		}
		memory[i] = memory[len(memory)-1]
		memory = memory[:len(memory)-1]
	}

	return memory
}

func fillMemory2(files map[int][]int, free [][]int) map[int][]int {

	fid := len(files)

	for fid > 0 {
		fid--
		filePos, fileSize := files[fid][0], files[fid][1]

		for i := 0; i < len(free); i++ {
			start, length := free[i][0], free[i][1]

			if start >= filePos {
				free = free[:i]
				break
			}

			if fileSize <= length {
				files[fid] = []int{start, fileSize}
				if fileSize == length {
					free = append(free[:i], free[i+1:]...)
				} else {
					free[i] = []int{start + fileSize, length - fileSize}
				}
				break
			}
		}
	}

	return files
}

func free(memory []int) []int {

	blanks := []int{}

	for i, val := range memory {
		if val == -1 {
			blanks = append(blanks, i)
		}
	}

	return blanks
}

func checksum(memory []int) int {

	sum := 0

	for i, x := range memory {
		sum += i * x
	}

	return sum
}

func checksum2(memory map[int][]int) int {

	sum := 0

	for i, x := range memory {
		p, s := x[0], x[1]
		for j := p; j < p+s; j++ {
			sum += i * j
		}
	}

	return sum
}

func main() {

	file, err := os.ReadFile("cmd/aoc9/data.txt")
	check(err)

	memory := fileRead(file)

	files, free, err := fileRead2(file)
	check(err)

	filledMemory := fillMemory(memory)

	filledMemory2 := fillMemory2(files, free)

	sum1 := checksum(filledMemory)

	sum2 := checksum2(filledMemory2)

	fmt.Println(sum1, sum2)
}
