package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strings"
	"unicode"
)

var ErrStringStartsWithDigit = errors.New("string starts with digit")
var ErrZeroDigit = errors.New("zero is invalid digit to unpack")
var ErrNumbersNotAllowed = errors.New("there is a number in string, only single digits allowed")

func Unpack(s string) (string, error) {
	var ans strings.Builder
	for pos, char := range s {
		if unicode.IsDigit(char){
			if ans.Len() != 0 {
				// get previous character
				prev := s[pos-1]
				if unicode.IsDigit(rune(prev)){
					// there is a number in string
					return "", ErrNumbersNotAllowed
				}
				// get number of repetitions minus one because there is already one symbol present
				count := (int(char) - '0') - 1
				if count < 0 {
					// we got zero digit somewhere in string
					return "", ErrZeroDigit
				}
				ans.WriteString(strings.Repeat(string(prev), count))
			} else {
				// string starts with digit
				return "", ErrStringStartsWithDigit
			}
		} else {
			ans.WriteString(string(char))
		}
	}
	return ans.String(), nil
}
