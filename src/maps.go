package main

import (
	"fmt"
)

type doc struct {
	id   int
	name string
	list []string
}

// maps
func Maps() {
	a := make(map[string]int)
	user := doc{
		id:   1,
		name: "majid",
		list: []string{
			"js",
			"go",
		},
	}
	for i := 1; i < 10; i++ {
		if i != 7 {
			a[fmt.Sprintf("id%v", i)] = i
		}
	}
	delete(a, "id2")
	fmt.Printf("A: %v\n", a)
	fmt.Printf("User: %v\n", user)
}
