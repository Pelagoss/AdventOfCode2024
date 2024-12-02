package Day01

import (
	"adventOfCode/utils"
	"sort"
	"strconv"
)

func ResolvePart1(data []string) int {

	var list1 []int
	var list2 []int

	for i := 0; i < len(data); i++ {
		splitted := utils.RegSplit(data[i], "\\s+")

		if value, err := strconv.Atoi(splitted[0]); err == nil {
			list1 = append(list1, value)
		}

		if value, err := strconv.Atoi(splitted[1]); err == nil {
			list2 = append(list2, value)
		}
	}

	sort.Slice(list1, func(i, j int) bool {
		return list1[i] > list1[j]
	})

	sort.Slice(list2, func(i, j int) bool {
		return list2[i] > list2[j]
	})

	var result []int

	for i := 0; i < len(list1); i++ {
		result = append(result, utils.Abs(list1[i]-list2[i]))
	}

	var sum int

	for i := 0; i < len(result); i++ {
		sum += result[i]
	}

	return sum
}

func ResolvePart2(data []string) int {

	var list1 []int
	var list2 []int

	for i := 0; i < len(data); i++ {
		splitted := utils.RegSplit(data[i], "\\s+")

		if value, err := strconv.Atoi(splitted[0]); err == nil {
			list1 = append(list1, value)
		}

		if value, err := strconv.Atoi(splitted[1]); err == nil {
			list2 = append(list2, value)
		}
	}

	var result []int

	for i := 0; i < len(list1); i++ {
		nFounded := 0

		for j := 0; j < len(list2); j++ {
			if list1[i] == list2[j] {
				nFounded++
			}
		}

		result = append(result, list1[i]*nFounded)
	}

	var sum int

	for i := 0; i < len(result); i++ {
		sum += result[i]
	}

	return sum
}

func Resolve(data []string) [2]int {
	return [2]int{
		ResolvePart1(data),
		ResolvePart2(data),
	}
}
