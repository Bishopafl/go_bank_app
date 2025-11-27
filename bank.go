package main

import "fmt"

func main() {
	var accountBalance float64 = 1000.0

	fmt.Println("-----------------------------")
	fmt.Println("| Welcome to Adam's Go Bank! |")
	fmt.Println("-----------------------------")
	// let us loop through
	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Funds")
		fmt.Println("3. Withdraw Funds")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ") // show a new line
		fmt.Scan(&choice)          // pointer to choice var

		// This would usually be put inline - but this is here
		// just to show the that this is possible in Go
		wantsCheckBalance := choice == 1

		// Let's check what the user input is
		if wantsCheckBalance {
			fmt.Println("Your balance is", accountBalance)
		} else if choice == 2 {
			fmt.Println("-----------------------------")
			fmt.Print("Your Deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			// check the deposit amount
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				// return
				continue
			}

			accountBalance += depositAmount // accountBalance = accountBalance + depositAmount
			fmt.Println("Balance Updated! New Amount:", accountBalance)
			fmt.Println("-----------------------------")
		} else if choice == 3 {
			fmt.Println("-----------------------------")
			fmt.Print("Your Withdrawl: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			// withdraw functionality checks
			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				return
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. You cannot withdraw more than you have.")
				return
			}

			accountBalance -= withdrawAmount
			fmt.Println("Balance Updated! New Amount:", accountBalance)
		} else {
			fmt.Println("-----------------------------")
			fmt.Println("Goodbye!")
			break
		}
	}

	fmt.Println("Thanks for choosing our bank!")

}
