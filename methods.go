/*
func(<reciever>) <method_name>(<parameters>)<return_params>{ <-- code --> }
*/

package main
import "fmt"

type Circle struct {
	radius float64
	area float64
}

func (c *Circle) calcArea(){
	c.area = 3.14 * c.radius * c.radius
}

func main(){
	c := Circle{radius: 5}
	c.calcArea()
	fmt.Printf("%+v\n", c)
}