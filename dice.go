package main

import (
	"fmt"
	"math/rand"
	"regexp"
)

func generateNumber(min int, max int) (result int) {
	return rand.Intn(max-min) + min
}

func main() {
	var dice string
	fmt.Print("Input dice: ")
	fmt.Scan(&dice)

	r, _ := regexp.Compile(`^(\d*)d(\d+)$`)
	var IsValid bool = r.MatchString(dice)
	fmt.Printf("%s is valid: %t", dice, IsValid)
	// fmt.Printf("Printing random number between 1 - %d: %d\n", upperBound, generateNumber(1, upperBound))
}
