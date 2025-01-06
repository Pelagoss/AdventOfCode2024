package Day19

import (
	"container/heap"
	"fmt"
	"regexp"
	"slices"
	"strings"
)

func ResolvePart1(data []string) int {
	var patterns string
	design := false
	count := 0

	for i := 0; i < len(data); i++ {
		if data[i] == "" {
			design = true
			continue
		}

		if design {
			re := regexp.MustCompile(patterns)

			if re.ReplaceAllString(data[i], "") == "" {
				count++
			}
		} else {
			strSet := strings.Split(data[i], ", ")

			slices.SortFunc(strSet, func(i, j string) int {
				return len(j) - len(i)
			})

			patterns = fmt.Sprintf("^(%s)*$", strings.Join(strSet, "|"))
		}
	}

	return count
}

type Result struct {
	score int
	pos   [2]int
	path  [][2]int
}

type ResultHeap []*Result

func (h ResultHeap) Len() int           { return len(h) }
func (h ResultHeap) Less(i, j int) bool { return h[i].score < h[j].score }
func (h ResultHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *ResultHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*Result))
}

func (h *ResultHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func ResolvePart2(data []string) int {
	var patterns string
	var patternSet []string
	design := false
	count := 0

	for i := 0; i < len(data); i++ {
		if data[i] == "" {
			design = true
			continue
		}

		if design {
			re := regexp.MustCompile(patterns)

			if re.ReplaceAllString(data[i], "") == "" {
				var t [][]int
				t2 := make(map[int][][2]int)

				for j := 0; j < len(patternSet); j++ {
					reBis := regexp.MustCompile(patternSet[j])
					found := reBis.FindAllStringSubmatchIndex(data[i], -1)

					if len(found) > 0 {
						for l := 0; l < len(found); l++ {
							t = append(t, found[l])
							t2[found[l][0]] = append(t2[found[l][0]], [2]int{found[l][0], found[l][1]})
						}
					}
				}

				var path [][][2]int

				heapResult := &ResultHeap{}
				heap.Init(heapResult)
				heap.Push(heapResult, &Result{
					score: 0,
					pos:   [2]int{0, 0},
					path:  make([][2]int, 0),
				})
				visited := make(map[[3]int]bool)

				for heapResult.Len() > 0 {
					h := heap.Pop(heapResult).(*Result)

					if _, ok := visited[[3]int{h.pos[0], h.pos[1], h.score}]; ok {
						continue
					}

					if len(h.path) > 1 && h.path[len(h.path)-1][1] == len(data[i]) {
						path = append(path, h.path)
						continue
					}

					visited[[3]int{h.pos[0], h.pos[1], h.score}] = true

					if _, ok := t2[h.pos[1]]; ok {
						for j := 0; j < len(t2[h.pos[1]]); j++ {
							toPush := t2[h.pos[1]][j]
							newPath := make([][2]int, 0)
							for k := 0; k < len(h.path); k++ {
								newPath = append(newPath, h.path[k])
							}
							newPath = append(newPath, toPush)

							heap.Push(heapResult, &Result{
								score: h.score + 1,
								pos:   toPush,
								path:  newPath,
							})
						}
					}
				}

				fmt.Println(path, len(path))
				fmt.Println(t2)
				fmt.Println()

				count += len(path)
			}
		} else {
			patternSet = strings.Split(data[i], ", ")

			slices.SortFunc(patternSet, func(i, j string) int {
				return len(j) - len(i)
			})

			patterns = fmt.Sprintf("^(%s)*$", strings.Join(patternSet, "|"))
		}
	}

	return count
}
func Resolve(data []string) [2]any {
	return [2]any{
		ResolvePart1(data),
		ResolvePart2(data), // 45570 too low
	}
}
