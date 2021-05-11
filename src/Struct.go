package main

import "fmt"

type Pet struct {
	Age        int8
	Snack      string
	Cute       bool
	AnimalType string
}

func Struct() {
	quark := Pet{
		Age:        12,
		Snack:      "chips",
		Cute:       true,
		AnimalType: "cat",
	}
	sanic := Pet{
		Age:        10,
		Snack:      "carrot",
		Cute:       true,
		AnimalType: "dog",
	}
	none := Pet{
		Age:        3,
		Snack:      "carrot",
		Cute:       false,
		AnimalType: "cactus",
	}
	myMap := map[string]Pet{
		"Quark": quark,
		"Sanic": sanic,
		"None":  none,
	}
	petter(myMap)
}

func petter(myPets map[string]Pet) int {
	numPets := 0
	for petName, petData := range myPets {
		switch petData.AnimalType {
		case "dog":
			fmt.Println(petName + " woof")
		case "cat":
			fmt.Println(petName + " meow")
		default:
			fmt.Println(petName + " what??")
		}
		numPets++
	}
	return numPets
}
