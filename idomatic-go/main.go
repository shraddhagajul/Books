package main

import (
	"fmt"
	"idomatic-go/pointers"
)

func main() {
	pointers.PointerType()
	member := &pointers.Person{Firstname: "Alex",
		Middlename: pointers.PointerToPrimitiveType("William"),
		Lastname:   "Percy"}
	fmt.Println("Firstname : ", member.Firstname)
	fmt.Println("Middlename : ", *member.Middlename)
	fmt.Println("Lastname : ", member.Lastname)

	pointers.PointersAndImmutability()
	pointers.PointersAndMutability()
	pointers.Slices()
}
