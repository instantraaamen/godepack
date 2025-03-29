package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Directories and files to exclude when bundling
var excludeDirs = []string{".git", "node_modules", "vendor", ".idea"}
var excludeExtensions = []string{".png", ".jpg", ".jpeg", ".gif", ".bin", ".exe"}

// Check if the path should be excluded
func isExcluded(path string) bool {
	for _, dir := range excludeDirs {
		if strings.Contains(path, dir) {
			return true
		}
	}
	for _, ext := range excludeExtensions {
		if strings.HasSuffix(path, ext) {
			return true
		}
	}
	return false
}

func main() {
	outputFile, err := os.Create("repomix-output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || isExcluded(path) {
			return nil
		}

		fileContent, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		// Output with separator
		separator := fmt.Sprintf("\n\n// ----- FILE: %s ----- //\n\n", path)
		writer.WriteString(separator)
		writer.Write(fileContent)

		return nil
	})

	writer.Flush()
	fmt.Println("Repository bundling completed: repomix-output.txt")
}
