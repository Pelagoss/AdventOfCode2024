package Day07

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func ResolvePart1(data []string) int {
	globalSum := 0

	for i := 0; i < len(data); i++ {
		split := strings.Split(data[i], ":")

		resValue, err := strconv.Atoi(split[0])

		if err != nil {
			panic("aaaah")
		}

		values := strings.Split(strings.Trim(split[1], " "), " ")

		for j := 0; j < int(math.Pow(2, float64(len(values))-1)); j++ {
			firstValue, err := strconv.Atoi(values[0])

			if err != nil {
				panic("aaaah")
			}

			calc := firstValue

			for k := 1; k < len(values); k++ {
				currValue, err := strconv.Atoi(values[k])

				if err != nil {
					panic("aaaah")
				}

				if (j>>(k-1))&1 == 1 {
					calc = calc * currValue
				} else {
					calc += currValue
				}
			}

			if resValue == calc {
				globalSum = globalSum + resValue
				break
			}
		}
	}

	return globalSum
}
func ResolvePart2(data []string) int {
	globalSum := 0

	for i := 0; i < len(data); i++ {
		split := strings.Split(data[i], ":")

		resValue, err := strconv.Atoi(split[0])

		if err != nil {
			panic("aaaah")
		}

		values := strings.Split(strings.Trim(split[1], " "), " ")

		for j := 0; j < int(math.Pow(3, float64(len(values)-1))); j++ {
			firstValue, err := strconv.Atoi(values[0])

			if err != nil {
				panic("aaaah")
			}

			calc := firstValue

			for k := 1; k < len(values); k++ {
				currValue, err := strconv.Atoi(values[k])

				if err != nil {
					panic("aaaah")
				}

				reverseJBase3 := strings.Split(strconv.FormatInt(int64(j), 3), "")

				slices.Reverse(reverseJBase3)

				for len(reverseJBase3) < len(values)-1 {
					reverseJBase3 = append(reverseJBase3, "0")
				}

				slices.Reverse(reverseJBase3)

				if reverseJBase3[k-1] == "0" {
					calc += currValue
				} else if reverseJBase3[k-1] == "1" {
					calc = calc * currValue
				} else if reverseJBase3[k-1] == "2" {
					value, err := strconv.Atoi(fmt.Sprintf("%v%v", calc, currValue))
					if err != nil {
						panic("aaaaaaa")
					}
					calc = value
				}
			}

			if resValue == calc {
				globalSum = globalSum + resValue
				break
			}
		}
	}

	return globalSum
}
func Resolve(data []string) [2]any {
	return [2]any{
		ResolvePart1(data),
		ResolvePart2(data),
	}
}
