package main

func main() {
	filePath := "list.txt"

	fileContent := readFile(filePath)
	listA, listB := splitSlices(fileContent)
	calculateDistance(listA, listB)
}
