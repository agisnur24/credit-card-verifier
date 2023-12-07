package source

import (
	"strconv"
	"strings"
)

func CardVerifier(s string) string {
	var sum, sumDoublingOfDigits, sumOddOfDigits, digits, value int
	s = strings.Replace(s, " ", "", -1)

	byteNumber := []byte(s)
	byteToInt, _ := strconv.Atoi(string(byteNumber))

	arrOfDig := []int{}

	for byteToInt > 0 {
		value = byteToInt % 10

		arrOfDig = append(arrOfDig, value)

		byteToInt = byteToInt / 10
	}

	for i := 0; i < len(arrOfDig); i++ {
		if i%2 == 0 {
			sumOddOfDigits += arrOfDig[i]
		} else {
			digits = arrOfDig[i] * 2
			if digits > 9 {
				digits -= 9
			}
			sumDoublingOfDigits += digits
		}
	}

	sum = sumDoublingOfDigits + sumOddOfDigits

	if sum%10 != 0 {
		return "NO"
	}

	return "YES"
}
