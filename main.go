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
	"adventOfCode/utils"
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type ResolverFunc func([]string) [2]int

func main() {
	solutionMap := map[int]ResolverFunc{
		1: Day01.Resolve,
		2: Day02.Resolve,
		3: Day03.Resolve,
		4: Day04.Resolve,
		5: Day05.Resolve,
		6: Day06.Resolve,
		7: Day07.Resolve,
		8: Day08.Resolve,
		9: Day09.Resolve,
	}

	fmt.Println("\033[1m\033[32mAdvent of code 2024\033[0m")
	fmt.Println("List of available solutions:")

	// Récupérer les dossiers correspondant aux jours
	folders := getDirectories(".")
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
				createDay(day)
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

func createDay(day int) {
	fmt.Println("Invalid day, but creating it ...")
	dirName := fmt.Sprintf("Day%02d", day)
	err := os.Mkdir(dirName, 0750)

	_, err = os.Create(fmt.Sprintf("%s/data", dirName))

	if err != nil {
		fmt.Println("Can't creat day, stopping ...")
	}

	mainFile, err := os.Create(fmt.Sprintf("%s/main.go", dirName))

	if err != nil {
		fmt.Println("Can't creat day, stopping ...")
	} else {
		_, err := mainFile.WriteString(
			fmt.Sprintf("package %s\n%s\n%s",
				dirName,
				"import (\n\t\"adventOfCode/utils\"\n)",
				"func ResolvePart1(data []string) int {\n\treturn 0\n}\nfunc ResolvePart2(data []string) int {\n\treturn 0\n}\nfunc Resolve(data []string) [2]int {\n\treturn [2]int{\n\t\tResolvePart1(data),\n\t\tResolvePart2(data),\n\t}\n}",
			),
		)

		if err != nil {
			return
		}
	}
}

// getDirectories retourne la liste des dossiers contenant "Day"
func getDirectories(path string) []string {
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		os.Exit(1)
	}

	var folders []string
	re := regexp.MustCompile(`Day\s*(\d+)`)
	for _, file := range files {
		if file.IsDir() && re.MatchString(file.Name()) {
			folders = append(folders, file.Name())
		}
	}

	// Trier par numéro de jour
	sort.Slice(folders, func(i, j int) bool {
		num1, _ := strconv.Atoi(re.FindStringSubmatch(folders[i])[1])
		num2, _ := strconv.Atoi(re.FindStringSubmatch(folders[j])[1])
		return num1 < num2
	})

	return folders
}

func executeDay(day int, solutionMap map[int]ResolverFunc) {
	dayFolder := fmt.Sprintf("Day%02d", day)
	dataFile := filepath.Join(dayFolder, "data")
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
		createDay(day)
	}
}
