package main
import "fmt"

func main(){
	var a,b string = "foo", "bar"
	fmt.Println(a+b)

	var c,d int = 24, 7
	fmt.Println(c - d)
	fmt.Println(c * d)
	fmt.Println(c / d)
	fmt.Println(c % d)
	c++
	fmt.Println(c)


}