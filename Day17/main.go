package Day17

import (
	"fmt"
	"strconv"
	"strings"
)

type Register struct {
	letter string
	value  int
}

func parseData(data []string) ([]int, []Register) {
	var (
		registers []Register
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
			fmt.Println(data[i])
			_, err := fmt.Sscanf(strings.ReplaceAll(data[i], ":", ""), "Register %s %d", &r.letter, &r.value)

			if err != nil {
				panic(err)
			}

			registers = append(registers, r)
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
func ResolvePart1(data []string) int {
	program, registers := parseData(data)
	fmt.Println(program, registers)
	return 0
}
func ResolvePart2(data []string) int {
	return 0
}
func Resolve(data []string) [2]int {
	return [2]int{
		ResolvePart1(data),
		ResolvePart2(data),
	}
}
