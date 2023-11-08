package main

import (
	"fmt"
	"strconv"
	"strings"
)

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
