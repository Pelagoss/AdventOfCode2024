package Day15

import (
	"slices"
)

type Point struct {
	x, y     int
	typePt   string
	letter   string
	brotherX int
}

type VectorMove struct {
	x, y int
}

func Move(movable *Point, move VectorMove, carte *[][]Point) bool {
	initX, initY := movable.x, movable.y
	initBrotherX := movable.brotherX
	nextX := (*movable).x + move.x
	nextY := (*movable).y + move.y

	moveIsPossible := false

	if nextX < len((*carte)[0]) && nextY < len(*carte) {
		nextCase := (*carte)[nextY][nextX]
		switch nextCase.typePt {
		case "Box":
			moveIsPossible = Move(&nextCase, move, carte)
			break
		case "EmptySpace":
			moveIsPossible = true
			break
		}

		if moveIsPossible {
			(*movable).x, (*movable).y, (*movable).brotherX = nextX, nextY, initBrotherX+move.x
			(*carte)[nextY][nextX] = *movable
			(*carte)[initY][initX] = Point{typePt: "EmptySpace", x: initX, y: initY, letter: "."}
		}
	}

	return moveIsPossible
}

func CanHeMove(movable *Point, move VectorMove, carte *[][]Point) bool {
	nextX := (*movable).x + move.x
	nextY := (*movable).y + move.y

	if nextX < len((*carte)[0]) && nextY < len(*carte) {
		var BoxesToMove []Point
		BoxMap := make(map[Point]bool)
		nextCase := (*carte)[nextY][nextX]
		if nextCase.typePt == "EmptySpace" {
			return true
		} else if nextCase.typePt == "Wall" {
			return false
		}

		BoxesToMove = append(BoxesToMove, nextCase, (*carte)[nextY][nextCase.brotherX])
		BoxMap[nextCase] = true
		BoxMap[(*carte)[nextY][nextCase.brotherX]] = true

		for i := 0; i < len(BoxesToMove); i++ {
			box := BoxesToMove[i]
			box2 := (*carte)[box.y][box.brotherX]

			if _, ok := BoxMap[box2]; !ok {
				BoxesToMove = append(BoxesToMove, box2)
				BoxMap[box2] = true
			}

			next := (*carte)[box.y+move.y][box.x+move.x]
			next2 := (*carte)[box2.y+move.y][box2.x+move.x]

			if next.typePt == "Box" {
				if _, ok := BoxMap[next]; !ok {
					BoxesToMove = append(BoxesToMove, next)
					BoxMap[next] = true
				}
			} else if next.typePt == "Wall" {
				return false
			}

			if next2.typePt == "Box" {
				if _, ok := BoxMap[next2]; !ok {
					BoxesToMove = append(BoxesToMove, next2)
					BoxMap[next2] = true
				}
			} else if next2.typePt == "Wall" {
				return false
			}
		}

		slices.Reverse(BoxesToMove)

		for i := 0; i < len(BoxesToMove); i++ {
			box := BoxesToMove[i]
			initX, initY := box.x, box.y
			nextX, nextY := initX+move.x, initY+move.y

			toMove := (*carte)[initY][initX]

			toMove.y, toMove.x = nextY, nextX
			(*carte)[nextY][nextX] = toMove
			(*carte)[initY][initX] = Point{typePt: "EmptySpace", x: initX, y: initY, letter: "."}
		}
	}

	return true
}
func Move2(movable *Point, move VectorMove, carte *[][]Point) bool {
	initX, initY := movable.x, movable.y
	nextX := (*movable).x + move.x
	nextY := (*movable).y + move.y

	if nextX < len((*carte)[0]) && nextY < len(*carte) {
		moveIsPossible := CanHeMove(movable, move, carte)
		if moveIsPossible {
			(*movable).x, (*movable).y = nextX, nextY
			(*carte)[nextY][nextX] = *movable
			(*carte)[initY][initX] = Point{typePt: "EmptySpace", x: initX, y: initY, letter: "."}
		}
	}

	return false
}

