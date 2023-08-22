package main

import (
	"fmt"
)

func main() {

	println("Ba dum, tss!")

	println("Ba dum, tss!")
	println("Ba dum, tss!")
	println("Ba dum, tss!")
	println("Ba dum, tss!")

	println("For a release")

	if a() && b() {
		fmt.Println("a && b")
	}
}

func a() bool {
	fmt.Println("a")

	return false
}

func b() bool {
	fmt.Println("b")

	return true
}
