package Day07

import (
	"fmt"
	"math"
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

		for j := 0; j < int(math.Pow(2, float64(len(values))))-1; j++ {
			calc := 0

			for k := 0; k < len(values); k++ {
				currValue, err := strconv.Atoi(values[k])

				if err != nil {
					panic("aaaah")
				}

				if (j>>k)&1 == 1 && k != 0 {
					calc = calc * currValue
				} else if k == 0 && j == 0 {
					calc = currValue
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
		fmt.Printf("%v => donc %v boucles => de 0 a %v\n", values, int(math.Pow(3, float64(len(values)-1))), int(math.Pow(3, float64(len(values)-1)))-1)

		for j := 0; j < int(math.Pow(3, float64(len(values)-1))); j++ {
			calc := 0

			for k := 0; k < len(values); k++ {
				currValue, err := strconv.Atoi(values[k])

				if err != nil {
					panic("aaaah")
				}

				reverseJBase3 := strings.Split(strconv.FormatInt(int64(j), 3), "")
				for len(reverseJBase3) < len(values) {
					reverseJBase3 = append(reverseJBase3, "0")
				}

				if i == 6 {
					fmt.Println(reverseJBase3)
					fmt.Printf("j : %v  | k : %v | j en base3 : %v | bit k dans j : %v \n", j, k, strconv.FormatInt(int64(j), 3), reverseJBase3[k])
				}

				if k == 0 {
					calc = currValue
					fmt.Printf("%v\n", currValue)
				} else if reverseJBase3[k] == "0" {
					fmt.Printf("%v+%v\n", calc, currValue)
					calc += currValue
				} else if reverseJBase3[k] == "1" {
					fmt.Printf("%v*%v\n", calc, currValue)
					calc = calc * currValue
				} else if reverseJBase3[k] == "2" {
					value, err := strconv.Atoi(fmt.Sprintf("%v%v", calc, currValue))
					if err != nil {
						panic("aaaaaaa")
					}
					fmt.Printf("%v%v\n", calc, currValue)
					calc = value
				}
			}

			if i == 6 {
				fmt.Println(calc)
			}

			if resValue == calc {
				globalSum = globalSum + resValue
				if i != 6 {
					fmt.Println(calc)
				}
				fmt.Println("break !")
				break
			}
		}
	}

	return globalSum
}
func Resolve(data []string) [2]int {
	return [2]int{
		ResolvePart1(data),
		ResolvePart2(data),
	}
}
