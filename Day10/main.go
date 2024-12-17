package Day10

import (
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	X, Y        int
	accessibles []Point
}

func followTrail(x int, y int, level int, data []string, accessibleEnd *map[string]int) {
	currentPoint := Point{X: x, Y: y}

	if level == 9 {
		(*accessibleEnd)[fmt.Sprintf("%v_%v", currentPoint.X, currentPoint.Y)]++
	}

	right := currentPoint.X < len(data)-1
	left := currentPoint.X > 0
	top := currentPoint.Y > 0
	bottom := currentPoint.Y < len(data)-1

	var neighbor []Point

	if right {
		neighbor = append(neighbor, Point{X: currentPoint.X + 1, Y: currentPoint.Y})
	}
	if left {
		neighbor = append(neighbor, Point{X: currentPoint.X - 1, Y: currentPoint.Y})
	}
	if top {
		neighbor = append(neighbor, Point{X: currentPoint.X, Y: currentPoint.Y - 1})
	}
	if bottom {
		neighbor = append(neighbor, Point{X: currentPoint.X, Y: currentPoint.Y + 1})
	}

	for i := 0; i < len(neighbor); i++ {
		abscissesNeighbor := strings.Split(data[neighbor[i].Y], "")

		levelNeighbor, err := strconv.Atoi(abscissesNeighbor[neighbor[i].X])

		if err != nil {
			panic(err)
		}

		if levelNeighbor == level+1 {
			currentPoint.accessibles = append(currentPoint.accessibles, neighbor[i])
		}
	}

	if len(currentPoint.accessibles) > 0 {
		for i := 0; i < len(currentPoint.accessibles); i++ {
			followTrail(currentPoint.accessibles[i].X, currentPoint.accessibles[i].Y, level+1, data, accessibleEnd)
		}
	}
}

func solve(data []string, part int) int {
	sum := 0
	for y := 0; y < len(data); y++ {
		abscisses := strings.Split(data[y], "")

		for x := 0; x < len(abscisses); x++ {
			if abscisses[x] != "0" {
				continue
			}

			level, err := strconv.Atoi(abscisses[x])

			if err != nil {
				panic(err)
			}

			accessibleEnd := make(map[string]int)

			followTrail(x, y, level, data, &accessibleEnd)

			if part == 1 {
				sum += len(accessibleEnd)
			} else if part == 2 {
				for k, _ := range accessibleEnd {
					sum += accessibleEnd[k]
				}
			}
		}
	}

	return sum
}
func Resolve(data []string) [2]int {
	return [2]int{
		solve(data, 1),
		solve(data, 2),
	}
}
