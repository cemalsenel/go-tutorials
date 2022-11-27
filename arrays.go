package main

import "fmt"

func main() {
	var grades_1 [3]int = [3]int{1, 2, 3}
	grades_2 := [3]int{1, 2, 3}
	grades_3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(grades_1)
	fmt.Println(grades_2)
	fmt.Println(grades_3)


	var fruits [2]string = [2]string{"apple", "orange"}
	fmt.Println(len(fruits))
	fruits[0] = "kiwi"
	fmt.Println(fruits[0])

	for i := 0; i < len(fruits); i++{
		fmt.Println(fruits[i])
	}

	for index, element := range fruits{
		fmt.Println(index, "=>", element)
	}

	var multi [4][3]int = [4][3]int{{1,2,3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}}

	fmt.Println(multi[0][0])

	for i := 0; i < len(multi); i++{
		fmt.Println(multi[i])
	}

	arr1 := [5]bool{true, true, true, true}
        for i := 0; i < len(arr1); i++ {
                if arr1[i] {
                        fmt.Println(i)
                }
        }
	
		arr := [5]string{"a", "b", "c"}
        for index, element := range arr {
                fmt.Println(index, "->", element)
        }
}
