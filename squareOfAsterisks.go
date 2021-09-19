package main

import (
	"fmt"
	"strconv"
)

func main() {
	var inputSize string

	for {
		fmt.Print("Enter size of the side of a square (-1 to exit program) : ")
		_, e := fmt.Scanf("%s", &inputSize)
		if e != nil {
			fmt.Println("we have an error : ", e)
			break
		}

		size, err := strconv.Atoi(inputSize)
		if err != nil {
			fmt.Println("We accept only integer number, please try again.")
			fmt.Println("")
			continue
		}

		fmt.Println("Size of the side of a square is :", size)
		if size >= 1 && size <= 20 {
			for i := 0; (i) < size; i++ {
				for j := 0; (j) < size; j++ {
					if i == 0 || (i) == size-1 || j == 0 || (j) == size-1 {
						fmt.Print("* ")
					} else {
						fmt.Print(" ")
					}
				}
				fmt.Println()
			}
			fmt.Println("")
		} else if size == -1 {
			fmt.Println("-----Exit program.-----")
			break
		} else {
			fmt.Println("We accept size between 1 and 20, please try again.")
			fmt.Println("")
			continue
		}
	}
}
