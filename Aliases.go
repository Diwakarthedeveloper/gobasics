package main

import "fmt"

func main() {
	fmt.Println("Type Alias:")
	type dj int64 // dj is alias for int64, which means now onwards dj will do the same role as int64 keyword
	var x dj = 224

	fmt.Println("x=", x)
}
