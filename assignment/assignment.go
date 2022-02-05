package assignment

import (
	"math"
	"sort"
	"strings"
)

func AddUint32(x, y uint32) (uint32, bool) {
	sum := uint64(x) + uint64(y)
	oweflow := sum >= math.MaxUint32
	return uint32(sum), oweflow
}

func CeilNumber(f float64) float64 {
	intPart := int64(f)
	ceil := f - float64(intPart)
	result := float64(intPart)

	if ceil > 0.75 {
		result += 1
	} else if ceil > 0.5 {
		result += 0.75
	} else if ceil > 0.25 {
		result += 0.5
	} else if ceil > 0 {
		result += 0.25
	}

	return result
}

func AlphabetSoup(s string) string {
	chars := sort.StringSlice(strings.Split(s, ""))
	sort.Sort(chars)
	return strings.Join(chars, "")
}

func StringMask(s string, n uint) string {
	strLen := len(s)
	if n == 0 {
		return stringMask(strLen)
	} else if n >= uint(strLen) {
		return stringMask(strLen)
	}
	chars := strings.Split(s, "")
	visible := chars[:n]
	return strings.Join(visible, "") + stringMask(strLen-int(n))
}

func stringMask(n int) string {
	if n == 0 {
		n = 1
	}
	return strings.Repeat("*", n)
}

func WordSplit(arr [2]string) string {
	notPossibleText := "not possible"
	targetWord := arr[0]
	if targetWord == "" {
		return notPossibleText
	}

	words := strings.Split(arr[1], ",")
	result := notPossibleText
WordCheck:
	for _, word := range words {
		if strings.Index(targetWord, word) == 0 {
			chars := strings.Split(targetWord, "")
			last := chars[len(word):]
			result = word + "," + strings.Join(last, "")
			break WordCheck
		}
	}

	return result
}

func VariadicSet(i ...interface{}) []interface{} {
	result := make([]interface{}, 0)
	for _, v := range i {
		if contains(result, v) {
			continue
		}
		result = append(result, v)
	}
	return result
}

func contains(arr []interface{}, c interface{}) bool {
	for _, v := range arr {
		if v == c {
			return true
		}
	}
	return false
}
