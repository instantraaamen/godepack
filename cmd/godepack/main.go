package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	ExcludeDirs       []string `json:"excludeDirs"`
	ExcludeExtensions []string `json:"excludeExtensions"`
}

func loadConfig(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	err = json.Unmarshal(data, &cfg)
	return cfg, err
}

func isExcluded(path string, info os.FileInfo, cfg Config) bool {
	if info.IsDir() {
		for _, dir := range cfg.ExcludeDirs {
			if info.Name() == dir {
				return true
			}
		}
		return false
	}

	ext := strings.ToLower(filepath.Ext(info.Name()))
	for _, excludeExt := range cfg.ExcludeExtensions {
		if ext == excludeExt {
			return true
		}
	}
	return false
}

func main() {
	cfg, err := loadConfig("config.json")
	if err != nil {
		fmt.Println("Error loading configuration file:", err)
		return
	}

	outputFile, err := os.Create("gode-packing.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)

	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Perform exclusion check
		if isExcluded(path, info, cfg) {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		if info.IsDir() {
			return nil
		}

		fileContent, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		separator := fmt.Sprintf("\n\n// ----- FILE: %s ----- //\n\n", path)
		writer.WriteString(separator)
		writer.Write(fileContent)

		return nil
	})

	if err != nil {
		fmt.Println("An error occurred during processing:", err)
		return
	}

	writer.Flush()
	fmt.Println("Repository bundling completed: gode-packing.txt")
}
