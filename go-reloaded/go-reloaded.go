package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	arr []string
	size int
)

func main() {
	if len(os.Args) != 3 {
		fmt.Print("Incorrect Arguments")
		return
	}

	// Read input file & convert content to arr
	content, err := os.ReadFile(os.Args[1]) // content in []byte
	if err != nil {
		fmt.Print("File not Found")
		return
	}
	contentStr := string(content)
	arr = strings.Fields(contentStr) // split string based on spaces, \t and \n
	size = len(arr)

	for i := 0; i < size ; i++ {
		if arr[i] == "(hex)" {
			NumberConv(i-1, 16)
		} else if arr[i] == "(bin)" {
			NumberConv(i-1, 2)
		} else if arr[i] == "(up)" {
			toUpper(i - 1)
			remove(i)
		} else if arr[i] == "(up," && strings.HasSuffix(arr[i+1], ")") {
			ans := arr[i] + " " + arr[i+1]
			num := getNum(ans)
			repeatCaseConversion(i-1, num, toUpper)
			remove(i)
			remove(i)
		} else if arr[i] == "(low)" {
			toLower(i - 1)
			remove(i)
		} else if arr[i] == "(low," && strings.HasSuffix(arr[i+1], ")") {
			ans := arr[i] + " " + arr[i+1]
			num := getNum(ans)
			repeatCaseConversion(i-1, num, toLower)
			remove(i)
			remove(i)
		} else if arr[i] == "(cap)" {
			cap(i - 1)
			remove(i)
		} else if arr[i] == "(cap," && strings.HasSuffix(arr[i+1], ")") {
			ans := arr[i] + " " + arr[i+1]
			num := getNum(ans)
			repeatCaseConversion(i-1, num, cap)
			remove(i)
			remove(i)
		}
			adjustPunctuation(i)
			AdjustQuot(i)
			AdjustVowels(i)
	}

	result, err := os.Create(os.Args[2])
	for i := 0; i < size; i++ {
		result.Write([]byte(arr[i] + " "))
	}

}

func adjustPunctuation(index int) {
	runes := []rune(arr[index])
	if isPunc(arr[index]) {
		arr[index-1] = string(append([]byte(arr[index-1]), byte(runes[0])))
		remove(index)
	} else if checkPrefixes(arr[index]) {
		// Find the index of the first non-punctuation character
		nonPuncIndex := 0
		for i := 0; i < len(runes); i++ {
			if !isPunc(string(runes[i])) {
				nonPuncIndex = i
				break
			}
		}

		if nonPuncIndex == 0 {
			arr[index-1] = arr[index-1] + arr[index]
			remove(index)
		} else {
			// Add the non-punctuation characters to the previous word
			arr[index-1] = string(append([]byte(arr[index-1]), []byte(string(runes[:nonPuncIndex]))...))

			// Remove the punctuation characters from the current word
			arr[index] = string(runes[nonPuncIndex:])
		}
	}
}

func AdjustVowels(index int) {
	if size == index {
		return
	}
	if arr[index] == "a" || arr[index] == "A" {
		if startsWithVowel(arr[index+1]) {
			arr[index] = arr[index] + "n"
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

func AdjustQuot(index int) {
	if size == index {
		return
	}
	if strings.HasPrefix(arr[index], "'") && len(arr[index]) > 1 {
		for i := index; i < size; i++ {
			if arr[i] == "'" && size-1 > index+1 {
				arr[i-1] = arr[i-1] + arr[i]
				remove(i)
				return
			} else if strings.HasSuffix(arr[i], "'") {
				return
			}
		}
		return
	} else if arr[index] == "'" && index < size {
		arr[index+1] = arr[index] + arr[index+1]
		remove(index)
		for i := index + 1; i < size; i++ {
			if arr[i] == "'" {
				arr[i-1] = arr[i-1] + arr[i]
				remove(i)
				break
			} else if strings.HasSuffix(arr[i], "'") {
				break
			}
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

func removeCharAtIndex(str string, index int) string {
	runes := []rune(str)
	runes = append(runes[:index], runes[index+1:]...)
	return string(runes)
}

func remove(index int) {
	temp := append(arr[:index], arr[index+1:]...)
	arr = temp
	size = size-1
}

func NumberConv(i int, base int) {
	hex_num := arr[i]
	num, err := strconv.ParseInt(hex_num, base, 64)
	if err != nil {
		fmt.Print("Failed")
	}
	remove(i)
	arr[i] = strconv.FormatInt(num, 10) // cannot use string(num)
}

func toUpper(i int) {
	// runes := []rune(arr[i])
	// answer := ""
	// for _, ch := range runes {
	// 	if isLower(string(ch)) {
	// 		answer += string(ch - 32)
	// 	} else {
	// 		answer += string(ch)
	// 	}
	// }
	// remove(i)
	// arr[i] = answer
	arr[i] = strings.ToUpper(arr[i])
}

func toLower(i int) {
	// runes := []rune(arr[i])
	// answer := ""
	// for _, ch := range runes {
	// 	if isUpper(string(ch)) {
	// 		answer += string(ch + 32)
	// 	} else {
	// 		answer += string(ch)
	// 	}
	// }
	// remove(i)
	arr[i] = strings.ToLower(arr[i])
}

func cap(i int) {
	// runes := []rune(arr[i])
	// isWordStart := true
	// for i := 0; i < len(runes); i++ {
	// 	if !isAlpha(string(runes[i])) {
	// 		isWordStart = true
	// 	} else if isWordStart {
	// 		if runes[i] >= 'a' && runes[i] <= 'z' {
	// 			runes[i] = runes[i] - 32
	// 		}
	// 		isWordStart = false
	// 	} else if isUpper(string(runes[i])) {
	// 		runes[i] = runes[i] + 32
	// 	}
	// }
	// remove(i)
	// arr[i] = string(runes)
	arr[i] = strings.Title(arr[i])
}

func isLower(s string) bool {
	arr := []rune(s)
	for i := 0; i < len(s); i++ {
		if arr[i] >= 0 && arr[i] <= 96 || arr[i] >= 123 {
			return false
		}
	}
	return true
}

func isUpper(str string) bool {
	arrS := []rune(str)

	for i := 0; i < len(arrS); i++ {
		if (arrS[i] >= 0) && (arrS[i] <= 64) || (arrS[i] >= 91) && (arrS[i] <= 127) {
			return false
		}
	}
	return true
}

func isAlpha(s string) bool {
	arr := []rune(s)
	for _, ch := range arr {
		if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch >= 48 && ch <= 57 {
			continue
		}
		return false
	}
	return true
}

func repeatCaseConversion(index, count int, conversionFunc func(int)) {
	for i := index; i >= 0; i-- {
		if count <= 0 {
			return
		}
		conversionFunc(i)
		count--
	}
}

func getNum(s string) int {
	startIndex := strings.Index(s, ",")
	endIndex := strings.Index(s, ")")
	numStr := s[startIndex+2 : endIndex]
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	return num
}
