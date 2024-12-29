package Day05

import (
	"slices"
	"strconv"
	"strings"
)

func isAtTheRightPlace(pos int, pageSet []string, rules map[string][]string) (bool, int) {
	mustBeBeforePages := rules[pageSet[pos]]
	betterIndex := -1
	ruleLength := len(mustBeBeforePages)
	count := 0

	if ruleLength > 0 {
		for i := 0; i < ruleLength; i++ {
			posRule := slices.Index(pageSet, mustBeBeforePages[i])
			if posRule == -1 || pos < posRule {
				count++
			} else {
				if betterIndex == -1 || posRule < betterIndex {
					betterIndex = posRule

					if betterIndex == 0 {
						break
					}
				}
			}
		}
	} else {
		count = -1
	}

	if count == -1 || count == ruleLength {
		return true, betterIndex
	}

	return false, betterIndex
}

func insert(a []string, index int, value string) []string {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	if index < 0 {
		index = 0
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

func solve(data []string, part int) int {
	index := -1
	sum := 0
	sumIncorrect := 0

	rules := make(map[string][]string)

	for i := 0; i < len(data); i++ {
		if data[i] == "" {
			index = i
			continue
		}

		if index == -1 {
			rule := strings.Split(data[i], "|")

			rules[rule[0]] = append(rules[rule[0]], rule[1])
		} else {
			pageSet := strings.Split(data[i], ",")
			count := 0
			for j := 0; j < len(pageSet); j++ {
				atRightPlace, betterPlace := isAtTheRightPlace(j, pageSet, rules)
				if atRightPlace {
					count++
				} else if part == 2 && betterPlace != -1 {
					testedSlice := slices.Clone(pageSet)
					testedModifiedSlice := append(testedSlice[:j], testedSlice[j+1:]...)
					pageSet = insert(testedModifiedSlice, betterPlace, pageSet[j])
				}
			}

			value, err := strconv.Atoi(pageSet[(len(pageSet)-1)/2])

			if err != nil {
				panic("aaaaa")
			}

			if count == len(pageSet) {
				sum += value
			} else {
				sumIncorrect += value
			}
		}
	}

	if part == 1 {
		return sum
	}

	return sumIncorrect
}

func ResolvePart1(data []string) int {
	sum := solve(data, 1)

	return sum
}

func ResolvePart2(data []string) int {
	sum := solve(data, 2)

	return sum
}

func Resolve(data []string) [2]any {
	return [2]any{
		ResolvePart1(data),
		ResolvePart2(data),
	}
}
