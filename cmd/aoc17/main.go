package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type Registers struct {
	A int
	B int
	C int
}

func readInput(file *os.File) (Registers, []int) {
	registers := Registers{}
	commands := make([]int, 0)

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := strings.Split(scanner.Text(), ": ")
	a, err := strconv.Atoi(line[1])
	check(err)
	registers.A = a

	scanner.Scan()
	line = strings.Split(scanner.Text(), ": ")
	b, err := strconv.Atoi(line[1])
	check(err)
	registers.B = b

	scanner.Scan()
	line = strings.Split(scanner.Text(), ": ")
	c, err := strconv.Atoi(line[1])
	check(err)
	registers.C = c

	scanner.Scan()

	scanner.Scan()
	line = strings.Split(strings.Split(scanner.Text(), ": ")[1], ",")
	for _, command := range line {
		command, err := strconv.Atoi(command)
		check(err)
		commands = append(commands, command)
	}

	file.Close()

	return registers, commands

}

func combo(operand int, register Registers) int {

	switch operand {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 3
	case 4:
		return register.A
	case 5:
		return register.B
	case 6:
		return register.C
	default:
		err := errors.New("invalid operand " + strconv.Itoa(operand))
		check(err)
		return 0

	}
}

func adv(register Registers, operand int) Registers {
	// division
	register.A = register.A / int(math.Exp2(float64(combo(operand, register))))

	return register
}

func blx(register Registers, operand int) Registers {
	// bitwise xor
	register.B = register.B ^ operand

	return register
}

func bst(register Registers, operand int) Registers {
	// modulo 8
	register.B = combo(operand, register) % 8

	return register
}

func jnz(register Registers, operand int, set_point int) int {

	if register.A == 0 {
		return set_point + 2
	}

	return operand

}

func bxc(register Registers) Registers {
	// bitwise XOR
	register.B = register.B ^ register.C

	return register
}

func out(register Registers, operand int, output string) string {
	// output modulo 8
	output += strconv.Itoa(combo(operand, register)%8) + ","

	return output
}

func bdv(register Registers, operand int) Registers {
	// division
	register.B = register.A / int(math.Exp2(float64(combo(operand, register))))

	return register
}

func cdv(register Registers, operand int) Registers {
	// division
	register.C = register.A / int(math.Exp2(float64(combo(operand, register))))

	return register
}

func executeCommand(register Registers, operand int, command int, output string, set_point int) (Registers, int, string) {

	new_set_point := set_point + 2
	new_register := register
	new_output := output

	switch command {
	case 0:
		new_register = adv(register, operand)
	case 1:
		new_register = blx(register, operand)
	case 2:
		new_register = bst(register, operand)
	case 3:
		new_set_point = jnz(register, operand, set_point)
	case 4:
		new_register = bxc(register)
	case 5:
		new_output = out(register, operand, output)
	case 6:
		new_register = bdv(register, operand)
	case 7:
		new_register = cdv(register, operand)
	default:
		err := errors.New("invalid command")
		check(err)
	}

	return new_register, new_set_point, new_output

}

func runProgram(registers Registers, commands []int) string {

	output := ""
	sp := 0

	for sp < len(commands) {
		registers, sp, output = executeCommand(registers, commands[sp+1], commands[sp], output, sp)
	}

	return output

}

// brute force
// func findA(registers Registers, commands []int) int {

// 	i := 1

// 	for {

// 		Acandidate := 0 + i

// 		registers.A = Acandidate
// 		registers.B = 0
// 		registers.C = 0

// 		fmt.Println(registers)

// 		output := ""
// 		desired_output := []string{}

// 		for _, command := range commands {
// 			desired_output = append(desired_output, fmt.Sprintf("%d", command))

// 		}

// 		sp := 0

// 		for sp < len(commands) {
// 			registers, sp, output = executeCommand(registers, commands[sp+1], commands[sp], output, sp)

// 			if len(output) > 0 {
// 				if strings.Join(desired_output, ",") == output[:len(output)-1] {
// 					return Acandidate
// 				}
// 			}
// 		}

// 		i++

// 	}

// }

// reverse engineering
// 2,4,1,3,7,5,1,5,0,3,4,2,5,5,3,0
// b = a % 8
// b = b ^ 3
// c = a >> b
// b = b ^ 5
// a = a >> 3
// b = b ^ c
// # b % 8
// if a != 0
// 	jump 0

// at last iter we need to have
// 0 = a >> 3
// so a is between 0 and 7

// ((((5%8)^3)^5)^(5>>(5%8)^3))%8
// b % 8 = 0 -> a = 5

func reverseEngineer(commands []int, ans int) int {

	if len(commands) == 0 {
		return ans
	}

	//fmt.Println(commands, ans)

	for bb := 0; bb < 8; bb++ {
		a := (ans << 3) + bb

		b := a % 8
		b = b ^ 3
		c := a >> b
		b = b ^ 5
		b = b ^ c

		if b%8 == commands[len(commands)-1] {
			sub := reverseEngineer(commands[:len(commands)-1], a)
			if sub == -1 {
				continue
			}
			return sub
		}
	}

	return -1

}

func main() {

	file, err := os.Open("cmd/aoc17/data.txt")
	check(err)

	registers, commands := readInput(file)

	//fmt.Println(registers, commands)

	//findRegisters := Registers{0, 0, 0}
	//fmt.Println(findA(findRegisters, commands))

	fmt.Println(runProgram(registers, commands)[:len(runProgram(registers, commands))-1], reverseEngineer(commands, 0))

}
