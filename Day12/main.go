package Day12

import (
	"fmt"
	"strings"
)

type Region struct {
	id        int
	perimeter int
	area      int
	letter    string
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
		reg = Region{id: len(*listRegions) + 1, perimeter: 0, area: 0, letter: (*plants)[x]}
		*listRegions = append(*listRegions, reg)
		regId = reg.id
	}

	(*listRegions)[regId-1].area += 1

	(*RegionData)[y][x] = Data{x: x, y: y, regionId: regId}
	fmt.Println(Data{x: x, y: y, regionId: regId})
}

func greedySearchNextPoint(greedySearch []Travel, RegionData *map[int]map[int]Data, plants *[]string, listRegions *[]Region, data *[]string, perimeter *int, regId int) {
	var nextGreedySearch []Travel
	for i := 0; i < len(greedySearch); i++ {
		curX := (greedySearch)[i].x
		curY := (greedySearch)[i].y

		if (*RegionData)[curY][curX].regionId != 0 {
			continue
		}

		fmt.Printf("(x: %v, y: %v)", curX, curY)

		setRegion(curX, curY, RegionData, plants, listRegions, regId)
		regId = (*RegionData)[curY][curX].regionId

		if curX == 0 || curX == len(*plants)-1 {
			*perimeter++
		}
		if curY == 0 || curY == len(*data)-1 {
			*perimeter++
		}

		if curX > 0 && (*data)[curY][curX-1] != (*data)[curY][curX] {
			*perimeter++
		} else if curX > 0 && (*data)[curY][curX-1] == (*data)[curY][curX] && (*RegionData)[curY][curX-1].regionId == 0 {
			nextGreedySearch = append(nextGreedySearch, Travel{curX - 1, curY})
			//setRegion(curX-1, y, &RegionData, &plants, &listRegions, regId)
			//fmt.Println("change regId", (*data)[y][curX], "===", (*data)[y][curX-1])
		}

		if curX < len(*plants)-1 && (*data)[curY][curX+1] != (*data)[curY][curX] {
			*perimeter++
		} else if curX < len(*plants)-1 && (*data)[curY][curX+1] == (*data)[curY][curX] && (*RegionData)[curY][curX+1].regionId == 0 {
			nextGreedySearch = append(nextGreedySearch, Travel{curX + 1, curY})
			//setRegion(curX+1, y, &RegionData, &plants, &listRegions, regId)
			//fmt.Println("change regId", (*data)[y][curX], "===", (*data)[y][curX+1])
		}

		if curY > 0 && (*data)[curY-1][curX] != (*data)[curY][curX] {
			*perimeter++
		} else if curY > 0 && (*data)[curY-1][curX] == (*data)[curY][curX] && (*RegionData)[curY-1][curX].regionId == 0 {
			nextGreedySearch = append(nextGreedySearch, Travel{curX, curY - 1})
			//setRegion(curX, y-1, &RegionData, &plants, &listRegions, regId)
			//fmt.Println("change regId", (*data)[y][curX], "===", (*data)[y-1][curX])
		}

		if curY < len(*data)-1 && (*data)[curY+1][curX] != (*data)[curY][curX] {
			*perimeter++
		} else if curY < len(*data)-1 && (*data)[curY+1][curX] == (*data)[curY][curX] && (*RegionData)[curY+1][curX].regionId == 0 {
			nextGreedySearch = append(nextGreedySearch, Travel{curX, curY + 1})
			//setRegion(curX, y+1, &RegionData, &plants, &listRegions, regId)
			//fmt.Println("change regId", (*data)[y][curX], "===", (*data)[y+1][curX])
		}

		(*listRegions)[regId-1].perimeter += *perimeter
	}
	fmt.Println("")
	if len(nextGreedySearch) > 0 {
		greedySearchNextPoint(nextGreedySearch, RegionData, plants, listRegions, data, perimeter, regId)
	}
}

func ResolvePart1(data []string) int {
	RegionData := make(map[int]map[int]Data)

	var listRegions []Region

	for y := 0; y < len(data); y++ {
		plants := strings.Split(data[y], "")

		for x := 0; x < len(plants); x++ {
			if RegionData[y][x].regionId != 0 {
				continue
			}
			fmt.Printf("%v \n", RegionData[y][x])
			fmt.Printf("%v : %v \n", plants[x], Travel{x: x, y: y})

			perimeter := 0
			greedySearch := []Travel{{x: x, y: y}}

			greedySearchNextPoint(greedySearch, &RegionData, &plants, &listRegions, &data, &perimeter, RegionData[y][x].regionId)
			fmt.Printf("%v \n", RegionData[2])

			fmt.Println("======================")

			fmt.Sprintf("%v %v %v", x, y, RegionData[y][x])
			listRegions[RegionData[y][x].regionId-1].perimeter = perimeter

			//fmt.Println(data[y][x], regId == -1 || len(listRegions) == 0, regId)

			//fmt.Println("regData", RegionData)
			//fmt.Println(listRegions)
			//for k := range RegionData {
			//	for i := 0; i < len(RegionData[k]); i++ {
			//		fmt.Printf("%02d ", RegionData[k][i].regionId)
			//	}
			//	fmt.Printf("\n")
			//}
			//panic("aaa")
			//if y > 1 {
			//	panic("aaa")
			//}
		}
	}

	fmt.Println("regData", RegionData)
	for k := range RegionData {
		for i := 0; i < len(RegionData[k]); i++ {
			fmt.Printf("%02d ", RegionData[k][i].regionId)
		}
		fmt.Printf("\n")
	}
	fmt.Println(listRegions)

	cost := 0
	for k := range listRegions {
		cost += listRegions[k].area * listRegions[k].perimeter
	}

	return cost
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
