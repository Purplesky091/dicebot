package main

import (
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func generateNumber(min int, max int) (result int) {
	return rand.Intn(max-min) + min
}

func main() {
	var max int
	var dice string
	var sum int

	fmt.Print("Input dice: ")
	fmt.Scan(&dice)

	r, _ := regexp.Compile(`^(\d*)d(\d+)$`)
	var IsValid bool = r.MatchString(dice)
	fmt.Printf("%s is valid: %t\n", dice, IsValid)
	if !IsValid {
		os.Exit(1)
	}

	// parsing dice
	DiceValues := strings.Split(dice, "d")
	numOfRolls, _ := strconv.Atoi(DiceValues[0])         // can use _ because I verified before the second value is an integer.
	max, _ = strconv.Atoi(DiceValues[len(DiceValues)-1]) // can use _ because I verified before the second value is an integer.

	rolls := make([]int, numOfRolls)
	fmt.Printf("Rolling dice %s.\n", dice)
	for i := 0; i < numOfRolls; i++ {
		rolls[i] = generateNumber(1, max)
		fmt.Printf("Roll %d: %d\n", i+1, rolls[i])
		sum += rolls[i]
	}

	fmt.Printf("Sum: %d", sum)
}
