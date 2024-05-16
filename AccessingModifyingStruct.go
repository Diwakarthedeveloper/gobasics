package main

import "fmt"

type Restaurant struct {
	name             string
	typeOfRestaurant string
	yearEstablished  int
}

func main() {

	// Abelow wriiten code with lof

	//restaurant := Restaurant{name: "Codecademy Steakhouse", typeOfRestaurant: "Japanese",yearEstablished: 2011}

	restaurant := Restaurant{"Codecademy Steakhouse", "Japanese", 2011}
	fmt.Println(restaurant)

	fmt.Println(restaurant.name)
	fmt.Println(restaurant.typeOfRestaurant)
	fmt.Println(restaurant.yearEstablished)

	restaurant.name = "Skillsoft Steakhouse"
	restaurant.yearEstablished = 2022

	fmt.Println(restaurant)

}
