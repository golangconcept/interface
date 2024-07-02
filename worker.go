package main

import "fmt"

type Worker interface {
	work()
}

type Person struct {
	name string
}

func (a Person) work() {
	fmt.Println(a.name, "is working")
}

func describe(w Worker) {
	fmt.Println("Interface type %T value %v \n", w, w)
}

func main() {
	p := Person{
		name: "Jai Pal",
	}

	var w Worker = p
	describe(w)
	w.work()
}
