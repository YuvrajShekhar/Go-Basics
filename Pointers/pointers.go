package main

import "fmt"

func main() {
	age := 32

	var agePointer *int
	agePointer = &age

	fmt.Println(*agePointer, "age pointer value")

	getAdultYears(agePointer)
	fmt.Println(age, "=========adult years")
}

func getAdultYears(age *int) {
	*age = *age - 18
}
