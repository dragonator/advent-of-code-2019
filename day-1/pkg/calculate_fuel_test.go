package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CalculateModuleFuel(t *testing.T) {
	res := CalculateModuleFuel(12)
	assert.Equal(t, 2, res)

	res = CalculateModuleFuel(14)
	assert.Equal(t, 2, res)

	res = CalculateModuleFuel(1969)
	assert.Equal(t, 654, res)

	res = CalculateModuleFuel(100756)
	assert.Equal(t, 33583, res)
}

func Test_CalculateFuelForFuel(t *testing.T) {
	res := CalculateFuelForFuel(2)
	assert.Equal(t, 0, res)

	res = CalculateFuelForFuel(654)
	assert.Equal(t, 312, res)

	res = CalculateFuelForFuel(33583)
	assert.Equal(t, 16763, res)
}

func Test_CalculateTotalFuel(t *testing.T) {
	res, err := CalculateTotalFuel([]string{"12", "14", "1969", "100756"}, true)
	assert.NoError(t, err)
	assert.Equal(t, 34241, res)
}
