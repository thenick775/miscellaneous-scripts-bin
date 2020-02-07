package main

import (
	"fmt"
	"regexp"
)

func SplitBefore(s string, re *regexp.Regexp) []string { //split a string before a specified regex, unsupported but rolled my own
	var r []string
	var p int
	is := re.FindAllStringIndex(s, -1)
	if is == nil {
		return append(r, s)
	}
	for _, p := range is { //adjust ranges to split by so that the matched regex is left at the beginning of the line
		p[1] = p[1] - (p[1] - p[0] + 1)
	}
	for x, i := range is {
		if x == 0 {
			continue
		}
		r = append(r, s[p:i[1]])
		p = i[1]
	}
	return append(r, s[p:])
}

func main() {
	fmt.Println("Hello, playground")
	stri:=`Feb 07, 00:33      0.10      1.72     236      2.46     39.7
Feb 07, 00:03      0.10      0.00      83      2.86     39.7
Feb 06, 23:33      0.10      8.18     358      2.66     39.7
Feb 06, 23:03      0.13      2.63     240      2.53     39.7
Feb 06, 22:33      0.13      1.72     212      2.43     39.7
Feb 06, 22:03      0.13      5.00      57      2.51     39.7
Feb 06, 21:33      0.16      2.44     253      2.45     39.9
Feb 06, 21:03      0.16      2.50     243      2.43     39.9
Feb 06, 20:33      0.16      2.22     237      2.38     39.7
Feb 06, 20:03      0.20      2.17     236      2.37     40.1
Feb 06, 19:33      0.23      2.63     235      2.37     39.7
Feb 06, 19:03      0.26      2.78     230      2.37     39.9
Feb 06, 18:33      0.30      2.44     236      2.53     40.3
Feb 06, 18:03      0.33      2.78     228      2.52     40.3`

	x:=SplitBefore(stri,regexp.MustCompile(`^[a-zA-z]{3}`))
	fmt.Println(x)
}
