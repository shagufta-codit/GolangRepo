package main

import "fmt"

func main() {
	fmt.Println("hello world")
	foo()
	bar()
}

func foo() {
	fmt.Println("i am in foo")
}

func bar() {
	fmt.Println("exiting the program")
}

