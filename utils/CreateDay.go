package utils

import (
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
)

func CreateDay(day int) {
	fmt.Println("Invalid day, but creating it ...")
	dirName := fmt.Sprintf("Day%02d", day)
	err := os.Mkdir(dirName, 0750)

	_, err = os.Create(fmt.Sprintf("%s/data", dirName))

	if err != nil {
		fmt.Println("Can't creat day, stopping ...")
		return
	}

	dayFile, err := os.Create(fmt.Sprintf("%s/main.go", dirName))

	if err != nil {
		fmt.Println("Can't creat day, stopping ...")
		return
	} else {
		_, err := dayFile.WriteString(
			fmt.Sprintf("package %s\n%s",
				dirName,
				"func ResolvePart1(data []string) int {\n\treturn 0\n}\nfunc ResolvePart2(data []string) int {\n\treturn 0\n}\nfunc Resolve(data []string) [2]int {\n\treturn [2]int{\n\t\tResolvePart1(data),\n\t\tResolvePart2(data),\n\t}\n}",
			),
		)

		if err != nil {
			return
		}

		mainFileContent := ReadFile("../main.go")

		canStopSearchImportLine := false
		bypassImport := false
		canStopSearchResolverLine := false

		lineToInsertImport := 0
		lineToInsertResolver := 0

		for line := 0; line < len(mainFileContent); line++ {
			if !bypassImport {
				containImport := strings.Contains(string(mainFileContent[line]), "adventOfCode/Day")
				if !containImport && canStopSearchImportLine {
					lineToInsertImport = line
					bypassImport = true
				}

				if containImport {
					canStopSearchImportLine = true
				}
			}

			containResolver := strings.Contains(string(mainFileContent[line]), ".Resolve")
			if !containResolver && canStopSearchResolverLine {
				lineToInsertResolver = line
				break
			}

			if containResolver {
				canStopSearchResolverLine = true
			}
		}

		mainFileContent = slices.Insert(mainFileContent, lineToInsertImport, fmt.Sprintf("\t\"adventOfCode/%s\"", dirName))
		mainFileContent = slices.Insert(mainFileContent, lineToInsertResolver+1, fmt.Sprintf("\t\t%02d: %s.Resolve,", day, dirName))

		mainFile, err := os.OpenFile("main.go", os.O_RDWR, 0644)

		if err != nil {
			fmt.Println("Can't creat day, stopping ...")
			return
		} else {
			_, err := mainFile.WriteString(
				strings.Join(mainFileContent, "\n"),
			)

			if err != nil {
				return
			}

			cmd := exec.Command("git", "add", fmt.Sprintf("%s/.", dirName))
			cmd.Output()
			cmd = exec.Command("git", "commit", "-am", fmt.Sprintf("%s: created", dirName))
			cmd.Output()
			cmd = exec.Command("git", "push")
			cmd.Output()
		}
	}
}
