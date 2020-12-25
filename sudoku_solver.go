//soduku solver
//sample output:
// Hello from the Soduku solver
// ********
// Problem 1
// Initial Board:
// +-------+-------+-------+
// | 0 0 0 | 0 9 5 | 6 0 8 |
// | 7 0 0 | 1 0 8 | 0 0 0 |
// | 0 0 9 | 3 0 0 | 0 0 1 |
// +-------+-------+-------+
// | 0 3 7 | 0 8 0 | 0 0 0 |
// | 4 0 2 | 0 7 0 | 3 0 9 |
// | 0 0 0 | 0 3 0 | 2 6 0 |
// +-------+-------+-------+
// | 1 0 0 | 0 0 9 | 7 0 0 |
// | 0 0 0 | 6 0 3 | 0 0 4 |
// | 2 0 6 | 8 5 0 | 0 0 0 |
// +-------+-------+-------+
// solving using Depth First Search...
// Solved Board:
// +-------+-------+-------+
// | 3 2 1 | 7 9 5 | 6 4 8 |
// | 7 5 4 | 1 6 8 | 9 3 2 |
// | 8 6 9 | 3 2 4 | 5 7 1 |
// +-------+-------+-------+
// | 6 3 7 | 9 8 2 | 4 1 5 |
// | 4 1 2 | 5 7 6 | 3 8 9 |
// | 5 9 8 | 4 3 1 | 2 6 7 |
// +-------+-------+-------+
// | 1 8 3 | 2 4 9 | 7 5 6 |
// | 9 7 5 | 6 1 3 | 8 2 4 |
// | 2 4 6 | 8 5 7 | 1 9 3 |
// +-------+-------+-------+
// Solved in:1.585154ms

// ********
// Problem 2
// Initial Board:
// +-------+-------+-------+
// | 0 0 0 | 0 0 0 | 5 4 1 |
// | 0 0 0 | 7 1 6 | 8 0 0 |
// | 0 0 0 | 0 0 0 | 6 0 0 |
// +-------+-------+-------+
// | 0 0 3 | 0 0 9 | 0 5 0 |
// | 5 0 0 | 0 7 0 | 0 0 6 |
// | 0 8 0 | 4 0 0 | 2 0 0 |
// +-------+-------+-------+
// | 0 0 7 | 0 0 0 | 0 0 0 |
// | 0 0 6 | 1 9 8 | 0 0 0 |
// | 2 3 8 | 0 0 0 | 0 0 0 |
// +-------+-------+-------+
// solving using Depth First Search...
// Solved Board:
// +-------+-------+-------+
// | 6 7 2 | 9 8 3 | 5 4 1 |
// | 3 4 5 | 7 1 6 | 8 2 9 |
// | 8 9 1 | 2 5 4 | 6 3 7 |
// +-------+-------+-------+
// | 1 6 3 | 8 2 9 | 7 5 4 |
// | 5 2 4 | 3 7 1 | 9 8 6 |
// | 7 8 9 | 4 6 5 | 2 1 3 |
// +-------+-------+-------+
// | 9 1 7 | 5 3 2 | 4 6 8 |
// | 4 5 6 | 1 9 8 | 3 7 2 |
// | 2 3 8 | 6 4 7 | 1 9 5 |
// +-------+-------+-------+
// Solved in:30.323303ms

// ********
// Problem 3
// Initial Board:
// +-------+-------+-------+
// | 8 0 0 | 0 0 0 | 0 0 0 |
// | 0 0 3 | 6 0 0 | 0 0 0 |
// | 0 7 0 | 0 9 0 | 2 0 0 |
// +-------+-------+-------+
// | 0 5 0 | 0 0 7 | 0 0 0 |
// | 0 0 0 | 0 4 5 | 7 0 0 |
// | 0 0 0 | 1 0 0 | 0 3 0 |
// +-------+-------+-------+
// | 0 0 1 | 0 0 0 | 0 6 8 |
// | 0 0 8 | 5 0 0 | 0 1 0 |
// | 0 9 0 | 0 0 0 | 4 0 0 |
// +-------+-------+-------+
// solving using Depth First Search...
// Solved Board:
// +-------+-------+-------+
// | 8 1 2 | 7 5 3 | 6 4 9 |
// | 9 4 3 | 6 8 2 | 1 7 5 |
// | 6 7 5 | 4 9 1 | 2 8 3 |
// +-------+-------+-------+
// | 1 5 4 | 2 3 7 | 8 9 6 |
// | 3 6 9 | 8 4 5 | 7 2 1 |
// | 2 8 7 | 1 6 9 | 5 3 4 |
// +-------+-------+-------+
// | 5 2 1 | 9 7 4 | 3 6 8 |
// | 4 3 8 | 5 2 6 | 9 1 7 |
// | 7 9 6 | 3 1 8 | 4 5 2 |
// +-------+-------+-------+
// Solved in:1.130869143s

