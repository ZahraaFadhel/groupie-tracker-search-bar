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
					Remove(i)
					return
				} else {
					Arr[i-2] = Arr[i-2] + "'"
					Remove(i-1)
				}
			} else if strings.HasSuffix(Arr[i], "'") && !IsConverter(Arr[i][:len(Arr[i])-1]) {
				return
			} else if strings.HasSuffix(Arr[i], "'") && IsConverter(Arr[i][:len(Arr[i])-1]) {
				Arr[i-1] = Arr[i-1] + "'"
				GetConv(i,Arr)(i-1)
				Remove(i)
			}
		}
		return
	} else if Arr[index] == "'" && index+1 < Size {
		Arr[index+1] = Arr[index] + Arr[index+1]
		Remove(index)
		for i := index; i < Size; i++ {
			if Arr[i] == "'" && !IsConverter(Arr[i-1]) {
				Arr[i-1] = Arr[i-1] + Arr[i]
				Remove(i)
				break
			} else if Arr[i] == "'" && IsConverter(Arr[i-1]) {
				Arr[i-2] = Arr[i-2] + Arr[i]
				Remove(i)
			} else if strings.HasSuffix(Arr[i], "'") {
				if !IsConverter(Arr[i][:len(Arr[i])-1]) {
					break
				} else { // ' suffix to a convertor
					Arr[i-1] = Arr[i-1] + "'"
					if !IsNumConvertor(Arr[i][:len(Arr[i])-1]) {
						GetConv(i, Arr)(i-1)
						Remove(i)
					} else {
						x := GetConv2(i)
						if x == ("(hex)") {
							NumberConv(i-1, 16)
							Arr[i-1] = "'" + Arr[i-1] + "'"
						} else if x == ("(bin)") {
							NumberConv(i-1, 2)
							Arr[i-1] = "'" + Arr[i-1] + "'"
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

func GetConv2(i int) string {
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

