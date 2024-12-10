package main

func structuredElementsRework(splitString []string) []string {
	for start, end := 0, len(splitString)-1; start <= end; {
		if splitString[start] == "." && splitString[end] != "." {
			splitString[start], splitString[end] = splitString[end], splitString[start]
			start++
			end--
		} else {
			switch {
			case splitString[start] != ".":
				start++
			case splitString[end] != ".":
				end--
			case splitString[start] == "." && splitString[end] == ".":
				end--
			default:

			}
		}
	}

	return splitString
}
