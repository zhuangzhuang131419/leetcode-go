package leetcode

import (
	"strconv"
	"strings"
)

func complexNumberMultiply(num1 string, num2 string) string {
	real1, _ := strconv.ParseInt(strings.Split(num1, "+")[0], 10, 64)
	ima1, _ := strconv.ParseInt(strings.Split(strings.Split(num1, "+")[1], "i")[0], 10, 64)

	real2, _ := strconv.ParseInt(strings.Split(num2, "+")[0], 10, 64)
	ima2, _ := strconv.ParseInt(strings.Split(strings.Split(num2, "+")[1], "i")[0], 10, 64)

	result_re := real1*real2 - ima1*ima2
	result_im := ima1*real2 + real1*ima2

	return strconv.Itoa(int(result_re)) + "+" + strconv.Itoa(int(result_im)) + "i"
}
