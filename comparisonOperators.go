package main
import "fmt"

func main(){
	var city_1 string = "Berlin"
	var city_2 string = "Munich"
	fmt.Println(city_1 == city_2)
	fmt.Println(city_1 != city_2)

	var a, b int = 5, 6
	fmt.Println(a < b)
	fmt.Println(a > b)
	fmt.Println(a <= b)
	fmt.Println(a >= b)
}