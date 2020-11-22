package main

import (
	"fmt"
	"strings"
)

func closure(x []string, deps [][]string) string {
	res, old := "", ""
	fmt.Println("from X")
	for _, val := range x { //whatever is in x is a part of the closure
		for _, strval := range val {
			if !strings.ContainsRune(res, strval) {
				res += string(strval)
			}
		}
	}

	fmt.Println(res)
	fmt.Println("apply transitivity")

	for old != res {
		old = res

		for i, val := range deps {
			fmt.Println(i, val)
			notfound := false

			for _, lhsval := range val[0] {
				found := false
				for _, x := range res {
					if lhsval == x {
						found = true
					}
				}
				if !found {
					fmt.Println(res)
					fmt.Println("didnt find:" + string(lhsval))
					notfound = true
					break
				}
			}

			if notfound {
				continue
			} else {
				for _, val := range val[1] {
					fmt.Println("adding val: " + string(val))
					if !strings.ContainsRune(res, val) {
						res += string(val)
					}
				}

			}

		}
	}

	return res
}

//type to hold data
type xdepclosure struct {
	x    []string
	deps [][]string
}

func main() {
	//closures to compute here
	myclosures := []xdepclosure{
		xdepclosure{x: []string{"B"}, deps: [][]string{
			{"A", "BC"},
			{"CD", "E"},
			{"B", "D"},
			{"E", "A"},
		}},
		xdepclosure{x: []string{"A", "E"}, deps: [][]string{
			{"A", "D"},
			{"AB", "E"},
			{"BI", "E"},
			{"CD", "I"},
			{"E", "C"},
		}},
	}

	//compute closures
	for _,val :=range myclosures{
		fmt.Println("DEPS: ", val.deps)
		fmt.Println("X: ", val.x)
		
		res := closure(val.x, val.deps)
		fmt.Print("Final Attribute Closure: ", res, " for X: ", val.x,"\n\n")
	}
}
