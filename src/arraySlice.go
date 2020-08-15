package main

import "fmt"

// ArraySlice
func ArraySlice() {
	a := make([]int, 3, 50)
	d := a
	a = append(a, 3)
	fmt.Printf("A : %v\n", a)
	fmt.Printf("D : %v\n", d)
	b := []int{6, 12, 7}
	c := b[:len(b)-1]
	a = append(a, c...)
	fmt.Printf("Itself : %v\n", a)
	fmt.Printf("Length : %v\n", len(a))
	fmt.Printf("Capacity : %v\n", cap(a))
}
