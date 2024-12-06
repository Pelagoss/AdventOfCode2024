package Day06

import (
	"adventOfCode/utils"
	"strings"
)

func getGuardDirection(guard string) [2]int {
	switch guard {
	case "^":
		return [2]int{-1, 0}
	case ">":
		return [2]int{0, 1}
	case "<":
		return [2]int{0, -1}
	case "v":
		return [2]int{1, 0}
	}

	return [2]int{0, 0}
}

func initMap(data *[]string, mapArea *map[int][]string, caseVisited *map[int][]bool, guard *string, guardPos *[2]int) {
	for y := 0; y < len(*data); y++ {
		row := strings.Split((*data)[y], "")

		(*mapArea)[y] = append((*mapArea)[y], row...)

		for x := 0; x < len(row); x++ {
			regMatch := utils.RegMatch(row[x], "[\\^><v]")
			(*caseVisited)[y] = append((*caseVisited)[y], false)
			if len(regMatch) > 0 {
				*guard = regMatch[0][0]
				*guardPos = [2]int{y, x}
				(*caseVisited)[y][x] = true
				(*mapArea)[y][x] = "."
			}
		}
	}
}

func moveGuardianOnce(mapArea *map[int][]string, caseVisited *map[int][]bool, guard *string, guardPos *[2]int, guardDirection *[2]int) (bool, any) {
	futureY := guardPos[0] + guardDirection[0]
	futureX := guardPos[1] + guardDirection[1]

	var futurePos string
	if futureX >= 0 && futureX < len((*mapArea)[futureY]) {
		futurePos = (*mapArea)[futureY][futureX]
	}

	switch futurePos {
	case ".", "":
		(*guardPos)[0] = futureY
		(*guardPos)[1] = futureX

		if futurePos == "." {
			alreadyVisited := (*caseVisited)[futureY][futureX]
			(*caseVisited)[futureY][futureX] = true

			return alreadyVisited, nil
		}
		break
	case "#":
		switch *guard {
		case "^":
			*guard = ">"
			break
		case ">":
			*guard = "v"
			break
		case "v":
			*guard = "<"
			break
		case "<":
			*guard = "^"
			break
		}
	}

	return false, true
}

func ResolvePart1(data []string) int {
	isInArea := true
	var guardDirection [2]int
	var guard string
	var guardPos [2]int
	mapArea := make(map[int][]string)
	caseVisited := make(map[int][]bool)

	initMap(&data, &mapArea, &caseVisited, &guard, &guardPos)

	for isInArea == true {
		guardDirection = getGuardDirection(guard)

		moveGuardianOnce(&mapArea, &caseVisited, &guard, &guardPos, &guardDirection)

		if !(guardPos[0] >= 0 && guardPos[0] < len(data) && guardPos[1] >= 0 && guardPos[1] < len(data[0])) {
			isInArea = false
		}
	}

	sum := 0
	for k := range caseVisited {
		for i := 0; i < len(caseVisited[k]); i++ {
			if caseVisited[k][i] {
				sum++
			}
		}
	}

	return sum
}
func ResolvePart2(data []string) int {
	var guardDirection [2]int
	var guardInitial string
	var guardInitialPos [2]int
	mapArea := make(map[int][]string)
	caseVisited := make(map[int][]bool)

	initMap(&data, &mapArea, &caseVisited, &guardInitial, &guardInitialPos)

	sum := 0

	// On peut remplacer cette double boucle for par une boucle qui placerait des # sur tous les points successifs du chemin trouvé dans la partie 1, mais la flemme là
	for y := 0; y < len(mapArea); y++ {
		for x := 0; x < len(mapArea[y]); x++ {
			if "#" == mapArea[y][x] {
				continue
			}

			mapAreaTested := make(map[int][]string)
			for k, v := range mapArea {
				for i := 0; i < len(v); i++ {
					mapAreaTested[k] = append(mapAreaTested[k], v[i])
				}
			}

			caseVisitedTested := make(map[int][]bool)
			for k, v := range caseVisited {
				for i := 0; i < len(v); i++ {
					caseVisitedTested[k] = append(caseVisitedTested[k], v[i])
				}
			}

			mapAreaTested[y][x] = "#"

			caseVisitedSum := 0
			isInArea := true

			guard := guardInitial
			guardPos := guardInitialPos

			count := 0

			for isInArea == true && caseVisitedSum < len(caseVisitedTested) {
				guardDirection = getGuardDirection(guard)

				isAlreadyVisited, err := moveGuardianOnce(&mapAreaTested, &caseVisitedTested, &guard, &guardPos, &guardDirection)
				if err == nil {
					if isAlreadyVisited {
						caseVisitedSum++
					} else {
						caseVisitedSum = 0
					}
				}
				if !(guardPos[0] >= 0 && guardPos[0] < len(data) && guardPos[1] >= 0 && guardPos[1] < len(data[0])) {
					isInArea = false
				}

				count++
			}

			if isInArea {
				sum++
			}
		}
	}

	return sum
}
func Resolve(data []string) [2]int {
	return [2]int{
		ResolvePart1(data),
		ResolvePart2(data),
	}
}
