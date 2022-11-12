package main_test

import (
	"testing"

	m "recursive_data_fun"

	"github.com/stretchr/testify/assert"
)

func Test_NewHolding(t *testing.T) {
	result := m.Holding{
		Name:   "Holding Test",
		Weight: 1.0,
	}

	assert.IsType(t, m.Holding{}, result)
}

func Test_FindHoldingInSlice(t *testing.T) {
	holding1 := m.Holding{
		Name:   "Test Holding 1",
		Weight: 1.0,
	}

	holding2 := m.Holding{
		Name:   "Test Holding 2",
		Weight: 1.0,
	}

	holding3 := m.Holding{
		Name:   "Test Holding 3",
		Weight: 1.0,
	}

	holdingSlice := []m.Holding{
		holding1,
		holding2,
		holding3,
	}

	result := m.FindHoldingInSlice(holdingSlice, holding2.Name)

	assert.Equal(t, result, 1)
}
