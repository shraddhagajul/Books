package main

import (
	"fmt"
	"idomatic-go/tmi"
)

func main() {
	runCounter()
	updates()
	methodsSameAsFunc()
}

func runCounter() {
	var c tmi.Counter
	fmt.Println(c.String())
	c.Increment()
	fmt.Println(c.String())
}

func updates() {
	var c tmi.Counter
	tmi.DoUpdateWrong(c)
	fmt.Println("in main:", c.String())
	tmi.DoUpdateRight(&c)
	fmt.Println("in main:", c.String())
}

func methodsSameAsFunc() {
	myAdder := tmi.Adder{Start: 10}
	fmt.Println("Add : ", myAdder.AddTo(5))
}