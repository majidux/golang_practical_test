package main

import "fmt"

func main() {
	var a int = 42
	var b *int = &a
	fmt.Println(a, *b)
	a = 21
	fmt.Println(a, *b)
}
