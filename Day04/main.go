package Day04

import (
	"adventOfCode/utils"
	"fmt"
	"strings"
)

func countWordsFromPosition(wordToSearch string, firstLetter string, data [][]string, x int, y int, crossable bool) int {
	sum := 0
	wordsLength := len(wordToSearch)
	rightDiscoverable := false
	leftDiscoverable := false
	topDiscoverable := false
	bottomDiscoverable := false

	firstLetterPos := strings.Index(wordToSearch, firstLetter)

	// Right discoverable
	if x <= len(data[y])-wordsLength {
		rightDiscoverable = true
	}
	// Left discoverable
	if x >= (wordsLength - 1) {
		leftDiscoverable = true
	}
	// Top discoverable
	if y >= (wordsLength - 1) {
		topDiscoverable = true
	}
	// Bottom discoverable
	if y <= (len(data) - wordsLength) {
		bottomDiscoverable = true
	}

	words := make(map[string]string)

	for i := 0; i < wordsLength; i++ {
		letter := string(wordToSearch[i])
		offset := i - firstLetterPos
		offsetAbs := utils.Abs(offset)

		if rightDiscoverable && crossable {
			if data[y][x+offsetAbs] == letter {
				words["right"] = words["right"] + letter
			}
		}

		if leftDiscoverable && crossable {
			if data[y][x-offsetAbs] == letter {
				words["left"] = words["left"] + letter
			}
		}

		if topDiscoverable && crossable {
			if data[y-offsetAbs][x] == letter {
				words["top"] = words["top"] + letter
			}
		}

		if bottomDiscoverable && crossable {
			if data[y+offsetAbs][x] == letter {
				words["bottom"] = words["bottom"] + letter
			}
		}

		if rightDiscoverable && topDiscoverable {
			if data[y-offsetAbs][x+offsetAbs] == letter {
				//if offset >= 0 {
				words["rt"] = words["rt"] + letter
				//} else {
				//	words["rt"] = letter + words["rt"]
				//	words["lb"] = words["lb"] + letter
				//}
			}
		}

		if rightDiscoverable && bottomDiscoverable {
			if data[y+offsetAbs][x+offsetAbs] == letter {
				//if offset >= 0 {
				words["rb"] = words["rb"] + letter
				//} else {
				//	words["rb"] = letter + words["rb"]
				//	words["lt"] = words["lt"] + letter
				//}
			}
		}

		if leftDiscoverable && topDiscoverable {
			if data[y-offsetAbs][x-offsetAbs] == letter {
				//if offset >= 0 {
				words["lt"] = words["lt"] + letter
				//} else {
				//	words["lt"] = letter + words["lt"]
				//	words["rb"] = words["rb"] + letter
				//}
			}
		}

		if leftDiscoverable && bottomDiscoverable {
			if data[y+offsetAbs][x-offsetAbs] == letter {
				//if offset >= 0 {
				words["lb"] = words["lb"] + letter
				//} else {
				//	words["rt"] = words["rt"] + letter
				//	words["lb"] = letter + words["lb"]
				//}
			}
		}
	}

	if y == 1 && crossable == false {
		fmt.Println(words)
	}

	for k := range words {
		if words[k] == wordToSearch {
			sum++
		}
	}

	return sum
}

func ResolvePart1(data []string, wordToSearch string, indexFirstLetter int, crossable bool) int {
	sum := 0
	var matrix [][]string

	for i := 0; i < len(data); i++ {
		matrix = append(matrix, strings.Split(data[i], ""))
	}

	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[y]); x++ {
			if matrix[y][x] == string(wordToSearch[indexFirstLetter]) {
				sum += countWordsFromPosition(wordToSearch, string(wordToSearch[indexFirstLetter]), matrix, x, y, crossable)
				if y == 1 && !crossable {
					panic("aa")
				}
			}
		}
	}

	return sum
}

func ResolvePart2(data []string) int {
	return ResolvePart1(data, "MAS", 1, false)
}

func Resolve(data []string) [2]int {
	return [2]int{
		ResolvePart1(data, "XMAS", 0, true),
		ResolvePart2(data),
	}
}
