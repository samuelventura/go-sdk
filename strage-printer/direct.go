package main

import (
	"fmt"
)

func main() {
	fmt.Println(strangePrinter("tbgtgb"))                                            //4
	fmt.Println(strangePrinter("dddccbdbababaddcbcaabdbdddcccddbbaabddb"))           //15
	fmt.Println(strangePrinter("baacdddaaddaaaaccbddbcabdaabdbbcdcbbbacbddcabcaaa")) //19
}

func strangePrinter(so string) int {
	s := make([]rune, 0, len(so))
	lc := '?'
	for _, c := range so {
		if c != lc {
			s = append(s, c)
		}
		lc = c
	}
	n := len(s)
	dd := make([][]int, n)
	for i := range dd {
		dd[i] = make([]int, n)
		dd[i][i] = 1
	}
	for i := range dd {
		dd[i] = make([]int, n)
		dd[i][i] = 1
	}
	for k := 2; k <= n; k++ {
		for i := 0; i <= n-k; i++ {
			j := i + k - 1
			if k == 2 {
				dd[i][j] = 2
			} else {
				if s[i] == s[j] {
					//proved by induction
					dd[i][j] = min(dd[i][j-1], dd[i+1][j])
					//dd[i][j] = min(dd[i][j], 1+dd[i+1][j-1])
				} else {
					mc := 1 << 16 //just large enough to be > 100/2
					for mm := i; mm < j; mm++ {
						mc = min(mc, dd[i][mm]+dd[mm+1][j])
					}
					dd[i][j] = mc
				}
			}
		}
	}
	return dd[0][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
