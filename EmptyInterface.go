package main

import "fmt"

func describeInterface(i interface{}) {
	fmt.Println(" Type + %T, value = %v\n", i, i)
}

func ImplementInterface() {
	s := "Hello world"

	describeInterface(s)
}
