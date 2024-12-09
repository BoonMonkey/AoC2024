package main

import (
	"slices"
)

func calculateDistance(listA, listB []int) int {
	var distanceSlice []int
	totalDistance := 0

	slices.Sort(listA)
	slices.Sort(listB)

	for i := 0; i < len(listA); i++ {
		distance := listB[i] - listA[i]
		distanceSlice = append(distanceSlice, distance)
	}

	for _, distance := range distanceSlice {
		if distance > 0 {
			totalDistance = totalDistance + distance
		} else {
			totalDistance = totalDistance - distance
		}
	}

	return totalDistance
}

func calculateScore(listA, listB []int) int {
	var similarityScore []int
	totalScore := 0
	similarity := make(map[int]int)

	slices.Sort(listA)
	slices.Sort(listB)

	for i := 0; i < len(listA); i++ {
		matches := 0
		for _, item := range listB {
			if listA[i] == item {
				matches++
			}
		}

		if matches > 0 {
			similarity[listA[i]] = matches
		}
	}

	for k, v := range similarity {
		score := k * v
		similarityScore = append(similarityScore, score)
	}

	for _, score := range similarityScore {
		if score > 0 {
			totalScore = totalScore + score
		} else {
			totalScore = totalScore - score
		}
	}

	return totalScore
}
