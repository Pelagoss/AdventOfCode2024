package Day09

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func ResolvePart1(data []string) int {
	line := data[0]

	values := strings.Split(line, "")

	toggle := true
	highestId := int(math.RoundToEven(float64(len(values))/2) - 1)

	id := 0
	nGetEnd := 0
	used := make(map[int]int)
	used[highestId] = 0

	var disk []int

	for i := 0; nGetEnd+i < len(values)-1; i++ {
		val, err := strconv.Atoi(values[i])

		if err != nil {
			panic(err)
		}
		if toggle {
			for j := 0; j < val-used[id]; j++ {
				disk = append(disk, id)
			}
			used[id] += val
			toggle = false
			id++
		} else {
			for j := 1; j <= val; j++ {
				valBis, err := strconv.Atoi(values[len(values)-1-nGetEnd*2])

				if err != nil {
					panic(err)
				}

				disk = append(disk, highestId)

				if used[highestId] == valBis-1 {
					nGetEnd++
					highestId = highestId - 1
					used[highestId] = 0

					if highestId < id {
						break
					}
				} else {
					used[highestId]++
				}
			}
			toggle = true
		}

		if highestId < id {
			break
		}
	}

	sum := 0
	for i := 0; i < len(disk); i++ {
		sum += i * disk[i]
	}

	return sum
}

func insert(a []string, index int, value string) []string {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	if index < 0 {
		index = 0
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func ResolvePart2(data []string) int {
	line := data[0]

	values := strings.Split(line, "")

	toggle := true
	highestInitial := int(math.RoundToEven(float64(len(values))/2) - 1)

	id := 0
	nGetEnd := 0

	var disk []int
	var diskVide []int
	var diskDefragmented []string

	for i := 0; nGetEnd+i < len(values); i++ {
		val, err := strconv.Atoi(values[i])

		if err != nil {
			panic(err)
		}
		if toggle {
			for j := 0; j < val; j++ {
				diskDefragmented = append(diskDefragmented, fmt.Sprintf("%v", id))
			}
			//	disk = append(disk, fmt.Sprintf("%v", id))
			disk = append(disk, val)

			toggle = false
			id++
		} else {
			for j := 0; j < val; j++ {
				diskDefragmented = append(diskDefragmented, fmt.Sprintf("%v", "."))
			}
			diskVide = append(diskVide, val)
			toggle = true
		}
	}

	fmt.Println(disk)
	fmt.Println(diskVide)
	fmt.Println(highestInitial)
	fmt.Println(diskDefragmented)

	for i := 0; i < (len(disk)-1)+(len(diskVide)-1); i++ {
		diskDefragmented = append(diskDefragmented, ".")
	}

	sum := 0
	highest := highestInitial
	for i := len(diskVide); i > 0; i-- {
		index := slices.IndexFunc(diskVide, func(n int) bool { return n >= disk[i] })
		if index == -1 {
			highest--
			continue
		}
		fmt.Println(index)
		fmt.Println(diskVide[index])
		fmt.Println(disk[i])

		if diskVide[index] >= disk[i] {
			for j := 0; j < disk[i]; j++ {
				diskDefragmented = remove(diskDefragmented, (disk[i]-1)+(diskVide[index]-1)+j)
				diskDefragmented = insert(diskDefragmented, (disk[i]-1)+(diskVide[index]-1)+j, fmt.Sprintf("%v", highest))
			}
		}
		highest--

		val := disk[i]

		//if err != nil {
		//	panic(err)
		//}
		sum += i * val
		//fmt.Printf("%v + ", i*disk[i])
		fmt.Println(diskDefragmented)
		//panic("aa")
	}
	//fmt.Printf("\n")

	fmt.Println(diskDefragmented)
	return sum
}
func Resolve(data []string) [2]int {
	return [2]int{
		ResolvePart1(data),
		ResolvePart2(data),
	}
}
