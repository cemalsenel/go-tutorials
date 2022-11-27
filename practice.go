package main

import "fmt"

func main() {
        arr := [10]string{"a", "b", "c"}
        hashmap := make(map[string]int)
        my_slice := arr[:]
        fmt.Println(len(my_slice))
        fmt.Println(cap(my_slice))
        fmt.Println(len(hashmap))
}