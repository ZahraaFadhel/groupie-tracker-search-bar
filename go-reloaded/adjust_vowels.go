package main

import (
	"strings"
)

func AdjustVowels(index int) {
	if Size == index {
		return
	}
	if Arr[index] == "a" || Arr[index] == "A" {
		if startsWithVowel(Arr[index+1]) {
			Arr[index] = Arr[index] + "n"
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