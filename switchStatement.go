package main

import "fmt"

func main() {
        var a, b = 100, 5
        switch {
        case a/b == 10:
                fmt.Println("10")
        case a/b == 20:
                fmt.Println("20")
        case a/b == 10:
                fmt.Println("30")
        default:
                fmt.Println("default")
        }


		day := "sunday"
        switch day {
        case "monday":
                fmt.Println("monday")
        case "tuesday":
                fmt.Println("tuesday")
        case "wednesday":
                fmt.Println("wednesday")
        case "thursday":
                fmt.Println("thursday")
        case "friday":
                fmt.Println("friday")
        case "saturday", "sunday":
                fmt.Println("weekend")
        default:
                fmt.Println("default")
        }

		days := "wednesday"
        switch days {
        case "monday":
                fmt.Println("monday")
        case "tuesday":
                fmt.Println("tuesday")
        case "wednesday":
                fmt.Println("wednesday")
                fallthrough
        case "thursday":
                fmt.Println("thursday")
                fallthrough
        case "friday":
                fmt.Println("friday")
        case "saturday", "sunday":
                fmt.Println("weekend")
        default:
                fmt.Println("default")
        }

		var i, j = 10, 50

        switch {
        case i+j == 60:
                fmt.Println("equal to 60")
        case i+j <= 60:
                fmt.Println("less than or equal to 60")
                fallthrough
        default:
                fmt.Println("greater than 60")
        }

}