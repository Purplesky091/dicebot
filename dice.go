package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var MAX_VALUE int = 99

func printHelp() {
	fmt.Println("Please input a dice with format [number]d[number]")
	fmt.Println("[number] can be anywhere from 1 - 99")
}

func generateNumber(min int, max int) (result int) {
	return rand.Intn(max-min) + min
}

func parseDice(dice string) (FaceCount int, RollCount int) {
	DiceValues := strings.Split(dice, "d")
	FaceCount, faceErr := strconv.Atoi(DiceValues[len(DiceValues)-1])

	if DiceValues[0] == "" {
		RollCount = 1
	} else {
		parsedRoll, rollError := strconv.Atoi(DiceValues[0])
		if rollError != nil {
			printHelp()
			os.Exit(2)
		} else if parsedRoll > MAX_VALUE {
			fmt.Printf("You cannot roll more than %d dice.\n", MAX_VALUE)
			os.Exit(3)
		} else if parsedRoll <= 0 {
			printHelp()
			fmt.Println("Your first [number] is less than 0. It must be 1-99")
			os.Exit(4)
		}

		RollCount = parsedRoll
	}

	// parsing number of faces
	if faceErr != nil {
		fmt.Println("Incorrect format, please input a dice with format [number]d[number]")
		os.Exit(2)
	} else if FaceCount > MAX_VALUE {
		printHelp()
		fmt.Printf("You cannot have a dice with more than %d faces.\n", MAX_VALUE)
		os.Exit(3)
	} else if FaceCount <= 0 {
		printHelp()
		fmt.Println("Your second [number] is less than 0. It must be 1-99")
		os.Exit(4)
	}

	return FaceCount, RollCount
}

func main() {
	var FaceCount int
	var dice string
	var sum int
	var RollCount int

	// fmt.Print("Input dice: ")
	// fmt.Scan(&dice)
	dice = "9d50"
	dice = strings.ToLower(dice)

	// r, _ := regexp.Compile(`^(\d*)d(\d+)$`)
	// var IsValid bool = r.MatchString(dice)
	// fmt.Printf("%s is valid: %t\n", dice, IsValid)
	// if !IsValid {
	// 	os.Exit(1)
	// }

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
