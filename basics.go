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

	var greeting string 
	greeting = "Hello there"
	fmt.Println(greeting)

	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)


	var field string = "DevOps"
	var tools string = "Golang"
	fmt.Println("In", field, tools, "is necessary")
	fmt.Printf("In %v\n", field)

	var grades int = 33
	fmt.Printf("Marks: %d\n", grades)

	var first string = "Jack"
	var mark int = 34
	fmt.Printf("Hey, %v! You have scored %d/100 in Math\n", first, mark)

	var (
		city string = "Berlin"
		number = 3
	)
	fmt.Println(city, number)

	country := "Germany"
	country = "France"
	// country = 12 => type of variable cannot be changed
	fmt.Println(country)
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
---------------------------------------
:= short assignment statement can be used in place of a var declaration with implicit type

\n =>  newline 

Printf -- format specifier:
General:
%v	the value in a default format
	when printing structs, the plus flag (%+v) adds field names
%#v	a Go-syntax representation of the value
%T	a Go-syntax representation of the type of the value
%%	a literal percent sign; consumes no value

Boolean:
%t	the word true or false

Integer:
%b	base 2
%c	the character represented by the corresponding Unicode code point
%d	base 10
%o	base 8
%O	base 8 with 0o prefix
%q	a single-quoted character literal safely escaped with Go syntax.
%x	base 16, with lower-case letters for a-f
%X	base 16, with upper-case letters for A-F
%U	Unicode format: U+1234; same as "U+%04X"

Floating-point and complex constituents:
%b	decimalless scientific notation with exponent a power of two,
	in the manner of strconv.FormatFloat with the 'b' format,
	e.g. -123456p-78
%e	scientific notation, e.g. -1.234456e+78
%E	scientific notation, e.g. -1.234456E+78
%f	decimal point but no exponent, e.g. 123.456
%F	synonym for %f
%g	%e for large exponents, %f otherwise. Precision is discussed below.
%G	%E for large exponents, %F otherwise
%x	hexadecimal notation (with decimal power of two exponent), e.g. -0x1.23abcp+20
%X	upper-case hexadecimal notation, e.g. -0X1.23ABCP+20

String and slice of bytes (treated equivalently with these verbs):
%s	the uninterpreted bytes of the string or slice
%q	a double-quoted string safely escaped with Go syntax
%x	base 16, lower-case, two characters per byte
%X	base 16, upper-case, two characters per byte

Slice:
%p	address of 0th element in base 16 notation, with leading 0x

Pointer:
%p	base 16 notation, with leading 0x
The %b, %d, %o, %x and %X verbs also work with pointers,
formatting the value exactly as if it were an integer.
The default format for %v is:

bool:                    %t
int, int8 etc.:          %d
uint, uint8 etc.:        %d, %#x if printed with %#v
float32, complex64, etc: %g
string:                  %s
chan:                    %p
pointer:                 %p


%f     default width, default precision
%9f    width 9, default precision
%.2f   default width, precision 2
%9.2f  width 9, precision 2
%9.f   width 9, precision 0
*/
