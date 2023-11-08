// package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// var (
// 	arr  []string
// 	size int
// )

// func main() {
// 	if len(os.Args) != 3 {
// 		fmt.Print("Incorrect Arguments")
// 		return
// 	}

// 	// Read input file & convert content to arr
// 	content, err := os.ReadFile(os.Args[1]) // content in []byte
// 	if err != nil {
// 		fmt.Print("File not Found")
// 		return
// 	}
// 	contentStr := string(content)
// 	arr = strings.Fields(contentStr) // split string based on spaces, \t and \n
// 	size = len(arr)

// 	for i := 0; i < size; i++ {
// 		if arr[i] == "(hex)" {
// 			convertors.NumberConv(i-1, 16, size)
// 		} else if arr[i] == "(bin)" {
// 			convertors.NumberConv(i-1, 2, size)
// 		} else if arr[i] == "(up)" {
// 			convertors.toUpper(i - 1, arr)
// 			remove(i, arr, size)
// 		} else if arr[i] == "(up," && strings.HasSuffix(arr[i+1], ")") {
// 			ans := arr[i] + " " + arr[i+1]
// 			num := getNum(ans)
// 			repeatCaseConversion(i-1, num, convertors.toUpper, arr)
// 			remove(i, arr, size)
// 			remove(i, arr, size)
// 		} else if arr[i] == "(low)" {
// 			convertors.toLower(i - 1, arr)
// 			remove(i, arr, size)
// 		} else if arr[i] == "(low," && strings.HasSuffix(arr[i+1], ")") {
// 			ans := arr[i] + " " + arr[i+1]
// 			num := getNum(ans)
// 			repeatCaseConversion(i-1, num, convertors.toLower, arr)
// 			remove(i, arr, size)
// 			remove(i, arr, size)
// 		} else if arr[i] == "(cap)" {
// 			convertors.cap(i - 1, arr)
// 			remove(i, arr, size)
// 		} else if arr[i] == "(cap," && strings.HasSuffix(arr[i+1], ")") {
// 			ans := arr[i] + " " + arr[i+1]
// 			num := getNum(ans)
// 			repeatCaseConversion(i-1, num, convertors.cap, arr)
// 			remove(i, arr, size)
// 			remove(i, arr, size)
// 		}
	
// 		convertors.adjustPunctuation(i, arr, size)
// 		convertors.AdjustQuot(i, arr, size)
// 		convertors.AdjustVowels(i, arr, size)
// 	}

// 	result, err := os.Create(os.Args[2])
// 	for i := 0; i < size; i++ {
// 		result.Write([]byte(arr[i] + " "))
// 	}
// }


// func removeCharAtIndex(str string, index int) string {
// 	runes := []rune(str)
// 	runes = append(runes[:index], runes[index+1:]...)
// 	return string(runes)
// }


// func repeatCaseConversion(index, count int, conversionFunc func(int, []string), y []string) {
// 	for i := index; i >= 0; i-- {
// 		if count <= 0 {
// 			return
// 		}
// 		conversionFunc(i, y)
// 		count--
// 	}
// }

// func getNum(s string) int {
// 	startIndex := strings.Index(s, ",")
// 	endIndex := strings.Index(s, ")")
// 	numStr := s[startIndex+2 : endIndex]
// 	num, err := strconv.Atoi(numStr)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return -1
// 	}
// 	return num
// }

// func remove(index int, arr[]string, size int) {
// 	temp := append(arr[:index], arr[index+1:]...)
// 	arr = temp
// 	size = size - 1
// }
