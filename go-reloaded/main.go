package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"learn.reboot01.com/git/zfadhel/go-reloaded/convertors"
)

var (
	arr  []string
	Size int
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
	Size = len(arr)

	for i := 0; i < Size; i++ {
		if arr[i] == "(hex)" {
			convertors.NumberConv(i-1, 16, arr, Size)
		} else if arr[i] == "(bin)" {
			convertors.NumberConv(i-1, 2, arr, Size)
		} else if arr[i] == "(up)" {
			convertors.ToUpper(i - 1, arr)
			remove(i, arr, Size)
		} else if arr[i] == "(up," && strings.HasSuffix(arr[i+1], ")") {
			ans := arr[i] + " " + arr[i+1]
			num := getNum(ans)
			repeatCaseConversion(i-1, num, convertors.ToUpper, arr)
			remove(i, arr, Size)
			remove(i, arr, Size)
		} else if arr[i] == "(low)" {
			convertors.ToLower(i - 1, arr)
			remove(i, arr, Size)
		} else if arr[i] == "(low," && strings.HasSuffix(arr[i+1], ")") {
			ans := arr[i] + " " + arr[i+1]
			num := getNum(ans)
			repeatCaseConversion(i-1, num, convertors.ToLower, arr)
			remove(i, arr, Size)
			remove(i, arr, Size)
		} else if arr[i] == "(cap)" {
			convertors.Cap(i - 1, arr)
			remove(i, arr, Size)
		} else if arr[i] == "(cap," && strings.HasSuffix(arr[i+1], ")") {
			ans := arr[i] + " " + arr[i+1]
			num := getNum(ans)
			repeatCaseConversion(i-1, num, convertors.Cap, arr)
			remove(i, arr, Size)
			remove(i, arr, Size)
		}
	
		convertors.AdjustPunctuation(i, arr, Size)
		convertors.AdjustQuot(i, arr, Size)
		convertors.AdjustVowels(i, arr, Size)
	}

	result, err := os.Create(os.Args[2])
	for i := 0; i < Size; i++ {
		result.Write([]byte(arr[i] + " "))
	}
}


func removeCharAtIndex(str string, index int) string {
	runes := []rune(str)
	runes = append(runes[:index], runes[index+1:]...)
	return string(runes)
}


func repeatCaseConversion(index, count int, conversionFunc func(int, []string), y []string) {
	for i := index; i >= 0; i-- {
		if count <= 0 {
			return
		}
		conversionFunc(i, y)
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

func remove(index int, arr[]string, Size int) {
	temp := append(arr[:index], arr[index+1:]...)
	arr = temp
	Size = Size - 1
}
