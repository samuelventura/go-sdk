package main

import (
	"fmt"
	"os"
)

func main() {
	run(4, "tbgtgb")                                             //4
	run(15, "dddccbdbababaddcbcaabdbdddcccddbbaabddb")           //15
	run(19, "baacdddaaddaaaaccbddbcabdaabdbbcdcbbbacbddcabcaaa") //19
}

type node struct {
	mi int
	mj int
	mm int
	mc int
}

func run(ec int, so string) {
	c, s, a := strangePrinter(so)
	fmt.Printf("%02d %s\n", c, string(a)) //15
	for i := range s {
		if s[i] != a[i] {
			fmt.Printf("Mismatch at %d\n", i)
			os.Exit(1)
		}
	}
	if ec != c {
		fmt.Printf("Mismatch %d!=%d\n", ec, c)
		os.Exit(1)
	}
}

func strangePrinter(so string) (int, []rune, []rune) {
	s := make([]rune, 0, len(so))
	a := make([]rune, 0, len(so))
	lc := '?'
	for _, c := range so {
		if c != lc {
			s = append(s, c)
			a = append(a, '-')
		}
		lc = c
	}
	n := len(s)
	fmt.Println(len(so), so)
	fmt.Println(string(s), n)
	dd := make([][]int, n)
	dn := make([][]*node, n)
	for i := range dd {
		dd[i] = make([]int, n)
		dn[i] = make([]*node, n)
	}
	count := count(0, n-1, s, dd, dn)
	draw(0, 0, n-1, s, a, dd, dn)
	return count, s, a
}

func count(i, j int, s []rune, dd [][]int, dn [][]*node) int {
	d := dd[i][j]
	if d > 0 {
		return d
	}
	if i == j {
		dd[i][j] = 1
		dn[i][j] = &node{i, j, -4, 1}
		return 1
	}
	minc := 2147483647
	mini := -1
	minj := -1
	minm := -1
	if s[i] == s[j] {
		left := count(i, j-1, s, dd, dn)
		right := count(i+1, j, s, dd, dn)
		if left < right {
			minc = left
			mini = i
			minj = j - 1
			minm = -2
		} else {
			minc = right
			mini = i + 1
			minj = j
			minm = -3
		}
	} else {
		for m := i; m < j; m++ {
			left := count(i, m, s, dd, dn)
			right := count(m+1, j, s, dd, dn)
			both := left + right
			if both < minc {
				minc = both
				mini = i
				minj = j
				minm = m
			}
		}
	}
	//fmt.Printf("%02d %02d %02d %02d %02d %02d\n", i, j, mini, minj, minm, minc)
	dd[i][j] = minc
	dn[i][j] = &node{mini, minj, minm, minc}
	return minc
}

func draw(c, i, j int, s, a []rune, dd [][]int, dn [][]*node) int {
	n := dn[i][j]
	switch n.mm {
	case -2:
		c = line(c, n.mm, n.mi, n.mj+1, s[n.mi], s, a)
		c = draw(c, n.mi, n.mj, s, a, dd, dn)
	case -3:
		c = line(c, n.mm, n.mi-1, n.mj, s[n.mj], s, a)
		c = draw(c, n.mi, n.mj, s, a, dd, dn)
	case -4:
		c = line(c, n.mm, n.mi, n.mj, s[n.mi], s, a)
	default:
		c = draw(c, n.mi, n.mm, s, a, dd, dn)
		c = draw(c, n.mm+1, n.mj, s, a, dd, dn)
	}
	return c
}

func line(c, m, i, j int, r rune, s, a []rune) int {
	d := 0
	l := make([]rune, len(s))
	for p := range s {
		if p >= i && p <= j {
			l[p] = r
			if r != a[p] {
				d++
				a[p] = r
			}
		} else {
			l[p] = '-'
		}
	}
	if d > 0 {
		c++
		fmt.Printf("%02d %s\n", c, string(l))
		fmt.Printf("%02d %s\n", c, string(a))
	}
	return c
}
