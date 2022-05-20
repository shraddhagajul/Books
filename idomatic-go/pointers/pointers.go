package pointers

import "fmt"

type Person struct {
	Firstname string
	Middlename *string
	Lastname string
}


func PointerType() {
	x := 10
	var pointerToX *int
	pointerToX = &x
	fmt.Println("pointerToX : ", pointerToX)
}

// PointerToPrimitiveType : helper func to turn constant value into a pointer 
func PointerToPrimitiveType( middlename string ) *string {
	return &middlename
}

func PointersAndImmutability() {
	var f *int
	failedUpdate(f)
	fmt.Println("after value of f : ",f)
	fmt.Println("--------------")
}

func failedUpdate(px *int) {
	fmt.Println("before value of px : ", px)
	x := 10
	px = &x
	fmt.Println("after value of px : ", px)
}

func PointersAndMutability() {
	f := 2
	failedUpdate(&f)
	fmt.Println("after failedUpdate : ", f)
	update(&f)
	fmt.Println("after update : ", f)
}

func update(px *int) {
	*px = 15
}

func Slices() {
	inputSlice := make([]string, 3, 10)
	inputSlice[0] = "Hello"
	inputSlice[1] = "Rach"
	fmt.Println("before changeValueInSlice : ", inputSlice)
	fmt.Println("address of inputSlice : ", &inputSlice[0])
	changeValueInSlice(inputSlice)
	fmt.Println("after changeValueInSlice : ", inputSlice)
	fmt.Println("address of inputSlice : ", &inputSlice[0])

	appendToSlice(inputSlice)
	fmt.Println("after appendToSlice : ", inputSlice)
	
}

func changeValueInSlice(inputSlice []string) {
	inputSlice[1] = "Alex"
}

func appendToSlice(inputSlice []string) {
	inputSlice = append(inputSlice, "Marta")
	fmt.Println("after changeLengthOfSlice : ", inputSlice)
}
