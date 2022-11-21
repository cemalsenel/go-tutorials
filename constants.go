package main
import "fmt"

const PI float64 = 3.14 //* global constant
func main(){
	//* syntax ==> const <const-name> <data-type(optional)> = <value>

	const name string = "hello"
	//* it will not assign new value to variable =>> name = "hi"
	const age int = 2  //* if we try to assign a value later, it will not allow us. we have to assign value to a constant at the same line
					   //* it is not allowed to use shorthand declaration
    var radius float64 = 5.0
	var area float64
	area = PI * radius * radius
	fmt.Println("Area of Circle is : ", area)
}