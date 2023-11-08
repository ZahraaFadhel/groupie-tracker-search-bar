package convertors

import (
	"strings"
)

func AdjustVowels(index int, x[]string, size int) {
	if size == index {
		return
	}
	if x[index] == "a" || x[index] == "A" {
		if startsWithVowel(x[index+1]) {
			x[index] = x[index] + "n"
		}
	}
}

func startsWithVowel(s string) bool {
	vowels := []string{"a", "e", "i", "o", "u", "h"}
	for _, ch := range vowels {
		if strings.HasPrefix(s, ch) {
			return true
		}
	}
	return false
}