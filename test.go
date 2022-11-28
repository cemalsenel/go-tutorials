package main
import (
	"fmt"
)

func modify(s map[string]int) {
	s["A"] = 100
}
// func main() {
// 	ascii_codes := map[string]int{}
// 	ascii_codes["A"] = 65
// 	fmt.Println(ascii_codes)
// 	modify(ascii_codes)
// 	fmt.Println(ascii_codes)
// }

// type Movie struct {
// 	name   string
// 	rating float32
// }

// func getMovie(s string, r float32) (m Movie) {
// 	m = Movie{
// 			name:   s,
// 			rating: r,
// 	}
// 	return
// }

// func main() {
// 	mov := getMovie("xyz", 2.1)
// 	mov1 := getMovie("abc", 3.3)
// 	movies := make([]Movie, 5)
// 	movies = append(movies, mov)
// 	movies = append(movies, mov1)
// 	for _, value := range movies {
// 			fmt.Println(value)
// 	}
// }

type Movie struct {
	name   string
	rating float32
}

// func main() {
// 	mov := Movie{"xyz", 2.1}
// 	mov1 := Movie{"abc", 2.1}
// 	if mov.rating == mov1.rating || mov != mov1 {
// 			fmt.Println("condition met")
// 	} else if mov.rating == mov1.rating {
// 			fmt.Println("condition_2 met")
// 	}
// }



// func main() {
// 	mov := Movie{"xyz", "", 2.1}
// 	fmt.Printf("%+v", mov)
// 	mov.increaseRating()
// 	mov.getSummary()
// 	fmt.Printf("%+v", mov)
// }

// func (m Movie) getSummary() {
// 	m.summary = "summary"
// }

// func (m *Movie) increaseRating() {
// 	m.rating += 1
// }


// type Rectangle struct {
// 	length  int
// 	breadth int
// }

// func (r Rectangle) area() int {
// 	return r.length * r.breadth
// }

// func (r *Rectangle) incLength(n int) {
// 	for i := 0; i < n; i++ {
// 		r.length += i
// 	}
// }

// func main() {
// 	r := Rectangle{breadth: 10, length: 5}
// 	fmt.Println(r.area())
// 	fmt.Println(r)
// 	r.incLength(7)
// 	fmt.Println(r.area())
// 	fmt.Println(r)
// }


// type Employee struct {
// 	eid int
// 	id  int
// }

// func main() {
// 	employees := make([]Employee, 5)
// 	for i := range employees {
// 			employees[i] = Employee{i, i + 10}
// 			fmt.Println(employees[i])
// 	}
// }

// type Employee struct {
// 	eid int
// 	id  int
// }

// func (e Employee) get_id() int {
// 	return e.eid + 10
// }

// func main() {
// 	employees := make([]Employee, 5)
// 	for i := range employees {
// 			employees[i] = Employee{eid: i}
// 			employees[i].id = employees[i].get_id()
// 			fmt.Printf("%+v\n", employees[i])
// 	}
// }

// type Student interface {
// 	getPercentage() int
// 	getName()
// }

// type Undergrad struct {
// 	name   string
// 	grades []int
// }

// func (u Undergrad) getPercentage() int {
// 	sum := 0
// 	for _, v := range u.grades {
// 		sum += v
// 	}
// 	return sum / len(u.grades)
// }

// func printPercentage(s Student) {
// 	fmt.Println(s.getPercentage())
// }

// func main() {
// 	grades := []int{90, 75, 80}
// 	u := Undergrad{"Ross", grades}
// 	printPercentage(u)
// }

// type Student interface {
// 	getPercentage() int
// 	getName() string
// }

// type Undergrad struct {
// 	name   string
// 	grades []int
// }

// func (u Undergrad) getPercentage() int {
// 	sum := 0
// 	for _, v := range u.grades {
// 		sum += v
// 	}
// 	return sum / len(u.grades)
// }
// func (u Undergrad) getName() string {
// 	return u.name
// }

// func printData(s Student) {
// 	fmt.Println(s.getName())
// 	fmt.Println(s.getPercentage())
// }

// func main() {
// 	grades := []int{90, 75, 80}
// 	u := Undergrad{"Ross", grades}
// 	printData(u)
// }

type Student interface {
	getPercentage() int
	getName() string
}

type Undergrad struct {
	name   string
	grades []int
}

type Postgrad struct {
	name   string
	grades []int
}

func (p Postgrad) getPercentage() int {
	sum := 0
	for _, v := range p.grades {
			sum += v
	}
	return ((sum * 100) / (len(p.grades) * 200))
}
func (p Postgrad) getName() string {
	return p.name
}

func (u Undergrad) getPercentage() int {
	sum := 0
	for _, v := range u.grades {
			sum += v
	}
	return sum / len(u.grades)
}
func (u Undergrad) getName() string {
	return u.name
}

func printData(s Student) {
	fmt.Println(s.getName())
	fmt.Println(s.getPercentage())
}

func main() {
	u := Undergrad{"Ross", []int{90, 75, 80}}
	p := Postgrad{"Joe", []int{150, 190, 185}}
	printData(u)
	printData(p)
}