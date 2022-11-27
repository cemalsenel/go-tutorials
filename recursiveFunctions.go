package main
import "fmt"

func factorial(n int) int{
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func print(n int) {
	if n == 0 {
			return
	}
	print(n - 1)
	fmt.Print(n)

}

func main(){

	n := 5
	result := factorial(n)
	fmt.Println("Factorial of", n, "is :", result)

	print(5)
}