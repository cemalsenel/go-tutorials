package main

import (
        "fmt"
)

func calcSquare(numbers []int) ([]int, bool) {
        squares := []int{}
        for _, v := range numbers {
                squares = append(squares, v*v)
        }
        return squares, true

}

func printStrings(s string, names ...string) {
        fmt.Println(s)
        for _, value := range names {
                fmt.Printf("%s, ", value)
        }

} 

func main() {
        arr := [10]string{"a", "b", "c"}
        hashmap := make(map[string]int)
        my_slice := arr[:]
        fmt.Println(len(my_slice))
        fmt.Println(cap(my_slice))
        fmt.Println(len(hashmap))

        nums := [3]int{10, 20, 15}
        fmt.Println(calcSquare(nums[:]))

        printStrings("Hey there", "Joe", "Monica", "Gunther")


}