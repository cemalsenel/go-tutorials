//* A pointer is a variable that holds memory address of another variable
//* They point where the memory is allocated and provide ways to find or even change the value located at the memory location
///* & =Y address-of aoperator
//* * => dereference operator

package main
import "fmt"

func main(){
	i := 10
	fmt.Printf("%T %v \n", &i, &i)
	fmt.Printf("%T %v \n", *(&i), *(&i))



	//* Declaring and Initializing a pointer
	var i1 *int
	var s1 *string
	fmt.Println(i1) // nil
	fmt.Println(s1) // nil

	//* initializing - 1
	a := 10
	var ptr_a *int = &a
	fmt.Println(ptr_a)

	//* initializing - 2
	s := "hello"
	var ptr_s = &s
	fmt.Println(ptr_s)

	//* initializing - 3
	str := "hello"
	ptr_str := &str
	fmt.Println(ptr_str)

	k := "hello"
	var j *string = &k
	fmt.Println(j)
	var h = &k
	fmt.Println(h)
	g := &k
	fmt.Println(g)


	//* Dereferencing a ppointer
	m := "hello"
	fmt.Printf("%T %v \n", m, m)
	mm := &m
	*mm = "world"
	fmt.Printf("%T %v \n", m, m)

	y := [3]int{10, 20, 30}
	fmt.Printf("%v \n", y)
	fmt.Println(&y)
	(*&y)[0] = 100
	fmt.Printf("%v \n", y)

	//* Passing by value
	//* all basic types (int, bool, float, string, array) are passed by value
	//* only the copy is accessed or modified, and the original value is never modified

	//* Passing by reference
	//* Go supports pointers, allowing you to pass references to values within your program
	//* In call by reference/pointer, the address of the variable is passed into the function call as the actual parameter
	//* All the operations in the function are performed on the value stored at the address of the actual parameters
	//* Slices are passed by reference, by default 
	//* Maps are passed by reference, by default 
}