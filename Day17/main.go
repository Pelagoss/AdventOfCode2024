package Day17

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

type Register struct {
	letter string
	value  int
}

func pmod(x, d int) int {
	x = x % d
	if x >= 0 {
		return x
	}
	if d < 0 {
		return x - d
	}
	return x + d
}

func getComboOperand(index int, program []int, registers map[string]Register, literal bool) (int, any) {
	if index > len(program) {
		return 0, true // halts
	}

	if literal {
		return program[index], nil
	}

	switch program[index] {
	case 0, 1, 2, 3:
		return program[index], nil
	case 4:
		return registers["A"].value, nil
	case 5:
		return registers["B"].value, nil
	case 6:
		return registers["C"].value, nil
	}

	return 0, nil
}

func parseData(data []string) ([]int, map[string]Register) {
	var (
		registers = make(map[string]Register)
		program   []int
	)
	parseRegister := true
	for i := 0; i < len(data); i++ {
		if len(data[i]) == 0 {
			parseRegister = false
			continue
		}

		if parseRegister {
			r := Register{}
			_, err := fmt.Sscanf(strings.ReplaceAll(data[i], ":", ""), "Register %s %d", &r.letter, &r.value)

			if err != nil {
				panic(err)
			}

			registers[r.letter] = r
		} else {
			var programStr string
			_, err := fmt.Sscanf(data[i], "Program: %s", &programStr)

			if err != nil {
				panic(err)
			}

			programStrSplitted := strings.Split(programStr, ",")

			for j := 0; j < len(programStrSplitted); j++ {
				val, err := strconv.Atoi(programStrSplitted[j])
				if err != nil {
					panic(err)
				}
				program = append(program, val)
			}

		}
	}

	return program, registers
}
func ResolvePart1(data []string) string {
	program, registers := parseData(data)
	var out []string
	instructionPointer := 0
	needToBreak := false
	for needToBreak == false {
		needToIncrementTwice := true

		switch program[instructionPointer] {
		case 0: // adv
			comboOperand, err := getComboOperand(instructionPointer+1, program, registers, false)
			if err != nil {
				needToBreak = true
			}
			r := registers["A"]
			r.value = int(float64(r.value) / math.Pow(float64(2), float64(comboOperand)))
			registers[r.letter] = r
			break
		case 1: // bxl
			literalComboOperand, err := getComboOperand(instructionPointer+1, program, registers, true)
			if err != nil {
				needToBreak = true
			}
			r := registers["B"]
			r.value = r.value ^ literalComboOperand
			registers[r.letter] = r
			break
		case 2: // bst
			comboOperand, err := getComboOperand(instructionPointer+1, program, registers, false)
			if err != nil {
				needToBreak = true
			}
			r := registers["B"]
			r.value = pmod(comboOperand, 8)
			registers[r.letter] = r
			break
		case 3: // jnz
			if registers["A"].value == 0 {
				break
			}
			literalComboOperand, err := getComboOperand(instructionPointer+1, program, registers, true)
			if err != nil {
				needToBreak = true
			}
			instructionPointer = literalComboOperand
			needToIncrementTwice = false
			break
		case 4: // bxc
			r := registers["B"]
			r.value = r.value ^ registers["C"].value
			registers[r.letter] = r
			break
		case 5: // out
			comboOperand, err := getComboOperand(instructionPointer+1, program, registers, false)
			if err != nil {
				needToBreak = true
			}
			out = append(out, fmt.Sprintf("%d", pmod(comboOperand, 8)))
			break
		case 6: // bdv
			comboOperand, err := getComboOperand(instructionPointer+1, program, registers, false)
			if err != nil {
				needToBreak = true
			}
			r := registers["B"]
			r.value = int(float64(registers["A"].value) / math.Pow(float64(2), float64(comboOperand)))
			registers[r.letter] = r
			break
		case 7: // cdv
			comboOperand, err := getComboOperand(instructionPointer+1, program, registers, false)
			if err != nil {
				needToBreak = true
			}
			r := registers["C"]
			r.value = int(float64(registers["A"].value) / math.Pow(float64(2), float64(comboOperand)))
			registers[r.letter] = r
			break
		}

		if needToIncrementTwice {
			instructionPointer += 2
		}

		if instructionPointer >= len(program) {
			needToBreak = true
		}
		needToIncrementTwice = true
	}

	return strings.Join(out, ",")
}

func ResolvePart2(data []string) int {

	program, registers := parseData(data)

	initValue := 0

	ra := registers["A"]
	ra.value = initValue
	registers[ra.letter] = ra

	var out []int
	instructionPointer := 0
	needToBreak := false
	for i := len(program) - 1; i >= 0; i-- {
		initValue <<= 3

		for {
			needToBreak = false
			out = nil
			rA := registers["A"]
			instructionPointer = 0
			rA.value = initValue
			registers[rA.letter] = rA

			for needToBreak == false {
				needToIncrementTwice := true

				switch program[instructionPointer] {
				case 0: // adv
					comboOperand, err := getComboOperand(instructionPointer+1, program, registers, false)
					if err != nil {
						needToBreak = true
					}
					r := registers["A"]
					r.value = int(float64(r.value) / math.Pow(float64(2), float64(comboOperand)))
					registers[r.letter] = r
					break
				case 1: // bxl
					literalComboOperand, err := getComboOperand(instructionPointer+1, program, registers, true)
					if err != nil {
						needToBreak = true
					}
					r := registers["B"]
					r.value = r.value ^ literalComboOperand
					registers[r.letter] = r
					break
				case 2: // bst
					comboOperand, err := getComboOperand(instructionPointer+1, program, registers, false)
					if err != nil {
						needToBreak = true
					}
					r := registers["B"]
					r.value = pmod(comboOperand, 8)
					registers[r.letter] = r
					break
				case 3: // jnz
					if registers["A"].value == 0 {
						break
					}
					literalComboOperand, err := getComboOperand(instructionPointer+1, program, registers, true)
					if err != nil {
						needToBreak = true
					}
					instructionPointer = literalComboOperand
					needToIncrementTwice = false
					break
				case 4: // bxc
					r := registers["B"]
					r.value = r.value ^ registers["C"].value
					registers[r.letter] = r
					break
				case 5: // out
					comboOperand, err := getComboOperand(instructionPointer+1, program, registers, false)
					if err != nil {
						needToBreak = true
					}
					out = append(out, pmod(comboOperand, 8))
					break
				case 6: // bdv
					comboOperand, err := getComboOperand(instructionPointer+1, program, registers, false)
					if err != nil {
						needToBreak = true
					}
					r := registers["B"]
					r.value = int(float64(registers["A"].value) / math.Pow(float64(2), float64(comboOperand)))
					registers[r.letter] = r
					break
				case 7: // cdv
					comboOperand, err := getComboOperand(instructionPointer+1, program, registers, false)
					if err != nil {
						needToBreak = true
					}
					r := registers["C"]
					r.value = int(float64(registers["A"].value) / math.Pow(float64(2), float64(comboOperand)))
					registers[r.letter] = r
					break
				}

				if needToIncrementTwice {
					instructionPointer += 2
				}

				if instructionPointer >= len(program) {
					needToBreak = true
				}
				needToIncrementTwice = true
			}
			if !slices.Equal(out, program[i:]) {
				initValue++
			} else {
				break
			}
		}
	}

	return initValue
}
func Resolve(data []string) [2]any {
	return [2]any{
		ResolvePart1(data),
		ResolvePart2(data),
	}
}
