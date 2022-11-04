package main

import (
	"encoding/json"
	"fmt"
	"os"
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
