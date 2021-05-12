package main

import "fmt"

func BiggestNumber() {
	ExpressionMatter(1, 2, 3)
}

func ExpressionMatter(a int, b int, c int) int {
	x := a * (b + c)
	y := a * b * c
	z := a + b*c
	k := (a + b) * c
	j := []int{x, y, z, k}
	p := 0
	for i := 0; i < len(j); i++ {
		p = j[i-1]
		if j[i] < p {
			fmt.Println(j[i])
		}
	}
	return 9
}
