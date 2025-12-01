package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = getFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------")
		panic(err) // this shows more details for error handling like where the line it died was and a better error message view
	}

	fmt.Println("-----------------------------")
	fmt.Println("| Welcome to Adam's Go Bank! |")
	fmt.Println("-----------------------------")
	// let us loop through
	for {
		displayMenuOptions() // because the package main is pointing to the bank.go and communication.go - it will pick up the other communication.go file

		var choice int
		fmt.Print("Your choice: ") // show a new line
		fmt.Scan(&choice)          // pointer to choice var

		// This would usually be put inline - but this is here
		// just to show the that this is possible in Go
		checkBalance := 1
		wantsDeposit := 2
		wantsWithdraw := 3

		switch choice {
		case checkBalance:
			fmt.Println("Your balance is", accountBalance)
		case wantsDeposit:
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
			writeFloatToFile(accountBalance, accountBalanceFile)

			fmt.Println("Balance Updated! New Amount:", accountBalance)
			fmt.Println("-----------------------------")
		case wantsWithdraw:
			fmt.Println("-----------------------------")
			fmt.Print("Your Withdrawl: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			// withdraw functionality checks
			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. You cannot withdraw more than you have.")
				continue
			}

			accountBalance -= withdrawAmount
			writeFloatToFile(accountBalance, accountBalanceFile)

			fmt.Println("Balance Updated! New Amount:", accountBalance)
		default:
			fmt.Println("-----------------------------")
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank!")
			return
		}
	}

}

func writeFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(accountBalanceFile, []byte(valueText), 0644) // text name, slice of bytes, permissions
}

func getFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("failed to find file")
	}

	valueText := string(data)
	balance, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored value")
	}

	return balance, nil
}
