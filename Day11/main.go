package Day11

import (
	"fmt"
	"strconv"
	"strings"
)

type Result struct {
	i      int
	stones []string
}

func BlinkStone(stone string) []string {
	var nextStones []string

	if stone == "0" {
		nextStones = append(nextStones, "1")
	} else if len(stone)%2 == 0 {
		middle := len(stone) / 2
		nextStones = append(nextStones, stone[:middle])
		stoneValue, _ := strconv.Atoi(stone[middle:])
		if stoneValue == 0 {
			nextStones = append(nextStones, "0")
		} else {
			nextStones = append(nextStones, fmt.Sprintf("%d", stoneValue))
		}
	} else {
		stoneValue, _ := strconv.Atoi(stone)

		nextStones = append(nextStones, fmt.Sprintf("%d", stoneValue*2024))
	}

	return nextStones
}

func Solve(data []string) [2]int {
	stones := strings.Split(data[0], " ")

	part1 := 0

	oldRocks := make(map[string]int, len(stones))
	for i := 0; i < len(stones); i++ {
		oldRocks[stones[i]] = 1
	}

	for blink := 0; blink < 75; blink++ {
		// Pour la partie 1 on s'arrete a 25
		if blink == 25 {
			for _, v := range oldRocks {
				part1 += v
			}
		}

		newRocks := make(map[string]int)
		for v, count := range oldRocks {
			r := BlinkStone(v)
			for _, rock := range r {
				if _, ok := newRocks[rock]; ok {
					newRocks[rock] += 1 * count
				} else {
					newRocks[rock] = 1 * count
				}
			}
		}
		oldRocks = newRocks
	}

	part2 := 0
	for _, v := range oldRocks {
		part2 += v
	}

	return [2]int{part1, part2}
}
func Resolve(data []string) [2]int {
	return Solve(data)
}
