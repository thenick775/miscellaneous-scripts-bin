//program that prints pascals triangle to a specified depth
package main

import (
	"fmt"
	"strconv"
)

func fixedLengthString(length int, str string) string {
	return fmt.Sprintf(fmt.Sprintf("%%%d.%ds", length, length), str)
}

func fact(n int64) int64 {
	res := int64(1)
	if n > 0 {
		for i := int64(1); i <= n; i++ {
			res *= i
		}
	}
	return res
}

func binom(n int64, k int64) int64 {
	return fact(n) / (fact(k) * fact(n-k))
}

func pascalsTriangle(depth int64) {
	res, leadspcnt := "", depth
	for i := int64(0); i < depth; i++ {
		for j := int64(0); j < i+1; j++ {
			if j == 0 {
				res += fixedLengthString(int(leadspcnt+4), "")
			}
			res += fixedLengthString(2, strconv.FormatInt(binom(i, j), 10)) + " "
		}
		leadspcnt -= 1
		res += "\n"
	}
	fmt.Print(res)
}

func main() {
	pascalsTriangle(12)
}
