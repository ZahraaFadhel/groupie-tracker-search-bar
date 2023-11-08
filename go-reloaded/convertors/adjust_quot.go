package convertors

import (
	"fmt"
	"strconv"
	"strings"
)

func AdjustQuot(index int, arr []string, size int) {
	if size == index {
		return
	}
	if strings.HasPrefix(arr[index], "'") && len(arr[index]) > 1 {
		for i := index; i < size; i++ {
			if arr[i] == "'" && size-1 > index+1 {
				if !IsConverter(arr[i-1][:len(arr[i-1])-1]) {
					arr[i-1] = arr[i-1] + arr[i]
					remove(i, arr, size)
					return
				} else {
					arr[i-2] = arr[i-2] + "'"
					remove(i-1, arr, size)
				}
			} else if strings.HasSuffix(arr[i], "'") && !IsConverter(arr[i][:len(arr[i])-1]) {
				return
			} else if strings.HasSuffix(arr[i], "'") && IsConverter(arr[i][:len(arr[i])-1]) {
				arr[i-1] = arr[i-1] + "'"
				GetConv(i,arr)(i-1, arr)
				remove(i, arr, size)
			}
		}
		return
	} else if arr[index] == "'" && index+1 < size {
		arr[index+1] = arr[index] + arr[index+1]
		remove(index, arr, size)
		for i := index; i < size; i++ {
			if arr[i] == "'" && !IsConverter(arr[i-1]) {
				arr[i-1] = arr[i-1] + arr[i]
				remove(i, arr, size)
				break
			} else if arr[i] == "'" && IsConverter(arr[i-1]) {
				arr[i-2] = arr[i-2] + arr[i]
				remove(i, arr, size)
			} else if strings.HasSuffix(arr[i], "'") {
				if !IsConverter(arr[i][:len(arr[i])-1]) {
					break
				} else { // ' suffix to a convertor
					arr[i-1] = arr[i-1] + "'"
					if !IsNumConvertor(arr[i]) {
						GetConv(i, arr)(i-1, arr)
						remove(i, arr, size)
					} else {
						x := GetConv2(i, arr)
						if x == ("hex") {
							NumberConv(i, 16, arr,size)
						} else if x == ("bin") {
							NumberConv(i, 2, arr, size)
						}
					}
				}
			}
		}
	}
}

func IsConverter(s string) bool {
	return s == "(cap)" || s == "(up)" || s == "(low)" || s == "(hex)" || s == "(bin)"
}

func NumberConv(i int, base int, arr []string, size int) {
	hex_num := arr[i]
	if strings.HasPrefix(hex_num, "'") {
		hex_num = arr[i][1:]
	}
	if strings.HasSuffix(hex_num, "'") {
		hex_num = hex_num[:len(arr)]
	}
	num, err := strconv.ParseInt(hex_num, base, 64)
	if err != nil {
		fmt.Print("Failed")
	}
	remove(i, arr, size)
	arr[i] = strconv.FormatInt(num, 10) // cannot use string(num)
}

func GetConv2(i int, arr[]string) string {
	if arr[i][:len(arr[i])-1] == "(hex)" {
		return "(hex)"
	} else {
		return "(bin)"
	}
}

func GetConv(i int, arr[]string) func(int, []string) {
	if arr[i][:len(arr[i])-1] == "(up)" {
		return ToUpper
	} else if arr[i][:len(arr[i])-1] == "(low)" {
		return ToLower
	} else if arr[i][:len(arr[i])-1] == "(cap)" {
		return Cap
	} else {
		return nil
	}
}

func ToUpper(i int, arr[]string) {
	arr[i] = strings.ToUpper(arr[i])
}

func ToLower(i int, arr[]string) {
	arr[i] = strings.ToLower(arr[i])
}

func Cap(i int, arr[]string) {
	arr[i] = strings.Title(arr[i])
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
	arrS := []rune(str)

	for i := 0; i < len(arrS); i++ {
		if (arrS[i] >= 0) && (arrS[i] <= 64) || (arrS[i] >= 91) && (arrS[i] <= 127) {
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