package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  isHeistOn := true
  eludedGuards := rand.Intn(100)

  if eludedGuards >= 50 {
    fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
  } else{ 
    isHeistOn := false
    fmt.Println("Plan a better disguise next time as ?",isHeistOn)
  }

openedVault := rand.Intn(100)

if isHeistOn == true && openedVault >= 70 {
  fmt.Println("Grab and GO!")
}else if openedVault >= 70 {
  isHeistOn := false
  fmt.Println("vault can’t be opened.",isHeistOn)
}
leftSafely := rand.Intn(5)
if isHeistOn == true








    fmt.Println("isHeistOn is currently:",isHeistOn)

}
