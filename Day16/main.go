package Day16

import (
	"container/heap"
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

	for y := 0; y < len(data); y++ {
		var line []string
		for x := 0; x < len(data[y]); x++ {
			letter := string(data[y][x])
			switch string(data[y][x]) {
			case "E":
				maze.end = [2]int{x, y}
				letter = "."
				break
			case "S":
				reindeer = Reindeer{direction: Direction{vecteur: [2]int{1, 0}, letter: ">", opposite: "<"}, pos: [2]int{x, y}, lastpos: [2]int{-1, -1}}
				maze.start = [2]int{x, y}
				letter = "."
				break
			}
			line = append(line, letter)
		}
		carte = append(carte, line)
	}

	maze.carte = carte

	return maze, reindeer
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
		if _, ok := visited[[3]int{h.dirIndex, x, y}]; maze.carte[y][x] == "." && !ok {
			heap.Push(heapResult, &Result{score: h.score + 1, dirIndex: h.dirIndex, pos: [2]int{x, y}})
		}

		left := pmod(h.dirIndex-1, 4)
		if _, ok1 := visited[[3]int{left, h.pos[0], h.pos[1]}]; !ok1 {
			heap.Push(heapResult, &Result{score: h.score + 1000, dirIndex: left, pos: h.pos})
		}

		right := pmod(h.dirIndex+1, 4)
		if _, ok2 := visited[[3]int{right, h.pos[0], h.pos[1]}]; !ok2 {
			heap.Push(heapResult, &Result{score: h.score + 1000, dirIndex: right, pos: h.pos})
		}
	}

	return 0
}
func canVisit(dirIndex int, pos [2]int, score int, visited *map[[3]int]int) bool {
	if prevScore, ok := (*visited)[[3]int{dirIndex, pos[0], pos[1]}]; ok && prevScore < score {
		return false
	}
	(*visited)[[3]int{dirIndex, pos[0], pos[1]}] = score
	return true
}

func ResolvePart2(data []string) int {
	maze, reindeer := buildMap(data)
	heapResult := &ResultHeap{}
	heap.Init(heapResult)
	validPath := make(map[[2]int]bool)
	validPath[reindeer.pos] = true

	heap.Push(heapResult, &Result{0, 0, reindeer.pos, validPath})
	visited := make(map[[3]int]int)
	lowestScore := -1

	for heapResult.Len() > 0 {
		h := heap.Pop(heapResult).(*Result)

		if lowestScore != -1 && lowestScore < h.score {
			break
		}

		if h.pos == maze.end {
			lowestScore = h.score

			sUnion := make(map[[2]int]bool)
			for k, _ := range h.path {
				sUnion[k] = true
			}
			for k, _ := range validPath {
				sUnion[k] = true
			}

			validPath = sUnion
			continue
		}

		if !canVisit(h.dirIndex, h.pos, h.score, &visited) {
			continue
		}

		x := h.pos[0] + nextDirections[h.dirIndex].vecteur[0]
		y := h.pos[1] + nextDirections[h.dirIndex].vecteur[1]
		if maze.carte[y][x] == "." && canVisit(h.dirIndex, [2]int{x, y}, h.score+1, &visited) {
			sUnion := make(map[[2]int]bool)
			for k, _ := range h.path {
				sUnion[k] = true
			}
			if _, ok := sUnion[[2]int{x, y}]; !ok {
				sUnion[[2]int{x, y}] = true
			}
			heap.Push(heapResult, &Result{score: h.score + 1, dirIndex: h.dirIndex, pos: [2]int{x, y}, path: sUnion})
		}

		left := pmod(h.dirIndex-1, 4)
		if canVisit(left, h.pos, h.score+1000, &visited) {
			heap.Push(heapResult, &Result{score: h.score + 1000, dirIndex: left, pos: h.pos, path: h.path})
		}

		right := pmod(h.dirIndex+1, 4)
		if canVisit(right, h.pos, h.score+1000, &visited) {
			heap.Push(heapResult, &Result{score: h.score + 1000, dirIndex: right, pos: h.pos, path: h.path})
		}
	}

	return len(validPath)
}
func Resolve(data []string) [2]any {
	return [2]any{
		ResolvePart1(data),
		ResolvePart2(data),
	}
}
