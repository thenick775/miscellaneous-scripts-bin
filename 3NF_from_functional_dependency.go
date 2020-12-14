//program to decompose functional dependencies to 3NF table form
//sample output:
//**Problem:  1 **
//[[A C] [AB C] [C D] [C I] [CD I] [EC A] [EC B] [EI C]]
//non redundant deps:  [[A C] [C D] [CD I] [EC A] [EC B] [EI C]]
//f (non-redundant lhs):  [[A C] [C D] [C I] [EC A] [EC B] [EI C]]
//synthesized: [[A C] [C DI] [EC AB] [EI C]]
//
//Relations:
//R1(EIC)
//R2(ECAB)
//R3(CDI)
//R4(AC)
//
//**Problem:  2 **
//[[FD P] [FD N] [P N] [N P]]
//non redundant deps:  [[FD N] [P N] [N P]]
//f (non-redundant lhs):  [[FD N] [P N] [N P]]
//synthesized: [[FD N] [P N] [N P]]
//found same table, different order, removing: PN
//
//Relations:
//R1(NP)
//R2(FDN)
//
//**Problem:  3 **
//[[ABC D] [ABC E] [ABC G] [E B] [E C] [E G] [F A] [F H]]
//non redundant deps:  [[ABC D] [ABC E] [E B] [E C] [E G] [F A] [F H]]
//f (non-redundant lhs):  [[ABC D] [ABC E] [E B] [E C] [E G] [F A] [F H]]
//synthesized: [[ABC DE] [E BCG] [F AH]]
//
//Relations:
//R1(FAH)
//R2(EBCG)
//R3(ABCDE)

package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func in(idx int, list []int) bool {
	for i, _ := range list {
		if i == idx {
			return true
		}
	}
	return false
}

func printTables(deps [][]string) {
	depsloc := make([][]string, len(deps))
	copy(depsloc, deps)
	res, count, alreadyhere := "", 1, false
	
	for i := len(depsloc) - 1; i >= 0; i-- { //check for duplicate tables
		dep := strings.Join(depsloc[i], "")
		for j := len(depsloc) - 1; j >= 0; j-- {
			innerdep := strings.Join(depsloc[j], "")
			if i == j {
				continue
			}
			if SortString(dep) == SortString(innerdep) && len(dep) == len(innerdep) {
				fmt.Println("found same table, different order, removing:", innerdep)
				depsloc = append(depsloc[:j], depsloc[j+1:]...)
				alreadyhere = true
				break
			}
		}
		if !alreadyhere {
			res += "R" + strconv.Itoa(count) + "(" + dep + ")\n"
			count++
		}
		alreadyhere = false
	}
	fmt.Println("\nRelations:\n" + res)
}

func remove(s [][]string, i int) [][]string {
	tmp := make([][]string, len(s))
	copy(tmp, s)
	tmp[len(tmp)-1], tmp[i] = tmp[i], tmp[len(tmp)-1]
	return tmp[:len(tmp)-1]
}

func splitDeps(inputdeps [][]string) [][]string {
	res := [][]string{}

	for _, dep := range inputdeps {
		if len(dep[1]) > 1 {
			for _, val := range dep[1] {
				res = append(res, []string{dep[0], string(val)})
			}
		} else if len(dep[1]) == 1 {
			res = append(res, dep)
		}
	}
	return res
}

func closure(x string, deps [][]string) string {
	res, old, count := x, "", 0
	for old != res {
		old = res
		for _, val := range deps {
			notfound := false

			for _, lhsval := range val[0] {
				found := false
				for _, x := range res {
					if lhsval == x {
						found = true
					}
				}
				if !found {
					notfound = true
					break
				}
			}

			if notfound {
				continue
			} else {
				for _, val := range val[1] {
					if !strings.ContainsRune(res, val) {
						res += string(val)
					}
				}
			}
		}
		count++
	}
	return res
}

func to3NF(deps [][]string) {
	//split the deps, step 1
	splitdeps := splitDeps(deps)
	fmt.Println(splitdeps)

	//redundancy check, step 2
	for i := 0; i < len(splitdeps); i++ {
		lhs, rhs := splitdeps[i][0], splitdeps[i][1]
		theclosure := closure(lhs, remove(splitdeps, i))
		if strings.Contains(theclosure, rhs) {
			//fmt.Println("is redundant")
			splitdeps = append(splitdeps[:i], splitdeps[i+1:]...)
			i--
		}
	}
	fmt.Println("non redundant deps: ", splitdeps)

	//eliminate redundant attributes from the left hand side, step 3
	for i, dep := range splitdeps {
		lhs := dep[0]
		if len(lhs) > 1 {
			for _, val := range lhs {
				theclosure := closure(string(val), remove(splitdeps, i))
				for _, valc := range strings.ReplaceAll(lhs, string(val), "") {
					if strings.Contains(theclosure, string(valc)) {
						//fmt.Println(valc, " is redundant round 2")
						splitdeps[i] = []string{string(val), dep[1]}
					}
				}
			}
		}
	}

	fmt.Println("f (non-redundant lhs): ", splitdeps)

	//3nf synthesization
	res := [][]string{}
	checked := make(map[string][]int)
	for i, val := range splitdeps {
		tmp := []string{}
		if _, ok := checked[val[0]]; !ok {
			for j, vali := range splitdeps {
				if val[0] == vali[0] && i != j && !in(j, checked[val[0]]) {
					checked[val[0]] = append(checked[val[0]], j)

					if len(tmp) >= 2 {
						tmp = []string{val[0], tmp[1] + vali[1]}
					} else {
						tmp = []string{val[0], val[1] + vali[1]}
					}
				}
			}
			if len(tmp) == 0 {
				res = append(res, val)
			} else {
				res = append(res, tmp)
			}
		}
	}
	fmt.Println("synthesized:", res)
	printTables(res)
}

func main() {
	//attributes:="ABCDEI" //for 1
	//attributes:="FDPN" //for 2
	//attributes:="ABCDEFGH' //for 3
	deps := [][][]string{
		[][]string{{"A", "C"},
			{"AB", "C"},
			{"C", "DI"},
			{"CD", "I"},
			{"EC", "AB"},
			{"EI", "C"}},
		[][]string{{"FD", "PN"},
			{"P", "N"},
			{"N", "P"}},
		[][]string{{"ABC", "DEG"},
			{"E", "BCG"},
			{"F", "AH"},
		},
	}

	for i, dep := range deps {
		fmt.Println("**Problem: ", i+1, "**")
		to3NF(dep)
	}

}
