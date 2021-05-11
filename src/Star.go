package main

import "fmt"

func Star() {
	h := starCreator()
	fmt.Println("\n", h)
}

func starCreator() int {
	o := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			o = j
		}
	}
	return o
}
