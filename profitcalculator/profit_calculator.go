package profitcalculator

import (
	"errors"
	"fmt"
	"os"
)

func ProfitCalculator() {
	revenue, err1 := getUserInput("Enter Revenue: ")
	expenses, err2 := getUserInput("Enter Expenses: ")
	taxRate, err3 := getUserInput("Enter Tax Rate: ")

	if err1 != nil || err2 != nil || err3 != nil {
		printErrors(err1, err2, err3)
		return
	}

	earningsBeforeTax, earningsAfterTax, ratio := calculateFinancials(revenue, expenses, taxRate)
	fmt.Println(earningsBeforeTax)
	fmt.Println(earningsAfterTax)
	fmt.Println(ratio)

	storeResults(earningsBeforeTax, earningsAfterTax, ratio)
}

func getUserInput(text string) (float64, error) {
	var userInputValue float64
	fmt.Print(text)
	fmt.Scan(&userInputValue)

	if userInputValue <= 0 {
		return 0, errors.New("value must be a positive number")
	}
	return userInputValue, nil
}

func outputText(text string, ebt, profit, ratio float64) {
	fmt.Printf(text, ebt, profit, ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningsBeforeTax := revenue - expenses
	earningsAfterTax := earningsBeforeTax * (1 - taxRate/100)
	ratio := earningsBeforeTax / earningsAfterTax
	return earningsBeforeTax, earningsAfterTax, ratio
}

func printErrors(errors ...error) {
	var errorsToPrint []error

	for _, err := range errors {
		if err != nil {
			errorsToPrint = append(errorsToPrint, err)
		}
	}

	if len(errorsToPrint) > 0 {
		fmt.Println("Errors:", errorsToPrint)
	}
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT is : %.1f\nProfit is : %.1f\nRatio is : %.3f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}
