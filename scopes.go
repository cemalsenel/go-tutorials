package main
import (
	"fmt"
)

var global string = "Global Variable"

func main() {
	city:= "Berlin"
	{
		country := "Germany"
		fmt.Println(country) // Germany  
		fmt.Println(city) // Berlin
	}
var local string = "Local Variable"
// fmt.Println(country) => throws error because of scope
fmt.Println(city)
fmt.Println(global)
fmt.Println(local)


}


/* 
Zero Values :
false(boolean), 0(integer), 0.0(float), ""(string), nil(pointers, functions, interfaces, maps)

*/