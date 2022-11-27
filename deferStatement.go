//* "defer" statement delays the execution of a function until the surrounding function returns.

package main

import (
        "fmt"
        "strings"
)

func printString(str string){
        fmt.Printf("%q ", str)
}

func printInt(i int){
        fmt.Printf("%d ", i)
}

func printFloat(f float64){
        fmt.Printf("%.2f ", f)
}
func main() {
        printString("browser")
        defer printInt(32)
        defer printFloat(0.24)
        printString("chrome")
        printInt(90)
        defer printFloat(89)
        printInt(900)

		_, lower := getString("BROWSER")
        fmt.Println(lower)

		fmt.Print(greetings())
}

func getString(str string) (string, string) {
	return strings.ToLower(str), strings.ToUpper(str)
}

func greetings() (x, y string) {
	x = "hello "
	y = "world"
	return
}

// func greetings() (x, y string) {
// 	x = "hello "
// 	y = "world"
// 	return
// }