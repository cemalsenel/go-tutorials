package main

import "fmt"

func main() {

        arr := [10]int{10, 20}
        slice := arr[2:8]
        fmt.Println(len(slice))
        fmt.Println(cap(slice))
        
        for i := 0; i <= 5; i++ {
                fmt.Println(i * i)
                if i == 3 {
                        continue
                }
        }


}