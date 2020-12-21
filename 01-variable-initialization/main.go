package main

import (
	"fmt"
)

func main() {

	var a float64 = 20.0
	var b string
	var c int


	a = 32.9
	b = "my b string"
	c = 123
	d := "my string value"

	fmt.Printf("a is of type %T\n", a)
	fmt.Printf("b is of type %T\n", b)
	fmt.Printf("c is of type %T\n", c)
	fmt.Printf("d is of type %T\n", d)

	fmt.Printf(fmt.Sprintf("a value is %f\n", a))
	fmt.Printf(fmt.Sprintf("b value is %s\n", b))
	fmt.Printf(fmt.Sprintf("c value is %d\n", c))
	fmt.Printf(fmt.Sprintf("d value is %s\n", d))

}
