package main

import "fmt"

func addHundred(x int) int {
        return x + 100
}
func partialSum(x ...int) func() {
        sum := 0
        for _, value := range x {
                sum += value
        }
        return func() {
                fmt.Println(addHundred(sum))
        }
}
func main() {
        partial := partialSum(1, 2, 3, 4, 5)
        partial()
}

/*
package main

import "fmt"

func addHundred(x int) int {
        return x + 100
}
func partialSum(add100 func(x int) int, x ...int) int {
        sum := 0
        for _, value := range x {
                sum += value
        }
        return add100(sum)

}
func main() {
        partial := partialSum(addHundred, 1, 2, 3)
        fmt.Println(partial)
}
*/