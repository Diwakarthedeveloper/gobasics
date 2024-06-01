package main

import "fmt"

func sum(x int, y int) int { // x int, y int here int are the input data type, and last int is the return data type
	return x + y
}

func swap(x, y string) (string, string) { // there will be two string values returned
	return y, x
}
func split(sum int) (x, y int) { // (x, y int) return variable already defined here and their data type
	x = sum * 4 / 9
	y = sum - x
	return // return variable not given here as in above line 6 and 10 because its already defined at line 12
}

func sum_of_num(nums ...int) int { //...int means inout value should be a array of int value
	total := 0
	for _, num := range nums { // for _, _ means as for loops needs index, but it is not required so provided underscore
		total += num // adding the array one y one
	}
	return total
}

func main() {
	result := sum(21, 24)
	fmt.Println("sum= ", result)

	fmt.Println("Multiple Return Values:")
	a, b := swap("Artificial ", "Intelligence")
	fmt.Println("a, b=", a, b)

	fmt.Println("Named Return Values:")
	sum := 100 // Example sum
	x, y := split(sum)
	fmt.Printf("The split of %d is: x = %d, y = %d\n", sum, x, y)

	fmt.Println("Variadic Functions:")
	sum_result := sum_of_num(1, 2, 3, 4)
	fmt.Println("The sum is:", sum_result)

}
