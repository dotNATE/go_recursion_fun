package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Fund struct {
	Name     string    `json:"name"`
	Holdings []Holding `json:"holdings"`
}

func ExtractCompaniesFromFundDataRecursively(fundData []Fund, myHolding Holding) []Holding {
	var weightedCompanies []Holding

	for _, fund := range fundData {
		if fund.Name == myHolding.Name {
			for _, h := range fund.Holdings {
				h.Weight *= myHolding.Weight
				var newCompanies []Holding

				if strings.HasPrefix(h.Name, "Fund") {
					newCompanies = ExtractCompaniesFromFundDataRecursively(fundData, h)
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

func GetFundDataFromDir() []Fund {
	files, err := os.ReadDir("fundData")
	if err != nil {
		fmt.Println("Error reading fund directory: ", err)
		os.Exit(1)
	}

	var fundData []Fund
	for _, file := range files {
		fundData = append(fundData, NewFundFromFile(file.Name()))
	}

	return fundData
}

func NewFund(n string, h []Holding) Fund {
	return Fund{
		Name:     n,
		Holdings: h,
	}
}

func NewFundFromFile(filename string) Fund {
	bs, err := os.ReadFile("fundData/" + filename)
	if err != nil {
		fmt.Println("Error reading fund from file: ", err)
		os.Exit(1)
	}

	fund := Fund{}
	err = json.Unmarshal(bs, &fund)
	if err != nil {
		fmt.Println("Error unmarshalling fund: ", err)
		os.Exit(1)
	}

	return fund
}

func (f Fund) WriteToFile(filename string) error {
	file, _ := json.MarshalIndent(f, "", " ")

	return os.WriteFile("fundData/"+filename, file, 0644)
}
