package Day12

import (
	"fmt"
	"sort"
	"strings"
)

type Region struct {
	id          int
	perimeter   int
	area        int
	letter      string
	perimetersY map[string]map[int]int
	perimetersX map[string]map[int]int
}

type Data struct {
	x, y     int
	regionId int
}

type Travel struct {
	x, y int
}

func setRegion(x int, y int, RegionData *map[int]map[int]Data, plants *[]string, listRegions *[]Region, regId int) {
	var reg Region

	if _, ok := (*RegionData)[y]; !ok {
		(*RegionData)[y] = make(map[int]Data)
	}

	if (*RegionData)[y][x].regionId != 0 {
		(*listRegions)[(*RegionData)[y][x].regionId-1].area += 1
		(*RegionData)[y][x] = Data{x: x, y: y, regionId: (*RegionData)[y][x].regionId}
		return
	}

	if regId == 0 {
		regId = (*RegionData)[y][x].regionId
	}

	if regId == 0 || len(*listRegions) == 0 {

		reg = Region{id: len(*listRegions) + 1, perimeter: 0, area: 0, letter: (*plants)[x], perimetersY: make(map[string]map[int]int), perimetersX: make(map[string]map[int]int)}
		*listRegions = append(*listRegions, reg)
		regId = reg.id
	}

	(*listRegions)[regId-1].area += 1

	(*RegionData)[y][x] = Data{x: x, y: y, regionId: regId}
}

func setPerimeterX(X string, Y int, listRegions *[]Region, regId int) {
	if _, ok := (*listRegions)[regId-1].perimetersX[X]; !ok {
		(*listRegions)[regId-1].perimetersX[X] = make(map[int]int)
	}

	(*listRegions)[regId-1].perimetersX[X][Y] = 1
}

func setPerimeterY(X int, Y string, listRegions *[]Region, regId int) {
	if _, ok := (*listRegions)[regId-1].perimetersY[Y]; !ok {
		(*listRegions)[regId-1].perimetersY[Y] = make(map[int]int)
	}

	(*listRegions)[regId-1].perimetersY[Y][X] = 1
}

func greedySearchNextPoint(greedySearch []Travel, RegionData *map[int]map[int]Data, plants *[]string, listRegions *[]Region, data *[]string, perimeter *int, regId int) {
	var nextGreedySearch []Travel
	for i := 0; i < len(greedySearch); i++ {
		curX := (greedySearch)[i].x
		curY := (greedySearch)[i].y

		if (*RegionData)[curY][curX].regionId != 0 {
			continue
		}

		setRegion(curX, curY, RegionData, plants, listRegions, regId)
		regId = (*RegionData)[curY][curX].regionId

		if curX == 0 || curX == len(*plants)-1 {
			*perimeter++
			indexX := "right"
			if curX == 0 {
				indexX = "left"
			}

			setPerimeterX(fmt.Sprintf("%v%v", indexX, curX), curY, listRegions, regId)
		}
		if curY == 0 || curY == len(*data)-1 {
			*perimeter++

			indexY := "bot"
			if curY == 0 {
				indexY = "top"
			}

			setPerimeterY(curX, fmt.Sprintf("%v%v", indexY, curY), listRegions, regId)
		}

		if curX > 0 && (*data)[curY][curX-1] != (*data)[curY][curX] {
			*perimeter++

			setPerimeterX(fmt.Sprintf("left%v", curX), curY, listRegions, regId)
		} else if curX > 0 && (*data)[curY][curX-1] == (*data)[curY][curX] && (*RegionData)[curY][curX-1].regionId == 0 {
			nextGreedySearch = append(nextGreedySearch, Travel{curX - 1, curY})
		}

		if curX < len(*plants)-1 && (*data)[curY][curX+1] != (*data)[curY][curX] {
			*perimeter++

			setPerimeterX(fmt.Sprintf("right%v", curX), curY, listRegions, regId)
		} else if curX < len(*plants)-1 && (*data)[curY][curX+1] == (*data)[curY][curX] && (*RegionData)[curY][curX+1].regionId == 0 {
			nextGreedySearch = append(nextGreedySearch, Travel{curX + 1, curY})
		}

		if curY > 0 && (*data)[curY-1][curX] != (*data)[curY][curX] {
			*perimeter++
			setPerimeterY(curX, fmt.Sprintf("top%v", curY), listRegions, regId)
		} else if curY > 0 && (*data)[curY-1][curX] == (*data)[curY][curX] && (*RegionData)[curY-1][curX].regionId == 0 {
			nextGreedySearch = append(nextGreedySearch, Travel{curX, curY - 1})
		}

		if curY < len(*data)-1 && (*data)[curY+1][curX] != (*data)[curY][curX] {
			*perimeter++
			setPerimeterY(curX, fmt.Sprintf("bot%v", curY), listRegions, regId)
		} else if curY < len(*data)-1 && (*data)[curY+1][curX] == (*data)[curY][curX] && (*RegionData)[curY+1][curX].regionId == 0 {
			nextGreedySearch = append(nextGreedySearch, Travel{curX, curY + 1})
		}

		(*listRegions)[regId-1].perimeter += *perimeter
	}

	if len(nextGreedySearch) > 0 {
		greedySearchNextPoint(nextGreedySearch, RegionData, plants, listRegions, data, perimeter, regId)
	}
}

func ResolveParts(data []string) [2]any {
	RegionData := make(map[int]map[int]Data)

	var listRegions []Region

	for y := 0; y < len(data); y++ {
		plants := strings.Split(data[y], "")

		for x := 0; x < len(plants); x++ {
			if RegionData[y][x].regionId != 0 {
				continue
			}

			perimeter := 0
			greedySearch := []Travel{{x: x, y: y}}

			greedySearchNextPoint(greedySearch, &RegionData, &plants, &listRegions, &data, &perimeter, RegionData[y][x].regionId)
			listRegions[RegionData[y][x].regionId-1].perimeter = perimeter
		}
	}

	cost1 := 0
	for k := range listRegions {
		cost1 += listRegions[k].area * listRegions[k].perimeter
	}

	cost2 := 0
	for k := range listRegions {
		newK := k

		sides1 := 0
		for kY := range listRegions[newK].perimetersY {
			newkY := kY
			previous1 := -1

			keys := make([]int, 0, len(listRegions[newK].perimetersY[newkY]))
			for k := range listRegions[newK].perimetersY[newkY] {
				keys = append(keys, k)
			}
			sort.Ints(keys)

			sides1++
			for i := 0; i < len(keys); i++ {
				newkX := keys[i]
				if previous1 == -1 || previous1+1 == newkX {
					previous1 = newkX
				} else {
					previous1 = newkX
					sides1++
				}
			}
		}

		sides2 := 0
		for kY := range listRegions[newK].perimetersX {
			newkY := kY
			previous2 := -1

			keys2 := make([]int, 0, len(listRegions[newK].perimetersX[newkY]))
			for k := range listRegions[newK].perimetersX[newkY] {
				keys2 = append(keys2, k)
			}
			sort.Ints(keys2)

			sides2++
			for i := 0; i < len(keys2); i++ {
				newkX := keys2[i]
				if previous2 == -1 || previous2+1 == newkX {
					previous2 = newkX
				} else {
					previous2 = newkX
					sides2++
				}
			}
		}

		cost2 += listRegions[newK].area * (sides1 + sides2)
	}
	return [2]any{cost1, cost2}
}

func Resolve(data []string) [2]any {
	return ResolveParts(data)
}
