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
