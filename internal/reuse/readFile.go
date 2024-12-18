package reuse

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(filePath string) []string {
	var fileContentSlice []string
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileContent := bufio.NewScanner(file)

	for fileContent.Scan() {
		fileContentSlice = append(fileContentSlice, fileContent.Text())
	}

	return fileContentSlice
}
