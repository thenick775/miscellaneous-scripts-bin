//cyk for fun
//requires grammar already be in CNF format
//sample output:
// nonterm index: map[A:2 B:3 C:4 S:1]
// Grammar 1
// S->AB
// S->BC
// A->BA
// A->a
// B->CC
// B->b
// C->AB
// C->a
//
// Problem for grammar 1 : 1 
// baaba
// I: baaba is a member of the language
//
// Problem for grammar 1 : 2 
// ab
// I: ab is a member of the language
//
// Problem for grammar 1 : 3 
// a
// I: a is not a member of the language
//
// Problem for grammar 1 : 4 
// bac
// I: bac is not a member of the language
//
// Problem for grammar 1 : 5 
// aa
// I: aa is not a member of the language
//
// Problem for grammar 1 : 6 
// baaa
// I: baaa is a member of the language
//
// Problem for grammar 1 : 7 
//
// I:  is not a member of the language
//
// nonterm index: map[A:2 B:3 S:1]
// Grammar 2
// S->AB
// S->
// A->a
// B->b
// B->BB
//
// Problem for grammar 2 : 1 
// a
// I: a is not a member of the language
//
// Problem for grammar 2 : 2 
// ab
// I: ab is a member of the language
//
// Problem for grammar 2 : 3 
// abb
// I: abb is a member of the language
//
// Problem for grammar 2 : 4 
// aabb
// I: aabb is not a member of the language
//
// Problem for grammar 2 : 5 
//
// I:  is a member of the language

package main

import (
	"fmt"
	"strings"
)

type grammar struct {
	Prod   [][]string
	ToTest []string
}

func (g grammar) Cyk(input string, n int, r int, nonTermInd map[string]int) bool {
	tiles := make([][][]bool, n+1)
	
	if g.edgecase(input,nonTermInd){
		return true
	} else if input == ""{
		return false
	}

	for i := range tiles { //initialize
		tiles[i] = make([][]bool, n+1)
		for j := range tiles[i] {
			tiles[i][j] = make([]bool, r+1)
		}
	}

	for s := 1; s <= n; s++ { //bottom row, unit productions
		for _, production := range g.Prod {
			if string(input[s-1]) == production[1] && len(production[1]) == 1 {
				tiles[1][s][nonTermInd[production[0]]] = true
			}
		}
	}

	for l := 2; l <= n; l++ { //move upwards
		for s := 1; s <= n-l+1; s++ {
			for p := 1; p <= (l - 1); p++ {
				for _, production := range g.Prod {
					if len(production[1]) > 1 {
						rhs:=strings.Split(production[1],"")
						if tiles[p][s][nonTermInd[rhs[0]]] && tiles[l-p][s+p][nonTermInd[rhs[1]]]{
							tiles[l][s][nonTermInd[production[0]]]=true
						}
					}
				}
			}
		}
	}

	if tiles[n][1][1] {
		return true
	} else {
		return false
	}
}

func (g grammar) edgecase(input string,nontermmap map[string]int) bool{
	for _,val:=range g.Prod{
		if nontermmap[val[0]]==1 && val[1]=="" && input ==""{
			return true
		}
	}
	return false
}

func (g grammar) nonTermMap() (map[string]int,int) {
	m,count := make(map[string]int), 1

	for _, rule := range g.Prod {
		if _, ok := m[rule[0]]; !ok {
			m[rule[0]] = count
			count++
		}
	}
	return m,len(m)
}

func PrintGrammar(gram [][]string){
	for _,val:=range gram{
		fmt.Println(val[0]+"->"+val[1])
	}
	fmt.Println("")
}

func main() {
	grammars := []grammar{
		grammar{Prod: [][]string{
			[]string{"S", "AB"},
			[]string{"S", "BC"},
			[]string{"A", "BA"},
			[]string{"A", "a"},
			[]string{"B", "CC"},
			[]string{"B", "b"},
			[]string{"C", "AB"},
			[]string{"C", "a"}},
			ToTest: []string{ "baaba", "ab", "a","bac","aa","baaa",""},
		},
		grammar{Prod: [][]string{
			[]string{"S", "AB"},
			[]string{"S", ""},
			[]string{"A", "a"},
			[]string{"B", "b"},
			[]string{"B", "BB"}},
			ToTest: []string{ "a", "ab", "abb","aabb",""},
		},
	}

	for i, gram := range grammars {
		tmap,r:=gram.nonTermMap()
		fmt.Println("nonterm index:",tmap)
		fmt.Println("Grammar",i+1)
		PrintGrammar(gram.Prod)
		for j, test := range gram.ToTest {
			fmt.Println("Problem for grammar", i+1, ":", j+1,"\n"+test)
			res := gram.Cyk(test, len(test), r,tmap)
			if res {
				fmt.Print("I: ", test, " is a member of the language\n\n")
			} else {
				fmt.Print("I: ", test, " is not a member of the language\n\n")
			}
		}
	}
}