func ResolvePart1(data []string) int {
	lineIsMap := true
	robot := Point{typePt: "Robot", letter: "@"}
	var carte [][]Point

	for y := 0; y < len(data); y++ {
		if len(data[y]) == 0 {
			lineIsMap = false
			continue
		}

		var object Point

		// Map
		// Setup la map
		var line []Point
		for x := 0; x < len(data[y]); x++ {
			if lineIsMap == true {
				switch string(data[y][x]) {
				case "@":
					object = robot
					// Robot
					robot.x = x
					robot.y = y
					break
				case "O":
					object = Point{typePt: "Box", x: x, y: y, letter: "O"}
					// Box
					break
				case "#":
					object = Point{typePt: "Wall", x: x, y: y, letter: "#"}
					// Wall
					break
				case ".":
					object = Point{typePt: "EmptySpace", x: x, y: y, letter: "."}
					// EmptySpace
					break
				}

				line = append(line, object)
			} else {
				vectorMove := VectorMove{}
				// Directions
				switch string(data[y][x]) {
				case ">":
					vectorMove.x = 1
					break
				case "^":
					vectorMove.y = -1
					break
				case "v":
					vectorMove.y = 1
					break
				case "<":
					vectorMove.x = -1
					break
				}

				Move(&robot, vectorMove, &carte)
			}
		}
		if lineIsMap == true {
			carte = append(carte, line)
		}
	}

	sum := 0
	for y := 0; y < len(carte); y++ {
		for x := 0; x < len(carte); x++ {
			if carte[y][x].letter == "O" {
				sum += 100*y + x
			}
		}
	}

	return sum
}
func ResolvePart2(data []string) int {
	lineIsMap := true
	robot := Point{typePt: "Robot", letter: "@"}
	var carte [][]Point

	for y := 0; y < len(data); y++ {
		if len(data[y]) == 0 {
			lineIsMap = false
			continue
		}

		var object Point
		var object2 Point

		// Map
		// Setup la map
		var line []Point
		for x := 0; x < len(data[y]); x++ {
			realX := x + x*1
			if lineIsMap == true {
				switch string(data[y][x]) {
				case "@":
					object = robot
					object2 = Point{typePt: "EmptySpace", x: realX + 1, y: y, letter: "."}
					// Robot
					robot.x = realX
					robot.y = y
					break
				case "O":
					object = Point{typePt: "Box", x: realX, y: y, letter: "[", brotherX: realX + 1}
					object2 = Point{typePt: "Box", x: realX + 1, y: y, letter: "]", brotherX: realX}
					// Box
					break
				case "#":
					object = Point{typePt: "Wall", x: realX, y: y, letter: "#"}
					object2 = Point{typePt: "Wall", x: realX + 1, y: y, letter: "#"}
					// Wall
					break
				case ".":
					object = Point{typePt: "EmptySpace", x: realX, y: y, letter: "."}
					object2 = Point{typePt: "EmptySpace", x: realX + 1, y: y, letter: "."}
					// EmptySpace
					break
				}

				line = append(line, object, object2)
			} else {
				vectorMove := VectorMove{}
				// Directions
				instruction := string(data[y][x])
				switch string(data[y][x]) {
				case ">":
					vectorMove.x = 1
					break
				case "^":
					vectorMove.y = -1
					break
				case "v":
					vectorMove.y = 1
					break
				case "<":
					vectorMove.x = -1
					break
				}

				if instruction == ">" || instruction == "<" {
					Move(&robot, vectorMove, &carte)
				} else {
					Move2(&robot, vectorMove, &carte)
				}
			}
		}

		if lineIsMap == true {
			carte = append(carte, line)
		}
	}

	sum := 0
	for y := 0; y < len(carte); y++ {
		for x := 0; x < len(carte[y]); x++ {
			if carte[y][x].letter == "[" {
				sum += 100*y + x
			}
		}
	}

	return sum
}

func Resolve(data []string) [2]any {
	return [2]any{
		ResolvePart1(data),
		ResolvePart2(data),
	}
}
