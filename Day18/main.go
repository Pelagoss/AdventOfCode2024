package Day18

import (
	"container/heap"
	"fmt"
	"slices"
	"strings"
)

type Memory struct {
	carte [][]string
	start [2]int
	end   [2]int
}

type Historian struct {
	pos       [2]int
	direction Direction
	cost      int
}

type Direction struct {
	vecteur [2]int
}

var (
	nextDirections = []Direction{{[2]int{1, 0}}, {[2]int{0, 1}}, {[2]int{-1, 0}}, {[2]int{0, -1}}}
)

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

func buildMap(data []string) (Memory, Historian) {
	var carte [][]string
	memory := Memory{}
	historians := Historian{}

	for y := 0; y <= 70; y++ {
		var line []string
		for x := 0; x <= 70; x++ {
			if y == 0 && x == 0 {
				memory.start = [2]int{x, y}
				historians = Historian{pos: [2]int{x, y}}
			}

			if y == 70 && x == 70 {
				memory.end = [2]int{x, y}
			}

			line = append(line, ".")
		}
		carte = append(carte, line)
	}

	memory.carte = loseBytes(carte, 0, 1024, data)

	return memory, historians
}

func loseBytes(carte [][]string, start int, end int, data []string) [][]string {
	for i := start; i < len(data) && i < end; i++ {
		var x, y int

		_, err := fmt.Sscanf(strings.ReplaceAll(data[i], ",", " , "), "%d , %d", &x, &y)
		if err != nil {
			panic(err)
		}

		carte[y][x] = "#"
	}

	return carte
}

func visitMemory(memory Memory, historians Historian) (bool, [][2]int, int) {
	heapResult := &ResultHeap{}
	heap.Init(heapResult)
	heap.Push(heapResult, &Result{0, historians.pos, make([][2]int, 0)})
	visited := make(map[[3]int]bool)

	for heapResult.Len() > 0 {
		h := heap.Pop(heapResult).(*Result)

		if h.pos == memory.end {
			return true, append(h.path, h.pos), h.score
		}

		if _, ok := visited[[3]int{h.pos[0], h.pos[1]}]; ok {
			continue
		}

		visited[[3]int{h.pos[0], h.pos[1]}] = true

		for i := 0; i < len(nextDirections); i++ {
			x := h.pos[0] + nextDirections[i].vecteur[0]
			y := h.pos[1] + nextDirections[i].vecteur[1]

			isNotInMap := x < 0 || y < 0 || x >= len(memory.carte[0]) || y >= len(memory.carte)

			if _, ok := visited[[3]int{x, y}]; !ok && !isNotInMap && memory.carte[y][x] == "." {
				heap.Push(heapResult, &Result{score: h.score + 1, pos: [2]int{x, y}, path: append(h.path, h.pos)})
			}
		}
	}

	return false, make([][2]int, 0), 0
}

func ResolvePart1(data []string) int {
	memory, historians := buildMap(data)

	_, _, score := visitMemory(memory, historians)

	return score
}
func ResolvePart2(data []string) string {
	memory, historians := buildMap(data)
	start := 1024

	for {
		finished, path, _ := visitMemory(memory, historians)

		if finished {
			var i int

			// Find the first lost memory position that is in the path, lost it and retry
			for i = start; i < len(data); i++ {
				var x, y int

				_, err := fmt.Sscanf(strings.ReplaceAll(data[i], ",", " , "), "%d , %d", &x, &y)
				if err != nil {
					panic(err)
				}

				memory.carte[y][x] = "#"

				if slices.Index(path, [2]int{x, y}) != -1 {
					start = i
					break
				}
			}
		} else {
			break
		}
	}

	return data[start]
}
func Resolve(data []string) [2]any {
	return [2]any{
		ResolvePart1(data),
		ResolvePart2(data),
	}
}
