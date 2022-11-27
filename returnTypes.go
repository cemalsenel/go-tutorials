package main 
import "fmt"

//* returning multiple values
//* Option-1
// func operation(a int, b int) (int, int){
// 	sum := a+b
// 	diff := a-b
// 	return sum, diff
// }

//* Option-2
func operation(a int, b int) (sum int, diff int) {
	sum = a+b
	diff = a-b
	return 
}

func sumNumbers(numbers ...int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}

//* Variadic Function
func printDetails(student string, subjects ...string){
	fmt.Println("Hey,", student, ", here are your subjects - ")
	for _, sub := range subjects {
		fmt.Printf("%s, ", sub)
	}
}

func f() (int, int, int) {
	return 45, 55, 65
}


func main(){

	sum, difference := operation(20, 10)
	fmt.Println(sum, difference)

	fmt.Println(sumNumbers())
	fmt.Println(sumNumbers(10))
	fmt.Println(sumNumbers(10, 20))
	fmt.Println(sumNumbers(10, 20, 30, 40, 50))

	printDetails("Joe", "Math", "Biology")

	fmt.Println()

	v, y, _ := f() //* blank identifier
	fmt.Println(v, y)
}