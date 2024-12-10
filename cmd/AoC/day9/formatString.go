package main

import (
	"log"
	"strconv"
	"strings"
)

func formatString(fileContent []string) string {
	characters := strings.Split(fileContent[0], "")
	translatedCharacters := []string{}
	evenCount := 0
	oddCount := 0

	for i, char := range characters {
		if i%2 == 0 {
			// fmt.Println(char + "is even positioned")
			timesToPrint, err := strconv.Atoi(char)
			characterToPrint := strconv.Itoa(evenCount)
			if err != nil {
				log.Fatal(err)
			}
			characterToAdd := strings.Repeat(characterToPrint+"_split", timesToPrint) // SUUUUUPER JANK but fixes my bug
			translatedCharacters = append(translatedCharacters, characterToAdd)
			evenCount++
		} else {
			// fmt.Println(char + "is odd positioned")
			timesToPrint, err := strconv.Atoi(char)
			characterToPrint := "."
			if err != nil {
				log.Fatal(err)
			}
			characterToAdd := strings.Repeat(characterToPrint+"_split", timesToPrint) // SUUUUUPER JANK but fixes my bug
			translatedCharacters = append(translatedCharacters, characterToAdd)
			oddCount++
		}
	}

	joinedCharacters := strings.Join(translatedCharacters, "")
	return joinedCharacters
}
