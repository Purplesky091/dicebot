package main

import (
	"fmt"
	"math/rand"
)

func generateNumber(min int, max int) (result int) {
	return rand.Intn(max-min) + min
}

func main() {
	fmt.Printf("Printing random number between 1 - 4: %d\n", generateNumber(1, 4))
	// for i := 0; i < 3; i++ {
	// fmt.Printf("Printing random number %d\n", rand.Intn(50))
	// }
}
