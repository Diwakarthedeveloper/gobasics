package main

import "fmt"

type Student struct {
	firstName string
	lastName  string
	age       int
	grade     int
}

func main() {
	// struct Instances defined below in this format --> instanceVariable := structname{labels: values, labels: values}
	peter := Student{"Peter", "Bookman", 16, 11}
	fmt.Println(peter)

	scott := Student{firstName: "Scott", lastName: "Peterson", grade: 12}
	fmt.Println(scott)

}
