package Day14

import (
	"fmt"
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
	for i := 0; i < nMoves; i++ {
		(*robot).x = pmod((*robot).x+(*robot).xStep, widthMax)
		(*robot).y = pmod((*robot).y+(*robot).yStep, heightMax)
	}
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

	//fmt.Printf("%v %v %v %v\n", topLeft, topRight, bottomLeft, bottomRight)

	return topLeft * topRight * bottomLeft * bottomRight
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
