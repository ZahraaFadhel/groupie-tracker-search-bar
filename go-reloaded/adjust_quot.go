package main

import (
	"strings"
)

func AdjustQuot(index int) {
	if Size == index {
		return
	}
	if strings.HasPrefix(Arr[index], "'") && len(Arr[index]) > 1 {
		for i := index; i < Size; i++ {
			if Arr[i] == "'" && Size-1 > index+1 {
				if !IsConverter(Arr[i-1][:len(Arr[i-1])-1]) {
					Arr[i-1] = Arr[i-1] + Arr[i]
					Remove(i, Arr)
					return
				} else {
					Arr[i-2] = Arr[i-2] + "'"
					Remove(i-1, Arr)
				}
			} else if strings.HasSuffix(Arr[i], "'") && !IsConverter(Arr[i][:len(Arr[i])-1]) {
				return
			} else if strings.HasSuffix(Arr[i], "'") && IsConverter(Arr[i][:len(Arr[i])-1]) {
				Arr[i-1] = Arr[i-1] + "'"
				GetConv(i,Arr)(i-1)
				Remove(i, Arr)
			}
		}
		return
	} else if Arr[index] == "'" && index+1 < Size {
		Arr[index+1] = Arr[index] + Arr[index+1]
		Remove(index, Arr)
		for i := index; i < Size; i++ {
			if Arr[i] == "'" && !IsConverter(Arr[i-1]) {
				Arr[i-1] = Arr[i-1] + Arr[i]
				Remove(i, Arr)
				break
			} else if Arr[i] == "'" && IsConverter(Arr[i-1]) {
				Arr[i-2] = Arr[i-2] + Arr[i]
				Remove(i, Arr)
			} else if strings.HasSuffix(Arr[i], "'") {
				if !IsConverter(Arr[i][:len(Arr[i])-1]) {
					break
				} else { // ' suffix to a convertor
					Arr[i-1] = Arr[i-1] + "'"
					if !IsNumConvertor(Arr[i]) {
						GetConv(i, Arr)(i-1)
						Remove(i, Arr)
					} else {
						x := GetConv2(i, Arr)
						if x == ("hex") {
							NumberConv(i, 16)
						} else if x == ("bin") {
							NumberConv(i, 2)
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

func GetConv2(i int, Arr[]string) string {
	if Arr[i][:len(Arr[i])-1] == "(hex)" {
		return "(hex)"
	} else {
		return "(bin)"
	}
}

func GetConv(i int, Arr[]string) func(int) {
	if Arr[i][:len(Arr[i])-1] == "(up)" {
		return ToUpper
	} else if Arr[i][:len(Arr[i])-1] == "(low)" {
		return ToLower
	} else if Arr[i][:len(Arr[i])-1] == "(cap)" {
		return Cap
	} else {
		return nil
	}
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