package main

import "fmt"

type Describe interface {
	Describe()
}

func main() {

	var d1 Describe
	if d1 == nil {
		fmt.Printf("d1 is nil and has type %T value %v\n", d1, d1)
	}
	//If we try to call a method on the nil interface, the program will panic since
	//the nil interface neither has a underlying value nor a concrete type.
	d1.Describe()
}
