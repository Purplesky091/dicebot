package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var MAX_VALUE int = 100

func printHelp() {
	fmt.Println("Please input a dice with format [number]d[number]")
	fmt.Printf("[number] can be anywhere from 1 - %d\n", MAX_VALUE)
}

func generateNumber(min int, max int) (result int) {
	return rand.Intn(max-min) + min
}

func parseValue(diceValue string) (value int) {
	value, rollError := strconv.Atoi(diceValue)
	if rollError != nil {
		printHelp()
		os.Exit(2)
	} else if value > MAX_VALUE {
		printHelp()
		os.Exit(3)
	} else if value <= 0 {
		printHelp()
		os.Exit(4)
	}

	return value
}

func parseDice(dice string) (FaceCount int, RollCount int) {
	DiceValues := strings.Split(dice, "d")

	if DiceValues[0] == "" {
		RollCount = 1
	} else {
		RollCount = parseValue(DiceValues[0])
	}
	FaceCount = parseValue(DiceValues[len(DiceValues)-1])

	return FaceCount, RollCount
}

func main() {
	var FaceCount int
	var dice string
	var sum int
	var RollCount int

	fmt.Print("Input dice: ")
	fmt.Scan(&dice)
	dice = strings.ToLower(dice)

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
