package main

import "fmt"

func main() {
        var a, b bool = true, false
        fmt.Println(a && b)
        fmt.Println(a || b)

		var c bool = false
        result := 10 > 50
        fmt.Println(!(c && result))
}