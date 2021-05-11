package main

import "fmt"

type aStruct struct {
	id    int8
	name  string
	langs []string
}

type zType string

func Pointer() {
	a := aStruct{id: 12, name: "majid", langs: []string{"go", "js"}}
	var b *aStruct = &a
	var z zType = "bar"
	var y *zType = &z
	fmt.Println(z, *y)
	z = "foo"
	fmt.Println(z, *y, *b)
}
