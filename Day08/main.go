package Day08

import (
	"strings"
)

func ResolvePart1(data []string) int {
	var mapAntennas [][]string
	mapAntinodes := make(map[int]map[int]string)
	lastPos := make(map[string][][2]int)
	sum := 0

	for y := 0; y < len(data); y++ {
		line := strings.Split(data[y], "")
		mapAntennas = append(mapAntennas, line)

		for x := 0; x < len(line); x++ {
			if line[x] == "." {
				continue
			}
			prevCoordsWAntennas, _ := lastPos[line[x]]
			if len(prevCoordsWAntennas) != 0 {
				for i := 0; i < len(prevCoordsWAntennas); i++ {
					lastY := prevCoordsWAntennas[i][0]
					lastX := prevCoordsWAntennas[i][1]

					dY := y - lastY
					dX := x - lastX

					dYbis := lastY - y
					dXbis := lastX - x

					if y+dY >= 0 && y+dY < len(data) && x+dX >= 0 && x+dX < len(data[0]) {
						if _, ok := mapAntinodes[y+dY]; !ok {
							mapAntinodes[y+dY] = make(map[int]string)
						}

						if _, ok := mapAntinodes[y+dY][x+dX]; !ok {
							sum++
						}
						mapAntinodes[y+dY][x+dX] = "#"
					}

					if lastY+dYbis >= 0 && lastY+dYbis < len(data) && lastX+dXbis >= 0 && lastX+dXbis < len(data[0]) {
						if _, ok := mapAntinodes[lastY+dYbis]; !ok {
							mapAntinodes[lastY+dYbis] = make(map[int]string)
						}

						if _, ok := mapAntinodes[lastY+dYbis][lastX+dXbis]; !ok {
							sum++
						}
						mapAntinodes[lastY+dYbis][lastX+dXbis] = "#"
					}
				}
			}

			lastPos[line[x]] = append(lastPos[line[x]], [2]int{y, x})
		}
	}

	return sum
}
func ResolvePart2(data []string) int {
	var mapAntennas [][]string
	mapAntinodes := make(map[int]map[int]string)
	lastPos := make(map[string][][2]int)
	sum := 0

	for y := 0; y < len(data); y++ {
		line := strings.Split(data[y], "")
		mapAntennas = append(mapAntennas, line)

		for x := 0; x < len(line); x++ {
			if line[x] == "." {
				continue
			}

			if _, ok := mapAntinodes[y]; !ok {
				mapAntinodes[y] = make(map[int]string)
			}

			if _, ok := mapAntinodes[y][x]; !ok {
				sum++
			}

			mapAntinodes[y][x] = "#"

			prevCoordsWAntennas, _ := lastPos[line[x]]

			if len(prevCoordsWAntennas) != 0 {
				for i := 0; i < len(prevCoordsWAntennas); i++ {
					lastY := prevCoordsWAntennas[i][0]
					lastX := prevCoordsWAntennas[i][1]

					possible := true
					possibleBis := true

					for j := 1; possible || possibleBis; j++ {
						dY := (y - lastY) * j
						dX := (x - lastX) * j

						dYbis := (lastY - y) * j
						dXbis := (lastX - x) * j

						if y+dY >= 0 && y+dY < len(data) && x+dX >= 0 && x+dX < len(data[0]) {
							if _, ok := mapAntinodes[y+dY]; !ok {
								mapAntinodes[y+dY] = make(map[int]string)
							}

							if _, ok := mapAntinodes[y+dY][x+dX]; !ok {
								sum++
							}
							mapAntinodes[y+dY][x+dX] = "#"
						} else {
							possible = false
						}

						if lastY+dYbis >= 0 && lastY+dYbis < len(data) && lastX+dXbis >= 0 && lastX+dXbis < len(data[0]) {
							if _, ok := mapAntinodes[lastY+dYbis]; !ok {
								mapAntinodes[lastY+dYbis] = make(map[int]string)
							}

							if _, ok := mapAntinodes[lastY+dYbis][lastX+dXbis]; !ok {
								sum++
							}
							mapAntinodes[lastY+dYbis][lastX+dXbis] = "#"
						} else {
							possibleBis = false
						}
					}
				}
			}

			lastPos[line[x]] = append(lastPos[line[x]], [2]int{y, x})
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
