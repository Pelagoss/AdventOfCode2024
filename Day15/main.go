package Day15

import (
	"fmt"
	"reflect"
)

type Movable struct {
	x, y int
}

type Robot Movable
type Box Movable

type Wall struct {
	x, y int
}

type EmptySpace struct {
	x, y int
}

func ResolvePart1(data []string) int {
	lineIsMap := true
	robot := Robot{}
	carte := [50][50]any{}

	for y := 0; y < len(data); y++ {
		if data[y] == "\n" {
			lineIsMap = false
			continue
		}

		if lineIsMap == true {
			var object any

			// Map
			// Setup la map
			for x := 0; x < len(data[y]); x++ {
				switch string(data[y][x]) {
				case "@":
					object = robot
					// Robot
					robot.x = x
					robot.y = y
					fmt.Println(reflect.TypeOf(object))
					break
				case "O":
					object = Box{x: x, y: y}
					// Box
					break
				case "#":
					object = Wall{x: x, y: y}
					// Wall
					break
				case ".":
					object = EmptySpace{x: x, y: y}
					// EmptySpace
					break
				}

				carte[y][x] = object
				//reflect.TypeOf(carte[y][x]) get type of variable
			}
		} else {
			// Directions
		}
	}
	return 0
}
func ResolvePart2(data []string) int {
	return 0
}
func Resolve(data []string) [2]int {
	return [2]int{
		ResolvePart1(data),
		ResolvePart2(data),
	}
}
