package main

import (
	"fmt"
)

type child struct {
	job string
}

type dataStruct struct {
	id   int
	name string
	list []string
	child
}

// maps
func MapsStruct() {
	a := make(map[string]int)
	user := dataStruct{
		id:   1,
		name: "majid",
		list: []string{
			"js",
			"go",
		},
		child: child{job: "react"},
	}
Loop:
	i := 1
	for ; i <= 10; i++ {
		if i != 7 {
			a[fmt.Sprintf("id%v\n", i)] = i

		}
		break Loop
	}
	delete(a, "id2")
	fmt.Printf("A: %v\n", a)
	fmt.Printf("User: %v\n", user)
}
