package main

import (
	"fmt"
	"unicode"
)

func main() {

	n := 1000010
	fmt.Printf("|%v|", numberToWords(n))
}

var numNames = map[int]string{
	1: "One",
	2: "Two",
	3: "Three",
	4: "Four",
	5: "Five",
	6: "Six",
	7: "Seven",
	8: "Eight",
	9: "Nine",
}
var tensNames = map[int]string{
	1: "Ten",
	2: "Twenty",
	3: "Thirty",
	4: "Forty",
	5: "Fifty",
	6: "Sixty",
	7: "Seventy",
	8: "Eighty",
	9: "Ninety",
}
var tensOneNames = map[int]string{
	1: "Eleven",
	2: "Twelve",
	3: "Thirteen",
	4: "Fourteen",
	5: "Fifteen",
	6: "Sixteen",
	7: "Seventeen",
	8: "Eighteen",
	9: "Nineteen",
}

var mb = map[int]string{
	2: "Thousand",
	3: "Million",
	4: "Billion",
}

// 2,147,483,647
func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	var ans []string
	for num > 0 {
		temp := ths(num)
		ans = append(ans, temp)
		num /= 1000
	}
	res := ""
	pt := len(ans)

	for i := len(ans) - 1; i >= 0; i-- {
		res += ans[i]
		if i != 0 {
			if pt > 1 {
				if ans[i] == "" {
					pt--
				} else {
					res += " " + mb[pt]
				}
				pt--

			}
			if ans[i-1] != "" {
				res += " "
			}
		}
	}
	return res
}

func ths(n int) string {
	c := n % 10
	n /= 10
	b := n % 10
	n /= 10
	a := n % 10
	n /= 10
	res := ""
	if a != 0 {
		res += numNames[a] + "Hundred"
	}
	if b != 0 && b != 1 {
		res += tensNames[b]
	} else if b == 1 {
		if c == 0 {
			res += "Ten"
			return spc(res)
		} else {
			res += tensOneNames[c]
			return spc(res)
		}

	}
	if c != 0 {
		res += numNames[c]
	}
	return spc(res)
}

func spc(s string) string {
	res := ""
	for i, val := range s {
		if i != 0 && unicode.IsUpper(val) {
			res += " "
		}
		res += string(val)
	}
	return res
}
