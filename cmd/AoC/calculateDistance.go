package main

import (
	"log"
	"slices"
)

func calculateDistance(listA, listB []int) {
	var distanceSlice []int

	slices.Sort(listA)
	slices.Sort(listB)

	totalDistance := 0

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

	log.Print(totalDistance)
}
