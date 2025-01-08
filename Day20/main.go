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

func visitRaceTrack(raceTrack RaceTrack, historians Historian, path [][2]int, cheatLength int) ([][2]int, map[int]int, map[[2][2]int]int) {
	heapResult := &ResultHeap{}
	heap.Init(heapResult)
	heap.Push(heapResult, &Result{0, historians.pos, make([][2]int, 0)})
	visited := make(map[[2]int]bool)
	cheats := make(map[int]int)
	cheatsMap := make(map[[2][2]int]int)

	for heapResult.Len() > 0 {
		h := heap.Pop(heapResult).(*Result)

		if h.pos == raceTrack.end {
			return append(h.path, h.pos), cheats, cheatsMap
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
				heapCheat := &ResultHeap{}
				heap.Init(heapCheat)
				heap.Push(heapCheat, &Result{0, h.pos, make([][2]int, 0)})

				visitedCheat := make(map[[2]int]bool)
				for heapCheat.Len() > 0 {
					h2 := heap.Pop(heapCheat).(*Result)

					index := slices.Index(path, h2.pos)

					if index != -1 { // Cheat exists
						picosecondsSaved := index - (h.score + h2.score)
						if picosecondsSaved >= 100 {
							_, ok := cheatsMap[[2][2]int{h.pos, h2.pos}]

							if !ok || cheatsMap[[2][2]int{h.pos, h2.pos}] > picosecondsSaved {
								cheatsMap[[2][2]int{h.pos, h2.pos}] = picosecondsSaved
							}
						}
					}

					if _, ok := visitedCheat[h2.pos]; ok {
						continue
					}

					visitedCheat[h2.pos] = true

					for i2 := 0; i2 < len(nextDirections); i2++ {
						x2 := h2.pos[0] + nextDirections[i2].vecteur[0]
						y2 := h2.pos[1] + nextDirections[i2].vecteur[1]

						isNotInMap2 := x2 < 0 || y2 < 0 || x2 >= len(raceTrack.carte[0]) || y2 >= len(raceTrack.carte)

						if _, ok := visitedCheat[[2]int{x2, y2}]; !ok && !isNotInMap2 && h2.score < cheatLength {
							heap.Push(heapCheat, &Result{score: h2.score + 1, pos: [2]int{x2, y2}, path: append(h2.path, h2.pos)})
						}
					}
				}
			}

			if _, ok := visited[[2]int{x, y}]; !ok && !isNotInMap && raceTrack.carte[y][x] == "." {
				heap.Push(heapResult, &Result{score: h.score + 1, pos: [2]int{x, y}, path: append(h.path, h.pos)})
			}
		}
	}

	return make([][2]int, 0), cheats, cheatsMap
}

func ResolvePart1(data []string) int {
	racetrack, historian := buildMap(data)

	path, _, _ := visitRaceTrack(racetrack, historian, nil, 2)

	_, _, cheatsMap := visitRaceTrack(racetrack, historian, path, 2)
	cheatsAtLeast100 := 0
	for k, _ := range cheatsMap {
		if cheatsMap[k] >= 100 {
			cheatsAtLeast100++
		}
	}
	return cheatsAtLeast100
}
func ResolvePart2(data []string) int {
	racetrack, historian := buildMap(data)

	path, _, _ := visitRaceTrack(racetrack, historian, nil, 2)

	_, _, cheatsMap := visitRaceTrack(racetrack, historian, path, 20)
	cheatsAtLeast100 := 0
	for k, _ := range cheatsMap {
		if cheatsMap[k] >= 100 {
			cheatsAtLeast100++
		}
	}
	return cheatsAtLeast100
}
func Resolve(data []string) [2]any {
	return [2]any{
		ResolvePart1(data),
		ResolvePart2(data),
	}
}
