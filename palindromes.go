package main

import (
	"fmt"
)

func main() {
	var inputValue string
	var isPalindrome bool

	for {
		fmt.Print("Enter number or text to check palindromes ('exit' to end program) : ")
		fmt.Scanln(&inputValue)

		if inputValue == "exit" {
			fmt.Println("-----End the program-----")
			break
		} else if len(inputValue) > 0 {
			for i := 0; i < len(inputValue)/2; i++ {
				if inputValue[i] == inputValue[len(inputValue)-i-1] {
					isPalindrome = true
					continue
				} else {
					isPalindrome = false
				}
			}

			if isPalindrome == true {
				fmt.Printf("%s is a palindrome.\n", inputValue)
				continue
			} else {
				fmt.Printf("%s is not a palindrome.\n", inputValue)
				continue
			}
		} else if len(inputValue) == 0 {
			fmt.Println("Please enter number or text again.")
			continue
		} else {
			fmt.Println("Something wrong, please try again.")
			continue
		}
	}
}
