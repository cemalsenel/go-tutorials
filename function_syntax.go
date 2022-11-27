package main 
import "fmt"

func main(){
	sumOfNumbers := addNumbers(5, 6)
	fmt.Println(sumOfNumbers)
	printGreeting("Jimmy")
}

func addNumbers(a int, b int) int {
	sum := a + b
	return sum
}

func printGreeting(str string) {
	fmt.Println("Hey there,", str)
}