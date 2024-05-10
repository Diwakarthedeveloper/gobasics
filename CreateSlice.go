package main

import "fmt"

func main() {
	myTutors := [4]string{"Kirsty", "Mishell", "Jose", "Neil"}
	Tutors := myTutors[:] //here array chnages into slice, where myTutorsis array and Tutors is slice
	subjects := []string{"Go", "Go", "Python"}
	fmt.Println(Tutors)
	fmt.Println(subjects)

}
