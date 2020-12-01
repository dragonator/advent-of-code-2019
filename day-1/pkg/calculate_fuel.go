package pkg

import (
	"fmt"
	"math"
	"strconv"
)

// CalculateModuleFuel -
func CalculateModuleFuel(mass int) int {
	return int(math.Floor(float64(mass)/3)) - 2
}

// CalculateFuelForFuel -
func CalculateFuelForFuel(fuelMass int) int {
	fuelForFuel := CalculateModuleFuel(fuelMass)
	if fuelForFuel <= 0 {
		return 0
	}
	return fuelForFuel + CalculateFuelForFuel(fuelForFuel)
}

// CalculateTotalFuel -
func CalculateTotalFuel(moduleMasses []string, modulesOnly bool) (int, error) {
	var totalFuel int
	for _, massStr := range moduleMasses {
		moduleMass, err := strconv.Atoi(massStr)
		if err != nil {
			return 0, fmt.Errorf("error: cannot convert value to integer: %s", err.Error())
		}

		fuelMass := CalculateModuleFuel(moduleMass)
		totalFuel += fuelMass

		if !modulesOnly {
			totalFuel += CalculateFuelForFuel(fuelMass)
		}
	}
	return totalFuel, nil
}
