package Day19

import (
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
	pos   Match
	path  []Match
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

type Match struct {
	start   int
	end     int
	pattern string
}

type Visited struct {
	match Match
	score int
}

func ResolvePart2(data []string) int {
	var patternSet []string
	design := false
	count := 0

	for i := 0; i < len(data); i++ {
		if data[i] == "" {
			design = true
			continue
		}

		if design {
			var ways func(string) int
			cache := map[string]int{}

			ways = func(design string) (n int) {
				if n, ok := cache[design]; ok {
					return n
				}
				defer func() { cache[design] = n }()

				if design == "" {
					return 1
				}

				for _, s := range patternSet {
					if strings.HasPrefix(design, s) {
						n += ways(design[len(s):])
					}
				}

				return n
			}

			count += ways(data[i])
		} else {
			patternSet = strings.Split(data[i], ", ")
		}
	}

	return count
}
func Resolve(data []string) [2]any {
	return [2]any{
		ResolvePart1(data),
		ResolvePart2(data),
	}
}