package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type board struct {
	Board      [9][9]int
	EmptyCount int
	CheckRow   [9]map[int]struct{}
	CheckCol   [9]map[int]struct{}
	Check3x3   [9]map[int]struct{}
}

func printBoard(b board) {
	res := ""
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			res += "+-------+-------+-------+\n"
		}
		res += "|"
		for j := 0; j < 9; j++ {
			if j%3 == 0 && j != 0 {
				res += " |"
			}
			res += " " + strconv.Itoa(b.Board[i][j])
		}
		res += " |\n"
	}
	res += "+-------+-------+-------+\n"
	fmt.Print(res)
}

func checkValid(b board) bool {
	checkmap := make(map[int]struct{})
	for i := 0; i < 9; i++ { //check rows
		for j := 0; j < 9; j++ {
			if _, ok := checkmap[b.Board[i][j]]; !ok && b.Board[i][j] != 0 {
				checkmap[b.Board[i][j]] = struct{}{}
			} else if b.Board[i][j] != 0 {
				//fmt.Println("duped row:", b.Board[i][j], " i: ", i, " j: ", j)
				return false
			}
		}
		checkmap = make(map[int]struct{})
	}
	checkmap = make(map[int]struct{})
	for i := 0; i < 9; i++ { //check columns
		for j := 0; j < 9; j++ {
			if _, ok := checkmap[b.Board[j][i]]; !ok && b.Board[j][i] != 0 {
				checkmap[b.Board[j][i]] = struct{}{}
			} else if b.Board[j][i] != 0 {
				//fmt.Println("duped col:", b.Board[j][i], " i: ", i, " j: ", j)
				return false
			}
		}
		checkmap = make(map[int]struct{})
	}
	checkmap = make(map[int]struct{})
	//check 3x3
	for i := 0; i < 9; i += 3 { //check 3x3
		for j := 0; j < 9; j += 3 {
			for k := i; k < i+3; k++ {
				for l := j; l < j+3; l++ {
					if _, ok := checkmap[b.Board[k][l]]; !ok && b.Board[k][l] != 0 {
						checkmap[b.Board[k][l]] = struct{}{}
					} else if b.Board[k][l] != 0 {
						//fmt.Println(checkmap)
						//fmt.Println("duped 3x3:", b.Board[k][l], " k: ", k, " l: ", l)
						return false
					}
				}
			}
			checkmap = make(map[int]struct{})
		}
	}
	return true //valid
}

func timeWrapper(b board, algostr string, f func(b *board) bool) {
	fmt.Println("solving using " + algostr + "...")
	start := time.Now()
	f(&b)
	duration := time.Since(start)
	fmt.Print("Solved in:", duration, "\n\n")
}

func depthFirst(b *board) bool {
	if b.EmptyCount == 0 {
		fmt.Println("Solved Board:")
		printBoard(*b)
		return true
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b.Board[i][j] == 0 { //this is a space not yet filled
				for k := 9; k >= 1; k-- {
					if b.IsMoveValid(k, i, j) {
						b.Board[i][j] = k
						b.EmptyCount -= 1
						if depthFirst(b) {
							return true
						}
						b.Board[i][j] = 0
						b.EmptyCount += 1
						delete(b.CheckRow[i], k)
						delete(b.CheckCol[j], k)
						delete(b.Check3x3[CountFor3x3(i, j)], k)
					}
				}
				return false
			}
		}
	}

	return false
}

