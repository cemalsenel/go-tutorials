//* blueprint for a method set
package main
import "fmt"

type shape interface{
	//* Method signatures
	area() float64
	perimeter() float64
}

type square struct {
	side float64
}

func (s square) area() float64{
	return s.side * s.side
}

func (s square) perimeter() float64{
	return 4 * s.side
}

type rect struct {
	length, breadth float64
}

func (r rect) area() float64{
	return r.length * r.breadth
}

func (r rect) perimeter() float64{
	return 2*r.length + 2*r.breadth
}

func printData(s shape){
	fmt.Println(s)
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}

func main(){

	r := rect{length: 3, breadth: 4}
	c := square{side: 5}
	printData(r)
	printData(c)

}