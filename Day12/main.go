package Day12

import (
	"fmt"
	"strings"
)

type Region struct {
	id        int
	perimeter int
	area      int
}

type Data struct {
	x, y   int
	region Region
}

func ResolvePart1(data []string) int {
	RegionData := make(map[int]map[int]Data)

	var listRegions []Region

	regions := 0
	for y := 0; y < len(data); y++ {
		plants := strings.Split(data[y], "")

		for x := 0; x < len(plants); x++ {
			reg := Region{id: regions, perimeter: 0, area: 0}

			perimeter := 0
			if x == 0 || x == len(plants)-1 {
				perimeter++
			}
			if y == 0 || y == len(data)-1 {
				perimeter++
			}
			if x > 0 && data[y][x-1] != data[y][x] {
				perimeter++
			} else if x > 0 {
				reg = RegionData[y][x-1].region
			}

			if x < len(plants)-1 && data[y][x+1] != data[y][x] {
				perimeter++
			} else if x < len(plants)-1 {
				reg = RegionData[y][x+1].region
			}
			if y > 0 && data[y-1][x] != data[y][x] {
				perimeter++
			} else if y > 0 {
				reg = RegionData[y-1][x].region
			}
			if y < len(data)-1 && data[y+1][x] != data[y][x] {
				perimeter++
			} else if y < len(data)-1 {
				reg = RegionData[y+1][x].region
			}

			reg.area += 1
			reg.perimeter += perimeter

			if _, ok := RegionData[y]; !ok {
				RegionData[y] = make(map[int]Data)
			}

			RegionData[y][x] = Data{x: x, y: y, region: reg}

			if reg.id == regions {
				listRegions = append(listRegions, reg)
			}

			regions++
		}
	}

	//fmt.Println("Perimeter : ", RegionPerimeter)
	//fmt.Println("Area : ", RegionArea)
	fmt.Println("Data : ", RegionData)

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