func (b *board) Init() bool {
	if checkValid(*b) {
		for i := 0; i < 9; i++ { //init row
			b.CheckRow[i] = make(map[int]struct{})
			for j := 0; j < 9; j++ {
				if b.Board[i][j] != 0 {
					b.CheckRow[i][b.Board[i][j]] = struct{}{}
				}
			}
		}

		for i := 0; i < 9; i++ { //init col
			for j := 0; j < 9; j++ {
				if b.CheckCol[i] == nil {
					b.CheckCol[i] = make(map[int]struct{})
				}
				if b.Board[j][i] != 0 {
					b.CheckCol[i][b.Board[j][i]] = struct{}{}
				}
			}
		}

		for i := 0; i < 9; i += 3 { //check 3x3
			for j := 0; j < 9; j += 3 {
				if b.Check3x3[CountFor3x3(i, j)] == nil {
					b.Check3x3[CountFor3x3(i, j)] = make(map[int]struct{})
				}
				for k := i; k < i+3; k++ {
					for l := j; l < j+3; l++ {
						if b.Board[k][l] != 0 {
							b.Check3x3[CountFor3x3(i, j)][b.Board[k][l]] = struct{}{}
						}
					}
				}
			}
		}
	} else {
		return false
	}
	return true
}

func CountFor3x3(i int, j int) int {
	if i < 3 && j < 3 {
		return 0
	} else if i < 3 && j >= 3 && j < 6 {
		return 1
	} else if i < 3 && j >= 6 {
		return 2
	} else if i >= 3 && i < 6 && j < 3 {
		return 3
	} else if i >= 3 && i < 6 && j >= 3 && j < 6 {
		return 4
	} else if i >= 3 && i < 6 && j >= 6 {
		return 5
	} else if i >= 6 && j < 3 {
		return 6
	} else if i >= 6 && j >= 3 && j < 6 {
		return 7
	} else if i >= 6 && j >= 6 {
		return 8
	}
	return -1
}

func (b *board) IsMoveValid(val int, i int, j int) bool {
	added := false
	if _, ok := b.CheckRow[i][val]; !ok {
		b.CheckRow[i][val] = struct{}{}
		added = true
	} else {
		return false
	}

	if _, ok := b.CheckCol[j][val]; !ok {
		b.CheckCol[j][val] = struct{}{}
		added = true
	} else {
		if added {
			delete(b.CheckRow[i], val)
		}
		return false
	}

	if _, ok := b.Check3x3[CountFor3x3(i, j)][val]; !ok {
		b.Check3x3[CountFor3x3(i, j)][val] = struct{}{}
		added = true
	} else {
		if added {
			delete(b.CheckRow[i], val)
			delete(b.CheckCol[j], val)
		}
		return false
	}
	return true
}

func parseinput(input string) board {
	scanner, res := bufio.NewScanner(strings.NewReader(input)), board{}
	count := 0
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ",")
		for i, cell := range s {
			v, err := strconv.Atoi(cell)
			if err == nil {
				res.Board[count][i] = v
				if v == 0 {
					res.EmptyCount += 1
				}
			} else {
				panic(err)
			}
		}
		count += 1
	}
	return res
}

func main() {
	fmt.Println("Hello from the Soduku solver")
	input := []string{
		`0,0,0,0,9,5,6,0,8
7,0,0,1,0,8,0,0,0
0,0,9,3,0,0,0,0,1
0,3,7,0,8,0,0,0,0
4,0,2,0,7,0,3,0,9
0,0,0,0,3,0,2,6,0
1,0,0,0,0,9,7,0,0
0,0,0,6,0,3,0,0,4
2,0,6,8,5,0,0,0,0`,
		`0,0,0,0,0,0,5,4,1
0,0,0,7,1,6,8,0,0
0,0,0,0,0,0,6,0,0
0,0,3,0,0,9,0,5,0
5,0,0,0,7,0,0,0,6
0,8,0,4,0,0,2,0,0
0,0,7,0,0,0,0,0,0
0,0,6,1,9,8,0,0,0
2,3,8,0,0,0,0,0,0`,
		`8,0,0,0,0,0,0,0,0
0,0,3,6,0,0,0,0,0
0,7,0,0,9,0,2,0,0
0,5,0,0,0,7,0,0,0
0,0,0,0,4,5,7,0,0
0,0,0,1,0,0,0,3,0
0,0,1,0,0,0,0,6,8
0,0,8,5,0,0,0,1,0
0,9,0,0,0,0,4,0,0`,
	}

	for i, inputstr := range input {
		fmt.Println("********\nProblem", i+1)
		b := parseinput(inputstr)
		if b.Init() {
			fmt.Println("Initial Board:")
			printBoard(b)
			timeWrapper(b, "Depth First Search", depthFirst)
		} else {
			fmt.Println("illegal board")
			printBoard(b)
		}
	}
}
