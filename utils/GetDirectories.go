package utils

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

// GetDirectories retourne la liste des dossiers contenant "Day"
func GetDirectories(path string) []string {
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
