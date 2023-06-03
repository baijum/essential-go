package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func main() {
	persons := []Person{
		Person{Name: "Huck", Age: 11},
		Person{Name: "Tom", Age: 10},
		Person{Name: "Polly", Age: 52},
	}
	fmt.Println(persons)
}
