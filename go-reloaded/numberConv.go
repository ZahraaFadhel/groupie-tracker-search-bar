package main

import (
	"strings"
	"fmt"
	"strconv"
)

func NumberConv(i int, base int) {
	hex_num := Arr[i]
	if strings.HasPrefix(hex_num, "'") {
		hex_num = Arr[i][1:]
	}
	if strings.HasSuffix(hex_num, "'") {
		hex_num = hex_num[:len(Arr)-2]
	}
	num, err := strconv.ParseInt(hex_num, base, 64)
	if err != nil {
		fmt.Print("Failed")
	}
	Remove(i, Arr)
	Arr[i] = strconv.FormatInt(num, 10) // cannot use string(num)
}

func Remove(index int, Arr[]string) {
	temp := append(Arr[:index], Arr[index+1:]...)
	Arr = temp
	Size = Size - 1
}