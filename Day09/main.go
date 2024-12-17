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
	highestInitial := 0

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
			disk = append(disk, val)

			if id > highestInitial {
				highestInitial = id
			}

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

	for i := 0; i < (len(disk)-1)+(len(diskVide)-1); i++ {
		diskDefragmented = append(diskDefragmented, ".")
	}

	highest := highestInitial
	for i := len(disk) - 1; i > 0; i-- {
		index := -1

		indexMax := slices.Index(diskDefragmented, fmt.Sprintf("%v", highest))
		countDot := 0

		for l := 0; l < indexMax; l++ {
			countDot = 0

			for k := 0; k < disk[i]; k++ {
				if diskDefragmented[l+k] == "." {
					countDot++
				}
			}

			if countDot == disk[i] && countDot > 0 {
				index = l
				break
			}
		}

		if index == -1 {
			highest--
			continue
		}

		indexToDel := slices.Index(diskDefragmented, fmt.Sprintf("%v", highest))
		if indexToDel < 0 {
			continue
		}

		for j := 0; j < disk[i]; j++ {
			diskDefragmented = remove(diskDefragmented, indexToDel+j)
			diskDefragmented = insert(diskDefragmented, indexToDel+j, ".")

			diskDefragmented = remove(diskDefragmented, index+j)
			diskDefragmented = insert(diskDefragmented, index+j, fmt.Sprintf("%v", highest))
		}
		highest--
	}

	sum := 0
	for i := 0; i < len(diskDefragmented); i++ {
		if diskDefragmented[i] == "." {
			continue
		}

		val, err := strconv.Atoi(diskDefragmented[i])
		if err != nil {
			panic(err)
		}

		sum += i * val
	}

	return sum
}
func Resolve(data []string) [2]int {
	return [2]int{
		ResolvePart1(data),
		ResolvePart2(data),
	}
}
