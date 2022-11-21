package main

import (
	"fmt" 
	"strconv")

func main(){
	var i int = 9
	var f float64 = float64(i)
	fmt.Printf("%.2f\n", f)

	var a float64 = 4.4
	var b int = int(a)
	fmt.Printf("%v\n", b)

	var c string = strconv.Itoa(i) // convert int to string
	fmt.Printf("%v, %T\n", c, c)

	var d string = "3"
	e, err := strconv.Atoi(d)
	fmt.Printf("%v, %T \n", e, e)
	fmt.Printf("%v, %T \n", err, err)
}