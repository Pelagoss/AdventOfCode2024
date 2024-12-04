package Day04

import (
	"adventOfCode/utils"
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
	if x <= (len(data[y]) - wordsLength + firstLetterPos) {
		rightDiscoverable = true
	}
	// Left discoverable
	if x >= (wordsLength - 1 - firstLetterPos) {
		leftDiscoverable = true
	}
	// Top discoverable
	if y >= (wordsLength - 1 - firstLetterPos) {
		topDiscoverable = true
	}
	// Bottom discoverable
	if y <= (len(data) - wordsLength + firstLetterPos) {
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
				if crossable {
					words["rt"] = words["rt"] + letter
				} else {
					words["rt"] = words["rt"] + letter
					if firstLetter != letter {
						words["lb"] = words["lb"] + letter
					}
				}
			}
		}

		if rightDiscoverable && bottomDiscoverable {
			if data[y+offsetAbs][x+offsetAbs] == letter {
				if crossable {
					words["rb"] = words["rb"] + letter
				} else {
					words["rb"] = words["rb"] + letter
					if firstLetter != letter {
						words["lt"] = words["lt"] + letter
					}
				}
			}
		}

		if leftDiscoverable && topDiscoverable {
			if data[y-offsetAbs][x-offsetAbs] == letter {
				if crossable {
					words["lt"] = words["lt"] + letter
				} else {
					words["lt"] = words["lt"] + letter
					if firstLetter != letter {
						words["rb"] = words["rb"] + letter
					}
				}
			}
		}

		if leftDiscoverable && bottomDiscoverable {
			if data[y+offsetAbs][x-offsetAbs] == letter {
				if crossable {
					words["lb"] = words["lb"] + letter
				} else {
					words["lb"] = words["lb"] + letter
					if firstLetter != letter {
						words["rt"] = words["rt"] + letter
					}
				}
			}
		}
	}

	for k := range words {
		if words[k] == wordToSearch {
			sum++
		}
	}

	if crossable {
		return sum
	}

	if sum == 4 {
		return 1
	}

	return 0
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
