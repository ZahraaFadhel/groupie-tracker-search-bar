package convertors

import (
	"strings"
)

func AdjustPunctuation(index int, arr []string, size int) {
	if index == size {
		return
	}
	runes := []rune(arr[index])
	if isPunc(arr[index]) {
		arr[index-1] = string(append([]byte(arr[index-1]), byte(runes[0])))
		remove(index, arr, size)
	} else if checkPrefixes(arr[index]) {
		// Find the index of the first non-punctuation character
		nonPuncIndex := -1
		for i := 0; i < len(runes); i++ {
			if string(runes[i]) == "'" {
				break
			}
			if !isPunc(string(runes[i])) {
				nonPuncIndex = i
				break
			}
		}

		if nonPuncIndex == -1 {
			arr[index-1] = arr[index-1] + arr[index]
			remove(index, arr, size)
		} else {
			// Add the non-punctuation characters to the previous word
			arr[index-1] = string(append([]byte(arr[index-1]), []byte(string(runes[:nonPuncIndex]))...))

			// Remove the punctuation characters from the current word
			arr[index] = string(runes[nonPuncIndex:])
		}
	}
}

func isPunc(s string) bool {
	return s == "." || s == "," || s == "!" || s == "?" || s == ":" || s == ";"
}

func remove(index int, arr []string, size int) {
	temp := append(arr[:index], arr[index+1:]...)
	arr = temp
	size = size - 1
}

func checkPrefixes(s string) bool {
	return strings.HasPrefix(s, ",") ||
		strings.HasPrefix(s, ".") ||
		strings.HasPrefix(s, ";") ||
		strings.HasPrefix(s, ":") ||
		strings.HasPrefix(s, "!") ||
		strings.HasPrefix(s, "?")
}
