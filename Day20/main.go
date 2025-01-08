package Day20

import (
	"container/heap"
	"slices"
)

type RaceTrack struct {
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

func buildMap(data []string) (RaceTrack, Historian) {
	var carte [][]string
	raceTrack := RaceTrack{}
	historian := Historian{}

	for y := 0; y < len(data); y++ {
		var line []string
		for x := 0; x < len(data[y]); x++ {
			letter := string(data[y][x])
			switch string(data[y][x]) {
			case "E":
				raceTrack.end = [2]int{x, y}
				letter = "."
				break
			case "S":
				historian = Historian{pos: [2]int{x, y}}
				raceTrack.start = [2]int{x, y}
				letter = "."
				break
			}

			line = append(line, letter)
		}
		carte = append(carte, line)
	}

	raceTrack.carte = carte

	return raceTrack, historian
}

func visitRaceTrack(raceTrack RaceTrack, historians Historian, path [][2]int) (bool, [][2]int, int, map[int]int, int) {
	heapResult := &ResultHeap{}
	heap.Init(heapResult)
	heap.Push(heapResult, &Result{0, historians.pos, make([][2]int, 0)})
	visited := make(map[[2]int]bool)
	cheats := make(map[int]int)
	cheatsAtLeast100 := 0

	for heapResult.Len() > 0 {
		h := heap.Pop(heapResult).(*Result)

		if h.pos == raceTrack.end {
			return true, append(h.path, h.pos), h.score, cheats, cheatsAtLeast100
		}

		if _, ok := visited[[2]int{h.pos[0], h.pos[1]}]; ok {
			continue
		}

		visited[[2]int{h.pos[0], h.pos[1]}] = true

		for i := 0; i < len(nextDirections); i++ {
			x := h.pos[0] + nextDirections[i].vecteur[0]
			y := h.pos[1] + nextDirections[i].vecteur[1]

			isNotInMap := x < 0 || y < 0 || x >= len(raceTrack.carte[0]) || y >= len(raceTrack.carte)

			if path != nil {
				x2 := h.pos[0] + (nextDirections[i].vecteur[0] * 2)
				y2 := h.pos[1] + (nextDirections[i].vecteur[1] * 2)

				index := slices.Index(path, [2]int{x2, y2})

				if index != -1 { // Cheat exists
					picosecondsSaved := index - (h.score + 2)

					if picosecondsSaved > 0 {
						cheats[picosecondsSaved]++

						if picosecondsSaved >= 100 {
							cheatsAtLeast100++
						}
					}
				}
			}

			if _, ok := visited[[2]int{x, y}]; !ok && !isNotInMap && raceTrack.carte[y][x] == "." {
				heap.Push(heapResult, &Result{score: h.score + 1, pos: [2]int{x, y}, path: append(h.path, h.pos)})
			}
		}
	}

	return false, make([][2]int, 0), 0, cheats, cheatsAtLeast100
}

func ResolvePart1(data []string) int {
	racetrack, historian := buildMap(data)

	_, path, _, _, _ := visitRaceTrack(racetrack, historian, nil)

	_, _, _, _, cheatsAtLeast100 := visitRaceTrack(racetrack, historian, path)

	return cheatsAtLeast100
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
