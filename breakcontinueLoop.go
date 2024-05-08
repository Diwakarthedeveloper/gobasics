package main

import (
	"fmt"
)

func main() {

	for count := 0; count < 20; count++ {
		//  CONTINUE STATEMENT BELOW  IF COUNT EQUALS 8 BELOW THIS LINE

		if count == 8 {
			continue
		}

		//  BREAK STATEMENT BELOW IF COUNT EQUALS 15 BELOW THIS LINE
		if count == 15 {
			break
		}

		fmt.Println(count)
	}

}
