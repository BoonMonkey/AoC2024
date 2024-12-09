package main

import (
	"log"
	"strconv"
	"strings"
)

func checkOrderIncreasingRework(fileContent []string) [][]int {
	var safeOutput [][]int
	for _, line := range fileContent {
		var output []int
		log.Print(line)
		items := strings.Split(line, " ")

		for i, item := range items {
			number, err := strconv.Atoi(item)
			if err != nil {
				log.Fatal(err)
			}

			if number < number+i {
				output = append(output, number)
			}
			safeOutput = append(safeOutput, output)
		}
	}
	return safeOutput
}

func checkOrderDecreasingRework(fileContent []string) [][]int {
	var safeOutput [][]int
	for _, line := range fileContent {
		var output []int
		log.Print(line)
		items := strings.Split(line, " ")

		for i, item := range items {
			number, err := strconv.Atoi(item)
			if err != nil {
				log.Fatal(err)
			}

			if number > number+i {
				output = append(output, number)
			}
			safeOutput = append(safeOutput, output)
		}
	}
	return safeOutput
}
