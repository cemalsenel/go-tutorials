package main
import (
	"fmt"
	"reflect"
)

func main(){
	var grades int = 42
	var message string = "hello world"
	var isCheck bool = true
	var amount float32 = 546.44

	fmt.Printf("grades = %v is type of %T \n", grades, grades)
	fmt.Printf("message = %v is type of %T \n", message, message)
	fmt.Printf("isCheck = %v is type of %T \n", isCheck, isCheck)
	fmt.Printf("amount = %v is type of %T \n", amount, amount)

	fmt.Printf("Type: %v \n", reflect.TypeOf(1000))
	fmt.Printf("Type: %v \n", reflect.TypeOf("1000"))
	fmt.Printf("Type: %v \n", reflect.TypeOf(1000.0))
	fmt.Printf("Type: %v \n", reflect.TypeOf(true))
}

// %T helps us to see type of variable