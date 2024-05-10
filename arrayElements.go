package main

import "fmt"

func main() {
	// array below
	//lottery := [5]int{5, 48, 32, 1, 6}
	lottery := [...]int{5, 48, 32, 1, 6} // here even if we dont mention numberOfElements . go will find  by number of elements provided

	fmt.Println(lottery)

}
