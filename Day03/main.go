package Day03

import (
	"adventOfCode/utils"
	"strconv"
	"strings"
)

func ResolvePart1(data []string) int {
	sum := 0

	for i := 0; i < len(data); i++ {
		splitted := utils.RegMatch(data[i], "mul\\((\\d+),(\\d+)\\)")

		for j := 0; j < len(splitted); j++ {
			value1, err := strconv.Atoi(splitted[j][1])
			if err != nil {
				panic(err)
			}

			value2, err := strconv.Atoi(splitted[j][2])
			if err != nil {
				panic(err)
			}

			sum += value1 * value2
		}
	}

	return sum
}
func ResolvePart2(data []string) int {
	var list []string

	enabled := true

	for i := 0; i < len(data); i++ {
		strLeft := strings.Clone(data[i])

		// strip parts between don't() & do()
		for len(strLeft) > 0 {
			reg := "do()"

			if enabled {
				reg = "don't()"
			}

			before, after, found := strings.Cut(strLeft, reg)

			if found {
				if enabled {
					list = append(list, before)
					enabled = false
				} else {
					enabled = true
				}

				strLeft = after
			} else {
				if enabled {
					list = append(list, strLeft)
				}
				strLeft = ""
			}
		}
	}

	return ResolvePart1(list)
}
func Resolve(data []string) [2]any {
	return [2]any{
		ResolvePart1(data),
		ResolvePart2(data),
	}
}
