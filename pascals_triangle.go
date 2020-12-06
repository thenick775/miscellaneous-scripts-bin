//program that prints pascals triangle to a specified depth
package main

import (
	"fmt"
	"math/big"
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

func pascalsTriangle(depth *big.Int) {
	res, leadspcnt := "", depth.Int64()
	for i := big.NewInt(0); i.Cmp(depth) == -1; i.Add(big.NewInt(1), i) {
		for j := big.NewInt(0); j.Cmp(big.NewInt(0).Add(big.NewInt(1), i)) == -1; j.Add(big.NewInt(1), j) {
			if j.Cmp(big.NewInt(0)) == 0 {
				res += fixedLengthString(int(leadspcnt+4), "")
			}
			res += binom(i, j).String() + " "
		}
		leadspcnt -= 1
		res += "\n"
	}
	fmt.Print(res)
}

func main() {
	pascalsTriangle(big.NewInt(12))
}
