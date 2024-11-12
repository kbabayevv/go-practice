package bank

import "fmt"

func Bank() {
	var accountBalance = 1000.0
	fmt.Println("Welcome to PASHA BANK!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your balance is", accountBalance)
		} else if choice == 2 {
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid deposit amount. Must be greater than 0.")
			} else if depositAmount > 0 {
				accountBalance += depositAmount
				fmt.Println("Balance updated! New amount: ", accountBalance)
			}
		} else if choice == 3 {
			fmt.Print("Your withdraw money: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Invalid deposit amount. Must be greater than 0.")
			} else if withdrawAmount > accountBalance {
				fmt.Println("No insufficient funds.Please increase your balance!")
			} else if withdrawAmount <= accountBalance {
				accountBalance -= withdrawAmount
				fmt.Println("Balance updated! New amount: ", accountBalance)
			}
		} else {
			fmt.Println("Goodbye!")
			break
		}

		fmt.Println()
	}
}
