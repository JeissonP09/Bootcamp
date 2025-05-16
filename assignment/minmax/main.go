/* Create a program that asks the user about a minimum value, a maximum value and a list of values. After that, filter out all the values that are out of range and print the values in range using a slice. */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
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

func getValueInput(logDescription string) (valFlt float64, err error) {
	fmt.Println(logDescription)
	inputValue, err := getInput()
	if err != nil {
		return getValueInput("The number imput is invalid")
	}
	valFlt, err = strconv.ParseFloat(inputValue, 64)
	if err != nil {
		return getValueInput(fmt.Sprintf("The value %s is invalid. Please enter other number:", inputValue))
	}
	return
}

func main() {
	min, err := getValueInput("Enter the minimum value:")
	if err != nil {
		fmt.Println("The number imput is invalid")
		return
	}
	max, err := getValueInput("Enter the maximum value:")
	if err != nil {
		fmt.Println("The number imput is invalid")
		return
	}

	fmt.Println("The ranks are: ", min, "--", max)

	values := make([]float64, 0)
	attemps := 0

	for attemps < 3 {
		fmt.Println("Enter a list of values separated by spaces:")
		input, err := getInput()
		if err != nil {
			fmt.Println("Error reading input:", err)
			attemps++
			continue
		}
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
