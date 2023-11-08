package main

import (
	"fmt"
	"strconv"
	"strings"
)

func NumberConv(i int, base int) {
	hex_num := Arr[i]
	if strings.HasPrefix(hex_num, "'") {
		hex_num = Arr[i][1:]
	}
	if strings.HasSuffix(hex_num, "'") {
		hex_num = hex_num[:len(Arr)]
	}
	num, err := strconv.ParseInt(hex_num, base, 64)
	if err != nil {
		fmt.Println("Cannot be converted")
	}
	Remove(i)
	Arr[i] = strconv.FormatInt(num, 10) // cannot use string(num)
}

func Remove(index int) {
	temp := append(Arr[:index], Arr[index+1:]...)
	Arr = temp
	Size = Size - 1
}

func RepeatCaseConversion(index, count int, conversionFunc func(int)) {
	for i := index; i >= 0; i-- {
		if count <= 0 {
			return
		}
		conversionFunc(i)
		count--
	}
}

func GetNum(s string) int {
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

func ToUpper(i int) {
	Arr[i] = strings.ToUpper(Arr[i])
}

func ToLower(i int) {
	Arr[i] = strings.ToLower(Arr[i])
}

func Cap(i int) {
	Arr[i] = strings.Title(Arr[i])
}

func IsLower(s string) bool {
	runes := []rune(s)
	for i := 0; i < len(s); i++ {
		if runes[i] >= 0 && runes[i] <= 96 || runes[i] >= 123 {
			return false
		}
	}
	return true
}

func IsUpper(str string) bool {
	ArrS := []rune(str)

	for i := 0; i < len(ArrS); i++ {
		if (ArrS[i] >= 0) && (ArrS[i] <= 64) || (ArrS[i] >= 91) && (ArrS[i] <= 127) {
			return false
		}
	}
	return true
}

func IsAlpha(s string) bool {
	runes := []rune(s)
	for _, ch := range runes {
		if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch >= 48 && ch <= 57 {
			continue
		}
		return false
	}
	return true
}

func IsNumConvertor(s string) bool {
	return s == "(hex)" || s == "(bin)"
}
