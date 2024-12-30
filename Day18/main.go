package Day18

import (
	"container/heap"
	"fmt"
	"strings"
)

type Maze struct {
	carte    [][]string
	start    [2]int
	end      [2]int
	cost     int
	finished bool
}

type Reindeer struct {
	pos       [2]int
	direction Direction
	cost      int
	lastpos   [2]int
}

type Direction struct {
	vecteur  [2]int
	letter   string
	opposite string
}

var (
	nextDirections = []Direction{{[2]int{1, 0}, ">", "<"}, {[2]int{0, 1}, "v", "^"}, {[2]int{-1, 0}, "<", ">"}, {[2]int{0, -1}, "^", "v"}}
)

type Result struct {
	score    int
	dirIndex int
	pos      [2]int
	path     map[[2]int]bool
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
func buildMap(data []string) (Maze, Reindeer) {
	var carte [][]string
	maze := Maze{}
	reindeer := Reindeer{}

	for y := 0; y <= 70; y++ {
		var line []string
		for x := 0; x <= 70; x++ {
			if y == 0 && x == 0 {
				maze.start = [2]int{x, y}
				reindeer = Reindeer{pos: [2]int{x, y}, lastpos: [2]int{-1, -1}}
			}

			if y == 70 && x == 70 {
				maze.end = [2]int{x, y}
			}

			line = append(line, ".")
		}
		carte = append(carte, line)
	}

	for i := 0; i < 1024; i++ {
		var x, y int

		_, err := fmt.Sscanf(strings.ReplaceAll(data[i], ",", " , "), "%d , %d", &x, &y)
		if err != nil {
			panic(err)
		}

		carte[y][x] = "#"
	}

	maze.carte = carte

	return maze, reindeer
}

func ResolvePart1(data []string) int {
	maze, reindeer := buildMap(data)

	heapResult := &ResultHeap{}
	heap.Init(heapResult)
	heap.Push(heapResult, &Result{0, 0, reindeer.pos, make(map[[2]int]bool)})
	visited := make(map[[3]int]bool)

	for heapResult.Len() > 0 {
		h := heap.Pop(heapResult).(*Result)

		if h.pos == maze.end {
			return h.score
		}

		if _, ok := visited[[3]int{h.dirIndex, h.pos[0], h.pos[1]}]; ok {
			continue
		}

		visited[[3]int{h.dirIndex, h.pos[0], h.pos[1]}] = true

		x := h.pos[0] + nextDirections[h.dirIndex].vecteur[0]
		y := h.pos[1] + nextDirections[h.dirIndex].vecteur[1]

		isNotInMap := x < 0 || y < 0 || x >= len(maze.carte[0]) || y >= len(maze.carte)

		if _, ok := visited[[3]int{h.dirIndex, x, y}]; !isNotInMap && maze.carte[y][x] == "." && !ok {
			heap.Push(heapResult, &Result{score: h.score + 1, dirIndex: h.dirIndex, pos: [2]int{x, y}})
		}

		left := pmod(h.dirIndex-1, 4)
		if _, ok1 := visited[[3]int{left, x, y}]; !ok1 {
			heap.Push(heapResult, &Result{score: h.score, dirIndex: left, pos: h.pos})
		}

		right := pmod(h.dirIndex+1, 4)
		if _, ok2 := visited[[3]int{right, x, y}]; !ok2 {
			heap.Push(heapResult, &Result{score: h.score, dirIndex: right, pos: h.pos})
		}
	}

	return 0
}
func ResolvePart2(data []string) int {
	return 0
}
func Resolve(data []string) [2]any {
	return [2]any{
		ResolvePart1(data),
		ResolvePart2(data),
	}
}
