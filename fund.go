package main

import (
	"encoding/json"
	"os"
)

type fund struct {
	Name     string    `json:"name"`
	Holdings []holding `json:"holdings"`
}

type fundData struct {
	FundData []fund `json:"fundData"`
}

func newFund(n string, h []holding) fund {
	return fund{
		Name:     n,
		Holdings: h,
	}
}

func (f fund) writeToFile(filename string) error {
	file, _ := json.MarshalIndent(f, "", " ")

	return os.WriteFile("fundData/"+filename, file, 0644)
}
