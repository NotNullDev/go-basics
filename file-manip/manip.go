package main

import (
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		panic("Usage: ./manip <filename to process>")
	}

	filename := os.Args[1]

	newFile := processFile(filename)

	println(newFile)
	os.WriteFile("itsme.txt", []byte(newFile), 0644)
}

func processFile(filename string) string {
	if strings.TrimSpace(filename) == "" {
		panic("No filename provided")
	}

	mainFile, err := os.ReadFile(filename)

	if err != nil {
		panic("Can''t read file: [" + err.Error() + "]")
	}

	mainFileString := string(mainFile)
	mainFileString = strings.Replace(mainFileString, "\r", "", -1)

	preparedListOfLines := strings.Split(mainFileString, "\n")

	for _, line := range preparedListOfLines {
		if strings.HasPrefix(line, "@") {
			lineWithoutAt := strings.ReplaceAll(line, "@", "")
			nextFileContent := processFile(lineWithoutAt)
			mainFileString = strings.Replace(mainFileString, line, nextFileContent, 1)
		}
	}
	return mainFileString
}
