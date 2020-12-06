//program that prints pascals triangle to a specified depth
//sample output:
//Row   0:                   1 
//Row   1:                  1 1 
//Row   2:                 1 2 1 
//Row   3:                1 3 3 1 
//Row   4:              1 4 6 4 1 
//Row   5:             1 5 10 10 5 1 
//Row   6:           1 6 15 20 15 6 1 
//Row   7:          1 7 21 35 35 21 7 1 
//Row   8:        1 8 28 56 70 56 28 8 1 
//Row   9:      1 9 36 84 126 126 84 36 9 1 
//Row  10:   1 10 45 120 210 252 210 120 45 10 1 
//Row  11:  1 11 55 165 330 462 462 330 165 55 11 1 

package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func fixedLengthString(length int, str string) string {
	return fmt.Sprintf(fmt.Sprintf("%%%d.%ds", length, length), str)
}

func fact(n *big.Int) *big.Int {
	res, t := big.NewInt(1), big.NewInt(0) //int64(1)
	if n.Cmp(t) == 1 {
		for i := big.NewInt(1); i.Cmp(new(big.Int).Add(big.NewInt(1), n)) == -1; i.Add(big.NewInt(1), i) {
			res.Mul(i, res)
		}
	}
	return res
}

func binom(n *big.Int, k *big.Int) *big.Int /*int64*/ {
	t := new(big.Int).Mul(fact(k), fact(big.NewInt(0).Sub(n, k)))
	return new(big.Int).Div(fact(n), t)
}

func reSpace(s string) string {
	res, t := "", strings.Split(s, "\n")
	finlen := len(t[len(t)-2])

	for i := 0; i < len(t)-1; i++ {
		res += "Row " + fixedLengthString(3, strconv.Itoa(i)) + ": " + fixedLengthString((finlen/2-len(t[i])/2)/2, "") + t[i] + "\n"
	}

	return res
}

func pascalsTriangle(depth *big.Int) {
	res, leadspcnt := "", depth.Int64()
	for i := big.NewInt(0); i.Cmp(depth) == -1; i.Add(big.NewInt(1), i) {
		for j := big.NewInt(0); j.Cmp(big.NewInt(0).Add(big.NewInt(1), i)) == -1; j.Add(big.NewInt(1), j) {
			if j.Cmp(big.NewInt(0)) == 0 {
				res += fixedLengthString(int(leadspcnt), "")
			}
			res += binom(i, j).String() + " "
		}
		leadspcnt -= 1
		res += "\n"
	}
	fmt.Print(reSpace(res))
}

func main() {
	pascalsTriangle(big.NewInt(12))
}
