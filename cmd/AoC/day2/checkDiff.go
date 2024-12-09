package main

import "log"

func checkDiff(report []int) []int {
	var result int

	for i, number := range report {
		if i > 0 {
			result = diff(number, report[i-1])
		}

		if result > 3 {
			log.Printf("Unsafe report: %d, results: %d", report, result)
			return nil
		}
	}

	return report
}

func checkPotentials(report []int) {
	var result int

	for i, number := range report {
		if i > 0 {
			result = diff(number, report[i-1])
		}

		if result > 3 {
			report = removeItem(report, number)
			log.Print(report)
		}
	}
}

func diff(first, second int) int {
	if first < second {
		return second - first
	}
	return first - second
}

func removeItem(report []int, item int) []int {
	return append(report[:item], report[item+1:]...)
}
