package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("Hey")

	name := "Lisa"
	fmt.Println(name)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var greeting string = "Hello there"
	fmt.Println(greeting)

	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}

// this is a single line comment

/*
this is a
multi-line
comment
*/

/*
Basic types :

- bool
- string
- int  int8  int16  int32  int64
- uint uint8 uint16 uint32 uint64 uintptr
- byte ==> alias for uint8
- rune => alias for int32
       => represents a Unicode code point
- float32 float64
- complex64 complex128

:= short assignment statement can be used in place of a var declaration with implicit type
*/
