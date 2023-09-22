package main

import (
	"fmt"
	"math/rand"
)

func getSides() int {
	var sides int
	fmt.Println("How many sides do you want your dice to roll?")
	fmt.Scanln(sides)
	return sides
}

func rollDice(sides int) int {
	var result int
	result = rand.Intn(sides - 1)
	return result
}

func main() {
	side = getSides()
	diceResult = rollDice(sides)
	fmt.Println("Your dice result is: ", diceResult)
}
