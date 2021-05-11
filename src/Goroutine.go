package main

import (
	"fmt"
	"time"
)

func Goroutine(){
	go compute(3)
	go compute(3)
	fmt.Scanln()
}

func compute(value int){
	for i := 0; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}