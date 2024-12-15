package Day09

import (
	"fmt"
	"math"
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
func ResolvePart2(data []string) int {
	line := data[0]

	values := strings.Split(line, "")

	toggle := true
	highestInitial := int(math.RoundToEven(float64(len(values))/2) - 1)

	id := 0
	nGetEnd := 0
	used := make(map[int]int)
	usedId := make(map[int]bool)

	var disk []string

	for i := 0; nGetEnd+i < len(values)-1; i++ {
		val, err := strconv.Atoi(values[i])

		if err != nil {
			panic(err)
		}
		if toggle {
			for j := 0; j < val-used[id]; j++ {
				disk = append(disk, fmt.Sprintf("%v", id))
			}

			if val-used[id] <= 0 {
				disk = append(disk, ".")
			}

			used[id] += val - used[id]
			toggle = false
			id++
		} else {
			valBis := -1
			index := -1
			highestId := highestInitial

			for k := 0; val > 0; k++ {
				for j := len(values); j > 0; j -= 2 {
					valBis, err = strconv.Atoi(values[j-1])

					if err != nil {
						panic(err)
					}

					if _, ok := usedId[j-1]; ok {
						highestId = highestId - 1
						continue
					}

					if valBis <= val {
						usedId[j-1] = true
						fmt.Printf("ezhehnegne %v %v\n", valBis, val)
						break
					}
					highestId = highestId - 1
				}

				if highestId < 0 {
					disk = append(disk, ".")
					usedId[index] = true
					val = 0
				}
				fmt.Println(highestId)
				fmt.Println(valBis)

				if _, ok := used[highestId]; ok {
					break
				}

				used[highestId] = valBis
				for j := 1; j <= valBis; j++ {
					disk = append(disk, fmt.Sprintf("%v", highestId))
				}

				fmt.Println(disk)

				val = val - valBis

				nGetEnd = 0
				highestId = highestInitial

				fmt.Printf("reste %v places, nGetEnd %v, highestId %v index %v\n", val, nGetEnd, highestId, len(values)-1-nGetEnd)
			}

			toggle = true
		}
	}

	//fmt.Println(disk)

	sum := 0
	for i := 0; i < len(disk); i++ {
		if disk[i] == "." {
			continue
		}

		val, err := strconv.Atoi(disk[i])

		if err != nil {
			panic(err)
		}
		sum += i * val
		//fmt.Printf("%v + ", i*disk[i])
	}
	//fmt.Printf("\n")

	return sum
}
func Resolve(data []string) [2]int {
	return [2]int{
		ResolvePart1(data),
		ResolvePart2(data),
	}
}
