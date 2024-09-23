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

func parseDice(dice string) (FaceCount int, RollCount int) {
	DiceValues := strings.Split(dice, "d")
	FaceCount, _ = strconv.Atoi(DiceValues[len(DiceValues)-1]) // can use _ because I verified before the second value is an integer.
	if DiceValues[0] == "" {
		RollCount = 1
	} else {
		RollCount, _ = strconv.Atoi(DiceValues[0]) // can use _ because I verified before the second value is an integer.
	}

	return FaceCount, RollCount
}

func main() {
	var FaceCount int
	var dice string
	var sum int
	var RollCount int

	fmt.Print("Input dice: ")
	fmt.Scan(&dice)

	r, _ := regexp.Compile(`^(\d*)d(\d+)$`)
	var IsValid bool = r.MatchString(dice)
	fmt.Printf("%s is valid: %t\n", dice, IsValid)
	if !IsValid {
		os.Exit(1)
	}

	FaceCount, RollCount = parseDice(dice)

	// rolls := make([]int, numOfRolls) // we'll add an array of values once we want to want to implement re-roll logic
	fmt.Printf("Rolling dice %s.\n", dice)
	for i := 0; i < RollCount; i++ {
		diceroll := generateNumber(1, FaceCount)
		fmt.Printf("Roll %d: %d\n", i+1, diceroll)
		sum += diceroll
	}

	fmt.Printf("Sum: %d", sum)
}
