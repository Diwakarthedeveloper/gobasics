package main

import "fmt"

func main() {
	fmt.Println("Below you see the diffrence in size of 4 bytes(float32) and 8 bytes(float64)")
	var a float32 = 3.14159265358979323846264338327950288419716939937510582097494
	var b float64 = 3.141592653589793238462643383279502884197169399375105820974944

	fmt.Println("Size(4 bytes) of 32 bit = ", a)
	fmt.Println("Size(8 bytes) of 64 bit = ", b)

}

//output
// Below you see the diffrence in size of 4 bytes(float32) and 8 bytes(float64)
// Size(4 bytes) of 32 bit =  3.1415927
// Size(8 bytes) of 64 bit =  3.141592653589793
