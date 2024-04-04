package main

import (
	"fmt"
	"time"
)

// here we add string as  the function type which helps to call the function in another function with same type variable
func isItLateInNewYork() string {
	var lateMessage string
	t := time.Now()
	tz, _ := time.LoadLocation("America/New_York")
	nyHour := t.In(tz).Hour()
	if nyHour < 5 {
		lateMessage = "Goodness it is late"
	} else if nyHour < 16 {
		lateMessage = "It's not late at all!"
	} else if nyHour < 19 {
		lateMessage = "I guess it's getting kind of late"
	} else {
		lateMessage = "It's late"
	}

	// Returning the string lateMessage so that it can be called laater in another function
	return lateMessage

}

func main() {
	var nyLate string
	nyLate = isItLateInNewYork()
	fmt.Println(nyLate)
}
