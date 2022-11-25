package main

import "fmt"

func main() {
        var x, y string = "foo", "bar"
        x += y
        fmt.Println(x)

		var a, b int = 27, 7
        a /= b
        fmt.Println(a)

		var c, d int = 100,9
        c /= d
        fmt.Println(c)
        c %= d
        fmt.Println(c)

}