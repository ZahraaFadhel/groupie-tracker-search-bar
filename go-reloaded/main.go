package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	Arr  []string
	Size int
)

func main() {
	if len(os.Args) != 3 {
		fmt.Print("Incorrect Arguments")
		return
	}

	// Read input file & convert content to Arr
	content, err := os.ReadFile(os.Args[1]) // content in []byte
	if err != nil {
		fmt.Print("File not Found")
		return
	}
	contentStr := string(content)
	Arr = strings.Fields(contentStr) // split string based on spaces, \t and \n
	Size = len(Arr)

	for i := 0; i < Size; i++ {
		if Arr[i] == "(hex)" {
			NumberConv(i-1, 16)
		} else if Arr[i] == "(bin)" {
			NumberConv(i-1, 2)
		} else if Arr[i] == "(up)" {
			ToUpper(i - 1)
			Remove(i)
		} else if Arr[i] == "(up," && strings.HasSuffix(Arr[i+1], ")") {
			ans := Arr[i] + " " + Arr[i+1]
			num := GetNum(ans)
			RepeatCaseConversion(i-1, num, ToUpper)
			Remove(i)
			Remove(i)
		} else if Arr[i] == "(low)" {
			ToLower(i - 1)
			Remove(i)
		} else if Arr[i] == "(low," && strings.HasSuffix(Arr[i+1], ")") {
			ans := Arr[i] + " " + Arr[i+1]
			num := GetNum(ans)
			RepeatCaseConversion(i-1, num, ToLower)
			Remove(i)
			Remove(i)
		} else if Arr[i] == "(cap)" {
			Cap(i - 1)
			Remove(i)
		} else if Arr[i] == "(cap," && strings.HasSuffix(Arr[i+1], ")") {
			ans := Arr[i] + " " + Arr[i+1]
			num := GetNum(ans)
			RepeatCaseConversion(i-1, num, Cap)
			Remove(i)
			Remove(i)
		}
		AdjustPunctuation(i)
		AdjustQuot(i)
		AdjustVowels(i)
	}

	result, err := os.Create(os.Args[2])
	for i := 0; i < Size; i++ {
		result.Write([]byte(Arr[i] + " "))
	}
}

