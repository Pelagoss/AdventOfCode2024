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
				t2 := make(map[int][][]int)

				//localCount := 0
				for j := 0; j < len(patternSet); j++ {
					reBis := regexp.MustCompile(patternSet[j])
					found := reBis.FindAllStringSubmatchIndex(data[i], -1)
					//foundStr := reBis.FindAllString(data[i], -1)
					//fmt.Println(foundStr)
					if len(found) > 0 {
						for l := 0; l < len(found); l++ {
							t = append(t, found[l])
							t2[found[l][0]] = append(t2[found[l][0]], found[l])
							//localCount++
						}
					}
				}

				currentIndex := 0
				path := make(map[int][][]int)

				for currentIndex < len(data[i]) {
					if _, ok := t2[currentIndex]; !ok {
						currentIndex++
						continue
					}

					//Todo: Fix this, partir de currentIndex == 0 et créer un path par len(t2[currentIndex]) > 1 puis compter le nombres de path qui partent de 0 et finissent a len(data[i])
					for j := 0; j < len(t2[currentIndex]); j++ {
						if currentIndex == 0 {
							path[j] = append(path[j], t2[currentIndex][j])
						} else {
							for k := 0; k < len(path); k++ {
								if _, ok := path[k]; !ok {
									continue
								}

								if path[k][len(path[k])-1][1] == t2[currentIndex][j][0] {
									path[k] = append(path[k], t2[currentIndex][j])
									if len(t2[currentIndex]) > 1 {
										path[len(path)] = append(path[len(path)], path[j]...)
									}
								}
							}
						}
					}

					currentIndex++
				}
				fmt.Println(path, len(path))
				//fmt.Println(t)
				fmt.Println(t2)
				fmt.Println()
				//panic("aa")

				//for k, v := range path {
				//	fmt.Println(k, v)
				//}
				fmt.Println()
				count += len(path)
			}
		} else {
			patternSet = strings.Split(data[i], ", ")

			slices.SortFunc(patternSet, func(i, j string) int {
				return len(j) - len(i)
			})

			patterns = fmt.Sprintf("^(%s)*$", strings.Join(patternSet, "|"))
			fmt.Println(patterns)
		}
	}

	return count
}
func Resolve(data []string) [2]any {
	return [2]any{
		ResolvePart1(data),
		ResolvePart2(data), // 928 too low
	}
}
