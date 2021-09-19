package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	for {
		var numbers []int
		var inputIsNumber bool

		fmt.Print("Enter the interger number ('exit' to end program) : ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		if scanner.Text() == "exit" {
			fmt.Println("-----End program-----")
			break
		}

		fmt.Println("Value input are : ", scanner.Text())
		inputValue := strings.Split(scanner.Text(), " ")

		for i := range inputValue {
			number, err := strconv.Atoi(inputValue[i])
			if err != nil {
				fmt.Println("We accept only interger number or something wrong, please try again.")
				inputIsNumber = false
				break
			} else {
				inputIsNumber = true
				numbers = append(numbers, number)
			}
		}

		if inputIsNumber {
			fmt.Println("Smallest value is :", findMinimumValue(numbers))
		}
	}
}

func findMinimumValue(numbers []int) int {
	smallestValue := numbers[0]
	for i := range numbers {
		if numbers[i] < smallestValue {
			smallestValue = numbers[i]
		}
	}
	return smallestValue
}
