package main

import (
	"log"
	"strconv"
)

func calculateCheckSum(structuredElements []string) int {
	checkSumTotal := 0
	for i := 0; i < len(structuredElements); i++ {
		if structuredElements[i] != "." {
			elementValue, err := strconv.Atoi(structuredElements[i])
			if err != nil {
				log.Fatal(err)
			}

			checkSumTotal += i * elementValue

		}
	}

	return checkSumTotal
}
