package main

import "fmt"

// arrays only have length while  slices have length(len) as wella s capacity(cap)
func main() {
	myTutors := []string{"Kirsty", "Mishell", "Jose", "Neil"}
	fmt.Println(myTutors)
	//fmt.Println(len(myTutors))
	//fmt.Println(cap(myTutors))
	fmt.Println(myTutors, len(myTutors), cap(myTutors)) // printing slice, initial length of slice and capacity of initial slice

	//Now we will append length of slice by 1 which will automatically DOUBLE he capacity
	myTutors = append(myTutors, "Josh")
	fmt.Println(myTutors, len(myTutors), cap(myTutors)) // output of lngth = 4 and capacity = 8
}
