package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type fund struct {
	Name     string    `json:"name"`
	Holdings []holding `json:"holdings"`
}

func getFundDataFromDir() []fund {
	files, err := os.ReadDir("fundData")
	if err != nil {
		fmt.Println("Error reading fund directory: ", err)
		os.Exit(1)
	}

	var fundData []fund
	for _, file := range files {
		fundData = append(fundData, newFundFromFile(file.Name()))
	}

	return fundData
}

func newFund(n string, h []holding) fund {
	return fund{
		Name:     n,
		Holdings: h,
	}
}

func newFundFromFile(filename string) fund {
	bs, err := os.ReadFile("fundData/" + filename)
	if err != nil {
		fmt.Println("Error reading fund from file: ", err)
		os.Exit(1)
	}

	fund := fund{}
	err = json.Unmarshal(bs, &fund)
	if err != nil {
		fmt.Println("Error unmarshalling fund: ", err)
		os.Exit(1)
	}

	return fund
}

func (f fund) writeToFile(filename string) error {
	file, _ := json.MarshalIndent(f, "", " ")

	return os.WriteFile("fundData/"+filename, file, 0644)
}

func extractCompaniesFromFundDataRecursively(fundData []fund, myHolding holding) []holding {
	var weightedCompanies []holding

	for _, fund := range fundData {
		if fund.Name == myHolding.Name {
			for _, h := range fund.Holdings {
				h.Weight *= myHolding.Weight
				var newCompanies []holding

				if strings.HasPrefix(h.Name, "Fund") {
					newCompanies = extractCompaniesFromFundDataRecursively(fundData, h)
				} else {
					newCompanies = append(newCompanies, h)
				}

				for _, newCompany := range newCompanies {
					i := FindHoldingInSlice(weightedCompanies, newCompany.Name)

					if i > -1 {
						weightedCompanies[i].Weight += newCompany.Weight
					} else {
						weightedCompanies = append(weightedCompanies, newCompany)
					}
				}
			}
		}
	}

	return weightedCompanies
}
