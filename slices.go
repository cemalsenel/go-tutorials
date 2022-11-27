package main
import "fmt"

func main(){
	//* more flexible than arrays

	slice := []int{1, 2, 3}
	fmt.Println(slice)
	fmt.Println(cap(slice))

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	slice_1 := arr[1:8]
	slice_1[0] = 11
	fmt.Println(slice_1) //* [11 3 4 5 6 7 8]
	fmt.Println(arr) //* [1 11 3 4 5 6 7 8 9 10]


	slice1 := make([]int, 5, 10)

	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	fmt.Println(cap(arr))
	fmt.Println(cap(slice_1))

	slice = append(slice, 10, 20, 30)
	fmt.Println(slice)
	fmt.Println(cap(slice))

	//* append two slices
	arr1 := [5]int{1, 2, 3, 4, 5}
	slice2 := arr1[:2]
	arr2 := [5]int{6, 7, 8, 9, 10}
	slice3 := arr2[:2]
	new_slice := append(slice2, slice3...)
	fmt.Println(new_slice)

	//* copy slice
	copy_slice := make([]int, 10)
	num := copy(copy_slice, new_slice)
	fmt.Println(copy_slice)
	fmt.Println("Number of elements copied: ", num)


	//* looping through a slice
	for index, value := range copy_slice {
		fmt.Println(index, "=>", value)
	}


	arr5 := [5]int{10, 20, 90, 70, 60}
	slice5 := arr5[:3]
	fmt.Println(cap(slice5))

	slice_5 := make([]int, 5, 20)
	new_slice5 := append(slice5, slice_5...)
	fmt.Println(cap(new_slice5))
	fmt.Println(new_slice5)
}