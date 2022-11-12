package main

import "fmt"

func main() {
	myHolding := NewHolding("Ethical Global Fund", 1)
	fundData := GetFundDataFromDir()

	myCompanies := ExtractCompaniesFromFundDataRecursively(fundData, myHolding)
	var totalWeight float32

	fmt.Print("Your holdings:\n\n")

	for _, holding := range myCompanies {
		totalWeight += holding.Weight
		fmt.Printf("Name: %s Weight: %.2f \n", holding.Name, holding.Weight)
	}

	fmt.Printf("\nTotal Weight: %f", totalWeight)
}
