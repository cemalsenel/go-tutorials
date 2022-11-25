/*
& => bitwise AND
| => bitwise OR
^ => bitwise XOR
>> => right shift
<< => left shift
*/

package main

import "fmt"

func main() {
        var x, y int = 100,90
        fmt.Println(x & y)
        fmt.Println(x ^ y)
        fmt.Println(x | y)
		fmt.Println((x+y) >> 2)
		fmt.Println(!(((x+y) >> 2 ) == 47))
}
