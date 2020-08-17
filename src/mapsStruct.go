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
	i := 1
	for ; i <= 10; i++ {
		if i != 7 {
			a[fmt.Sprintf("id%v\n", i)] = i
		} else if err := recover(); err != nil {
			panic("error")
		}
	}
	delete(a, "id2")
	fmt.Printf("A: %v\n", a)
	defer ranger(a)
	fmt.Printf("User: %v\n", user)
}

func ranger(a map[string]int) int {
	for _, v := range a {
		if v == 5 {
			break
		}
		fmt.Println("Each item: \n", v)
	}
	return 0
}
