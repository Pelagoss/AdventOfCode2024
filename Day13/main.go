package Day13

import (
	"fmt"
	"strconv"
	"strings"
)

type Result struct {
	a, b int
}

func getResult(Xa int, Xb int, Ya int, Yb int, Xend int, Yend int, part int) (Result, any) {
	b := ((Xa * Yend) - (Ya * Xend)) / (Xa*Yb - Ya*Xb)
	a := (Xend / Xa) - (Xb * (((Xa * Yend) - (Ya * Xend)) / (Xa*Yb - Ya*Xb)) / Xa)

	if part == 1 && (a > 100 || b > 100) || a*Xa+b*Xb != Xend || a*Ya+b*Yb != Yend {
		return Result{}, true
	}

	return Result{a: a, b: b}, nil
}

func ResolvePart1(data []string, part int) int {
	var (
		Xa, Xb, Ya, Yb, Xend, Yend int
	)

	tokens := 0

	for i := 0; i < len(data); i++ {
		if len(data[i]) == 0 {
			r, err := getResult(Xa, Xb, Ya, Yb, Xend, Yend, part)
			if err == nil {
				tokens += r.a*3 + r.b
			}

			continue
		}

		switch i % 4 {
		case 0, 1:
			buttonLetter := "B"
			if i%4 == 0 {
				buttonLetter = "A"
			}

			specsStr := strings.ReplaceAll(data[i], fmt.Sprintf("Button %v: X+", buttonLetter), "")
			specs := strings.Split(specsStr, ",")

			XStrVal := specs[0]
			YStrVal := strings.ReplaceAll(specs[1], " Y+", "")

			XVal, err := strconv.Atoi(XStrVal)
			if err != nil {
				panic(err)
			}

			YVal, err := strconv.Atoi(YStrVal)
			if err != nil {
				panic(err)
			}

			if i%4 == 0 {
				Xa = XVal
				Ya = YVal
			} else {
				Xb = XVal
				Yb = YVal
			}
			break
		case 2:
			specsStr := strings.ReplaceAll(data[i], "Prize: X=", "")
			specs := strings.Split(specsStr, ",")

			XStrVal := specs[0]
			YStrVal := strings.ReplaceAll(specs[1], " Y=", "")

			XVal, err := strconv.Atoi(XStrVal)
			if err != nil {
				panic(err)
			}

			YVal, err := strconv.Atoi(YStrVal)
			if err != nil {
				panic(err)
			}

			if part == 1 {
				Xend = XVal
				Yend = YVal
			} else {
				Xend = XVal + 10000000000000
				Yend = YVal + 10000000000000
			}
			break
		}
	}

	return tokens
}
func ResolvePart2(data []string) int {
	return ResolvePart1(data, 2)
}
func Resolve(data []string) [2]int {
	return [2]int{
		ResolvePart1(data, 1),
		ResolvePart2(data),
	}
}
