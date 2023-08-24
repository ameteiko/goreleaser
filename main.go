package main

import (
	"fmt"

	_ "k8s.io/apimachinery"
)

func main() {

	println("Ba dum, tss!")

	println("Ba dum, tss!")
	println("Ba dum, tss!")
	println("Ba dum, tss!")
	println("Ba dum, tss!")

	println("For a release")

	if a() && a() {
		fmt.Println("a && b")
	}
}

func a() bool {
	fmt.Println("a")

	return false
}

func b2() bool {
	fmt.Println("b")

	return true
}

// doc
func c1() bool {
	fmt.Println("b")

	return true
}

func d1() bool {
	fmt.Println("b")

	return true
}
