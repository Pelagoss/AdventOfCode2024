package main

import (
	"adventOfCode/Day01"
	"adventOfCode/Day02"
	"adventOfCode/Day03"
	"adventOfCode/Day04"
	"adventOfCode/Day05"
	"adventOfCode/Day06"
	"adventOfCode/Day07"
	"adventOfCode/Day08"
	"adventOfCode/Day09"
	"adventOfCode/Day10"
	"adventOfCode/Day11"
	"adventOfCode/Day12"
	"adventOfCode/Day13"
	"adventOfCode/Day14"
	"adventOfCode/Day15"
	"adventOfCode/Day16"
	"adventOfCode/Day17"
	"adventOfCode/Day18"
	"adventOfCode/Day19"
	"adventOfCode/utils"
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type ResolverFunc func([]string) [2]int

func main() {
	solutionMap := map[int]ResolverFunc{
		1:  Day01.Resolve,
		2:  Day02.Resolve,
		3:  Day03.Resolve,
		4:  Day04.Resolve,
		5:  Day05.Resolve,
		6:  Day06.Resolve,
		7:  Day07.Resolve,
		8:  Day08.Resolve,
		9:  Day09.Resolve,
		10: Day10.Resolve,
		11: Day11.Resolve,
		12: Day12.Resolve,
		13: Day13.Resolve,
		14: Day14.Resolve,
		15: Day15.Resolve,
		16: Day16.Resolve,
		17: Day17.Resolve,
		18: Day18.Resolve,
		19: Day19.Resolve,
	}

	fmt.Println("\033[1m\033[32mAdvent of code 2024\033[0m")
	fmt.Println("List of available solutions:")

	// Récupérer les dossiers correspondant aux jours
	folders := utils.GetDirectories(".")
	currentDate := time.Now()
	isDuringAdvent := currentDate.Month() == time.December && currentDate.Day() <= 25

	// Afficher les dossiers avec mise en surbrillance du jour actuel si applicable
	for i, folder := range folders {
		dayNumber := i + 1
		highlight := dayNumber == currentDate.Day() && isDuringAdvent
		colorText := "\033[33m" // Jaune par défaut
		if highlight {
			colorText = "\033[32m" // Vert si c'est le jour actuel
		}
		fmt.Printf("%s%2d \033[0m- %s\n", colorText, dayNumber, folder)
	}

	if isDuringAdvent {
		fmt.Println("\033[33mall\033[0m - Run all solutions")
	} else {
		fmt.Println("\033[32mall - Run all solutions\033[0m")
	}

	// Lecture de l'entrée utilisateur
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\nWhich day's solution do you want to see? \033[32m[%d]\033[0m ", currentDate.Day())
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "" {
		if isDuringAdvent {
			input = strconv.Itoa(currentDate.Day())
		} else {
			input = "all"
		}
	}

	// Vérification et exécution
	if input != "all" {
		day, err := strconv.Atoi(input)
		if err != nil || day < 1 || day > len(folders) {
			if day == len(folders)+1 {
				utils.CreateDay(day)
			} else {
				fmt.Println("Invalid day, stopping ...")
			}
			return
		}
		executeDay(day, solutionMap)
	} else {
		for day := 1; day <= len(folders); day++ {
			executeDay(day, solutionMap)
		}
	}
}

func executeDay(day int, solutionMap map[int]ResolverFunc) {
	dayFolder := fmt.Sprintf("Day%02d", day)
	dataFile := filepath.Join(dayFolder, "test")
	data := utils.ReadFile(dataFile)

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	table.DefaultHeaderFormatter = func(format string, vals ...interface{}) string {
		return strings.ToUpper(fmt.Sprintf(format, vals...))
	}

	tbl := table.New(fmt.Sprintf("Day%02d", day), "Part", "Value")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	if solutionMap[day] != nil {
		for i, value := range solutionMap[day](data) {
			tbl.AddRow("", i+1, value)
		}

		tbl.Print()
	} else {
		utils.CreateDay(day)
	}
}
