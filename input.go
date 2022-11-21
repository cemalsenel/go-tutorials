package main

import "fmt"

func main(){
	/*
	var name string
	var is_muggle bool
	fmt.Print("Please enter your name & are you a muggle: ")
	fmt.Scanf("%s %t", &name, &is_muggle)
	fmt.Println("Hey there, ", name, is_muggle)
	*/

	var a string
	var b int
	fmt.Println("Enter a string and a number: ")
	count, err := fmt.Scanf("%s %d", &a, &b)
	fmt.Println("count: ", count)
	fmt.Println("error: ", err)  //  if does not match ==> error: expected integer
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)

}

/*
Scanf Return Values:

count:
err:
*/