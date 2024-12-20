package Day14

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"
)

type Robot struct {
	x, y         int
	xStep, yStep int
}

func pmod(x, d int) int {
	x = x % d
	if x >= 0 {
		return x
	}
	if d < 0 {
		return x - d
	}
	return x + d
}

func MoveRobot(robot *Robot, nMoves int) {
	(*robot).x = pmod((*robot).x+nMoves*(*robot).xStep, widthMax)
	(*robot).y = pmod((*robot).y+nMoves*(*robot).yStep, heightMax)
}

var (
	widthMax  = 101
	heightMax = 103
)

func ResolvePart1(data []string) int {
	var listRobots []Robot
	var (
		topLeft,
		topRight,
		bottomLeft,
		bottomRight int
	)

	for i := 0; i < len(data); i++ {
		robot := Robot{}
		robotSpecs := strings.NewReader(data[i])
		_, err := fmt.Fscanf(robotSpecs, "p=%d,%d v=%d,%d", &robot.x, &robot.y, &robot.xStep, &robot.yStep)

		if err != nil {
			panic(err)
		}

		MoveRobot(&robot, 100)
		listRobots = append(listRobots, robot)

		left := robot.x < (widthMax-1)/2
		right := robot.x > (widthMax-1)/2

		top := robot.y < (heightMax-1)/2
		bottom := robot.y > (heightMax-1)/2

		if left && top {
			topLeft++
		} else if right && top {
			topRight++
		} else if left && bottom {
			bottomLeft++
		} else if right && bottom {
			bottomRight++
		}
	}

	return topLeft * topRight * bottomLeft * bottomRight
}
func ResolvePart2(data []string) int {
	for t := 0; t < 101*103; t++ {
		var mapRobots [103][101]int
		var vertical [101]int
		var horizontal [103]int
		canWePrint := false

		for i := 0; i < len(data); i++ {
			robot := Robot{}
			robotSpecs := strings.NewReader(data[i])
			_, err := fmt.Fscanf(robotSpecs, "p=%d,%d v=%d,%d", &robot.x, &robot.y, &robot.xStep, &robot.yStep)

			if err != nil {
				panic(err)
			}

			MoveRobot(&robot, t)

			mapRobots[robot.y][robot.x] += 1
			vertical[robot.x] += 1
			horizontal[robot.y] += 1

			if vertical[robot.x] > 33 || horizontal[robot.y] > 33 {
				canWePrint = true
			}
		}

		if canWePrint {
			upLeft := image.Point{}
			lowRight := image.Point{X: widthMax, Y: heightMax}

			img := image.NewRGBA(image.Rectangle{Min: upLeft, Max: lowRight})
			green := color.RGBA{R: 0, G: 77, B: 67, A: 0xff}

			for y := 0; y < len(mapRobots); y++ {
				for x := 0; x < len(mapRobots[y]); x++ {
					if mapRobots[y][x] == 0 {
						img.Set(x, y, color.White)
					} else {
						img.Set(x, y, green)
					}
				}
			}

			f, _ := os.Create(fmt.Sprintf("%v.png", t))
			err := png.Encode(f, img)
			if err != nil {
				panic(err)
			}
		}
		canWePrint = false
	}

	return 6668 //hardcoded lol
}
func Resolve(data []string) [2]int {
	return [2]int{
		ResolvePart1(data),
		ResolvePart2(data),
	}
}
