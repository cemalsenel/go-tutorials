//* User defined data type
//* a structure that groups together data elements
//* provide a way to reference a series of grouped values through a single variable name
//* used when it makes sense to group or associate two or more data variables


package main
import "fmt"
//* struct declaration
type Student struct {
	name string
	rollNo int
	// marks []int
	// grades map[string]int
}
//* struct initialization
type Circle struct {
	x int 
	y int 
	radius float64
	area float64
}

func calcArea(c *Circle){
	const PI float64 = 3.14
	var area float64
	area = ( PI * c.radius * c.radius)
	(*c).area = area
}

func main(){

	var s Student
	fmt.Printf("%+v \n", s) // {name: rollNo:0 marks:[] grades:map[]}
	//st := new(Student)
	// st := Student {
	// 	name: "Jimmy",
	// 	rollNo: 12,
	// }

	st := Student{"Joe", 12}
	fmt.Printf("%+v \n", st)

	// Accessing Fields
	type Circle1 struct {
		x int 
		y int 
		radius int
		area float64
	}

	var c Circle1
	c.x = 5
	c.y = 6
	c.radius = 3
	fmt.Printf("%+v \n", c)

	//* Passing structs to functions
	q := Circle{x:5, y:5, radius:5, area:0}
	fmt.Printf("%+v \n", q)
	calcArea(&q)
	fmt.Printf("%+v \n", q)


	//* Structs of the same type can be compared using Go's equality operators (==, !=)


}

