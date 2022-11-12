package main_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	m "recursive_data_fun"
)

func Test_NewFund(t *testing.T) {
	holdings := []m.Holding{
		{
			Name:   "Holding Test",
			Weight: 1,
		},
	}

	expected := m.Fund{
		Name:     "Fund Test",
		Holdings: holdings,
	}

	result := m.NewFund("Fund Test", holdings)

	assert.Equal(t, expected, result)
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
