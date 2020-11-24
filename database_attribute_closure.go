//computes closure of specified attribute given a set of functional dependencies
package main

import (
	"fmt"
	"strings"
)

//type to hold data
type xdepclosure struct {
	x    string
	deps [][]string
}

func closure(x string, deps [][]string) string {
	res, old, count := x, "", 0
	fmt.Println("from X,\n"+res+"\napply transitivity:")

	for old != res {
		old = res
		fmt.Println("Pass: ", count)

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
					//fmt.Println("adding val: " + string(val))
					if !strings.ContainsRune(res, val) {
						fmt.Println("adding val: " + string(val))
						res += string(val)
						fmt.Println(res)
					}
				}

			}

		}
		count++
	}

	return res
}

func printdeps(deps [][]string) {
	resstr := ""
	for _, val := range deps {
		resstr += fmt.Sprintf("%s->%s, ", val[0], val[1])
	}
	fmt.Println("Deps: ", resstr[:len(resstr)-2])
}

func main() {
	//closures to compute here
	myclosures := []xdepclosure{
		xdepclosure{x: "B", deps: [][]string{
			{"A", "BC"},
			{"CD", "E"},
			{"B", "D"},
			{"E", "A"},
		}},
		xdepclosure{x: "AE", deps: [][]string{
			{"A", "D"},
			{"AB", "E"},
			{"BI", "E"},
			{"CD", "I"},
			{"E", "C"},
		}},
		xdepclosure{x: "SI", deps: [][]string{
			{"S", "D"},
			{"I", "B"},
			{"IS", "Q"},
			{"B", "O"},
		}},
	}

	//compute closures
	for _, val := range myclosures {
		//fmt.Println("DEPS: ", val.deps)
		printdeps(val.deps)
		fmt.Println("Attribute to find closure for, X: ", val.x)

		res := closure(val.x, val.deps)
		fmt.Print("Final Attribute Closure: (", val.x, ")+ = ", res, "\n\n")
	}
}
