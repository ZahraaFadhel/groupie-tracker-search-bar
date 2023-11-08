package main

import (
	"strings"
)

func AdjustPunctuation(index int) {
	if index == Size {
		return
	}
	runes := []rune(Arr[index])
	if isPunc(Arr[index]) {
		Arr[index-1] = string(append([]byte(Arr[index-1]), byte(runes[0])))
		Remove(index)
	} else if checkPrefixes(Arr[index]) {
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
			Arr[index-1] = Arr[index-1] + Arr[index]
			Remove(index)
		} else {
			// Add the non-punctuation characters to the previous word
			Arr[index-1] = string(append([]byte(Arr[index-1]), []byte(string(runes[:nonPuncIndex]))...))

			// Remove the punctuation characters from the current word
			Arr[index] = string(runes[nonPuncIndex:])
		}
	}
}

func isPunc(s string) bool {
	return s == "." || s == "," || s == "!" || s == "?" || s == ":" || s == ";"
}

func checkPrefixes(s string) bool {
	return strings.HasPrefix(s, ",") ||
		strings.HasPrefix(s, ".") ||
		strings.HasPrefix(s, ";") ||
		strings.HasPrefix(s, ":") ||
		strings.HasPrefix(s, "!") ||
		strings.HasPrefix(s, "?")
}
