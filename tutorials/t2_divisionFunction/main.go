package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "Hola Mundo, este argumento lo pasé a la función printMe!"
	printMe(printValue)
	var numerator int = 10
	var denominator int = 3
	var result, remainder, err = intDivision(numerator, denominator)
	if err != nil {
		fmt.Println(err.Error())
	} else if remainder == 0 {
		fmt.Printf("No remainder, result is %v\n", result)
	} else {
		// println(result, remainder)
		fmt.Printf("Result is %v with remainder %v\n", result, remainder)
	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("not divisible by zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
