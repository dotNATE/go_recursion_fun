package main_test

import (
	"os"
	"testing"

	m "recursive_data_fun"

	"github.com/stretchr/testify/assert"
)

func Test_NewFund(t *testing.T) {
	holdings := []m.Holding{
		{
			Name:   "Holding Test",
			Weight: 1,
		},
	}

	result := m.NewFund("Fund Test", holdings)

	assert.IsType(t, m.Fund{}, result)
}

func Test_WriteToFileAndNewFundFromFile(t *testing.T) {
	os.Remove("fundData/_testfund.json")

	holdings := []m.Holding{
		{
			Name:   "Holding Test",
			Weight: 1,
		},
	}

	fund := m.Fund{
		Name:     "Fund Test",
		Holdings: holdings,
	}

	fund.WriteToFile("_testfund.json")

	result := m.NewFundFromFile("_testfund.json")

	assert.Equal(t, fund, result)

	os.Remove("fundData/_testfund.json")
}

func Test_GetFundDataFromDir(t *testing.T) {
	os.Remove("fundData/_testfund.json")

	holdings := []m.Holding{
		{
			Name:   "Holding Test",
			Weight: 1,
		},
	}

	fund := m.Fund{
		Name:     "Fund Test",
		Holdings: holdings,
	}

	fund.WriteToFile("_testfund.json")

	result := m.GetFundDataFromDir()

	assert.Contains(t, result, fund)

	os.Remove("fundData/_testfund.json")
}

func Test_ExtractCompaniesFromFundDataRecursively(t *testing.T) {
	myHolding := m.NewHolding("Fund 1", 1)

	companyHolding1 := m.NewHolding("Company 1", 0.5)
	companyHolding2 := m.NewHolding("Company 2", 0.3)
	companyHolding3 := m.NewHolding("Company 1", 0.7)
	fundHolding1 := m.NewHolding("Fund 2", 0.5)

	fund1 := m.NewFund("Fund 1", []m.Holding{
		companyHolding1,
		fundHolding1,
	})

	fund2 := m.NewFund("Fund 2", []m.Holding{
		companyHolding2,
		companyHolding3,
	})

	fundData := []m.Fund{
		fund1,
		fund2,
	}

	result := m.ExtractCompaniesFromFundDataRecursively(fundData, myHolding)

	assert.Equal(t, result[0], m.Holding{
		Name:   "Company 1",
		Weight: 0.85,
	})
	assert.Equal(t, result[1], m.Holding{
		Name:   "Company 2",
		Weight: 0.15,
	})
}
