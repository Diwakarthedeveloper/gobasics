package main

import "fmt"
// pass1 == pass1 -> true
// pass1 == badpass -> false are they are not equal
//123 != 12 -> true as the they are not equal and the operator make the statement true
//123 != 123 -> false as both are same and operator is of not equals too

func main() {
	lockCombo := "2-35-19"
	robAttempt := "1-1-1"

 
  if lockCombo != robAttempt { // this will make the whole statement true so the it will go in IF statement
  fmt.Println("The vault is now opened.")
  }
}
