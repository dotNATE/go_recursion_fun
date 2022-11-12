package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	myHolding := NewHolding("Ethical Global Fund", 1)
	fundData := GetFundDataFromDir()
	myCompanies := ExtractCompaniesFromFundDataRecursively(fundData, myHolding)

	fmt.Print("\nYour holdings:\n\n")
	for _, holding := range myCompanies {
		fmt.Printf("Name: %s Weight: %.2f \n", holding.Name, holding.Weight)
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\nHow much do you have invested?")
	var investmentValue float32
	fmt.Scanln(&investmentValue)

	fmt.Println("\nWhich company do you want to check?")
	scanner.Scan()
	holdingName := scanner.Text()

	n := FindHoldingInSlice(myCompanies, holdingName)
	holdingExists := n > -1
	if !holdingExists {
		fmt.Println("\nCompany not found")
		os.Exit(1)
	}

	fmt.Printf("\nInvestment Value: £%.2f", investmentValue)
	fmt.Printf("\nHolding Name: %s", holdingName)

	for _, holding := range myCompanies {
		if holding.Name == holdingName {
			fmt.Printf("\nHolding Value: £%.2f", investmentValue*holding.Weight)
		}
	}
}
