package main

import "fmt"

// Create the function fuelGauge() here
func fuelGauge(fuel int) {
	fmt.Print(fuel, "The ship has %d gallons of fuel left")

	// return fuelGauge
}

// Create the function calculateFuel() here
func calculateFuel(planet string) int {
	var fuel int

	switch planet {
	case "Venus":
		fuel = 300000
	case "Mercury":
		fuel = 500000
	case "Mars":
		fuel = 700000
	}

	return fuel
}

// Create the function greetPlanet() here

// Create the function cantFly() here

// Create the function flyToPlanet() here

func main() {
	// Test your functions!
	//fuelGauge(4)
	//fmt.Println(calculateFuel("Mars"))

	// Create `planetChoice` and `fuel`

	// And then liftoff!
	//

}
