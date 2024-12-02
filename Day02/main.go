package Day02

import (
	"adventOfCode/utils"
	"slices"
	"strconv"
)

func check(report []int) bool {
	splittedReversed := slices.Clone(report)
	slices.Reverse(splittedReversed)

	if slices.IsSorted(report) || slices.IsSorted(splittedReversed) {
		uniqueMap := make(map[int]bool)

		for i := 0; i < len(report); i++ {
			uniqueMap[report[i]] = true
		}

		if len(uniqueMap) == len(report) {
			var everySlice []bool

			for i := 0; i < len(report); i++ {
				if i < len(report)-1 {
					everySlice = append(everySlice, utils.Abs(report[i]-report[i+1]) <= 3)
				} else {
					everySlice = append(everySlice, true)
				}
			}
			counter := 0

			for i := 0; i < len(everySlice); i++ {
				if everySlice[i] {
					counter++
				}
			}

			if counter == len(everySlice) {
				return true
			}
		}
	}

	return false
}

func ResolvePart1(data []string) int {
	var list []string

	for i := 0; i < len(data); i++ {
		splitted := utils.RegSplit(data[i], "\\s+")

		var sliceAndInt []int

		for i := 0; i < len(splitted); i++ {
			if value, err := strconv.Atoi(splitted[i]); err == nil {
				sliceAndInt = append(sliceAndInt, value)
			} else {
				panic(err)
			}
		}

		if check(sliceAndInt) {
			list = append(list, data[i])
		}
	}

	return len(list)
}

func ResolvePart2(data []string) int {
	var badReports [][]int

	for i := 0; i < len(data); i++ {
		splitted := utils.RegSplit(data[i], "\\s+")

		var sliceAndInt []int

		for i := 0; i < len(splitted); i++ {
			if value, err := strconv.Atoi(splitted[i]); err == nil {
				sliceAndInt = append(sliceAndInt, value)
			} else {
				panic(err)
			}
		}

		if !check(sliceAndInt) {
			badReports = append(badReports, sliceAndInt)
		}
	}

	firstResult := ResolvePart1(data)

	var list [][]int

	for i := 0; i < len(badReports); i++ {
		for j := 0; j < len(badReports[i]); j++ {
			testedSlice := slices.Clone(badReports[i])
			testedModifiedSlice := append(testedSlice[:j], testedSlice[j+1:]...)

			if check(testedModifiedSlice) {
				list = append(list, badReports[i])
				break
			}
		}
	}

	return firstResult + len(list)
}

func Resolve(data []string) [2]int {
	return [2]int{
		ResolvePart1(data),
		ResolvePart2(data),
	}
}
