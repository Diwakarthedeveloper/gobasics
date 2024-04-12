package main

import "fmt"

// Create the function fuelGauge() here
func fuelGauge(fuel int) {
	fmt.Printf("The ship has %d gallons of fuel left ", fuel)

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
func greenPlanet(planet string) {
	fmt.Println("Welcome to the ", planet)

}

// Create the function cantFly() here
func cantFly() {
	fmt.Println("We do not have the available fuel to fly there.")
}

// Create the function flyToPlanet() here
func flyToPlanet(planet string, fuel int) int {
	var fuelRemaining, fuelCost int
	fuelRemaining = fuel
	fuelCost = calculateFuel(planet)

	if fuelRemaining >= fuelCost {
		greenPlanet(planet)

		//fuelRemaining = fuelRemaining - fuelCost
		fuelRemaining -= fuelCost
	} else {
		cantFly()
	}
	return fuelRemaining
}

func main() {
	// Test your functions!
	//fuelGauge(4)
	//fmt.Println(calculateFuel("Mercury", 600000))
	fmt.Println(flyToPlanet("Mercury", 600000))

	// Create `planetChoice` and `fuel`
	fuel := 1090000
	planetChoice := "Venus"
	fuel = flyToPlanet(planetChoice, fuel)
	fuelGauge(fuel)

	// And then liftoff!

}
