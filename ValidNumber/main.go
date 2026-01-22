package main

import (
	"fmt"
	"unicode"
)

func main() {

	s := "-1E+3"
	fmt.Println(isNumber(s))

}

func isNumber(s string) bool {
	dot := false
	num := false
	ep := false
	for i := 0; i < len(s); i++ {
		if s[i] <= '9' && s[i] >= '0' {
			num = true
		}
		if ep {
			if s[i] == '.' || s[i] == 'e' || s[i] == 'E' {
				return false
			}
		}

		if s[i] == 'e' || s[i] == 'E' {
			ep = true
		}
		if s[i] == '-' || s[i] == '+' {
			if i == 0 {
				continue
			} else {
				if s[i-1] == 'e' || s[i-1] == 'E' {
					continue
				} else {
					return false
				}
			}
		}
		if dot && s[i] == '.' {
			return false
		}
		if s[i] == '.' {
			dot = true
		}

		if i != 0 && unicode.IsLetter(rune(s[i])) && unicode.IsLetter(rune(s[i-1])) {
			return false
		}
		if unicode.IsLetter(rune(s[i])) && (s[i] != 'e' && s[i] != 'E') {
			return false
		}

		if s[i] == 'e' || s[i] == 'E' {
			if !num {
				return false
			}
			if i == len(s)-1 {
				return false
			} else {
				if !unicode.IsDigit(rune(s[i+1])) {
					if s[i+1] == '-' || s[i+1] == '+' {
						if i+1 == len(s)-1 {
							return false
						}
						continue
					} else {
						return false
					}
				}
			}

		}

	}

	if !num {
		return false
	}

	return true
}
