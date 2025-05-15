/* Create a program that asks the user about a minimum value, a maximum value and a list of values. After that, filter out all the values that are out of range and print the values in range using a slice. */

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(input)
}

func minMax(min, max float64, values ...float64) []float64 {
	var inRange []float64

	for _, value := range values {
		if value >= min && value <= max {
			inRange = append(inRange, value)
		}
	}
	return inRange
}

func main() {
	fmt.Println("Enter minimum value:")
	minValue := getInput()
	min, err := strconv.ParseFloat(minValue, 64)
	if err != nil {
		fmt.Println("Invalid minimum value. Please enter a number.")
		return
	}

	fmt.Println("Enter maximum value:")
	maxValue := getInput()
	max, err := strconv.ParseFloat(maxValue, 64)
	if err != nil {
		fmt.Println("Invalid maximum value. Please enter a number.")
		return
	}

	fmt.Println("The ranks are: ", min, "--", max)

	var values []float64
	attemps := 0

	for attemps < 3 {
		fmt.Println("Enter a list of values separated by spaces:")
		input := getInput()
		inputValues := strings.Fields(input)
		values = nil
		valid := true

		for _, part := range inputValues {
			value, err := strconv.ParseFloat(part, 64)
			if err != nil {
				fmt.Println("Invalid input. Please enter numbers only.")
				valid = false
				break
			}
			values = append(values, value)
		}
		if valid {
			break
		}
		attemps++
	}

	if len(values) == 0 {
		fmt.Println("No valid values entered. Exiting.")
		return
	}

	inRange := minMax(min, max, values...)

	fmt.Println("Values in range:", inRange)

}
