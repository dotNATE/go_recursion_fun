package main_test

import (
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

	assert.Equal(t, expected, result, "Both values should be equal")
}
