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
