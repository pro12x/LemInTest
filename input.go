package main

import (
	"fmt"
)

func Input(data []string) (Farm, error) {
	if len(data) != 0 {
		var (
			err  error
			farm Farm
		)
		farm, err = FillFarm(data)
		if err != nil {
			return Farm{}, fmt.Errorf("ERROR: invalid data format\n")
		}

		return farm, err
	}

	return Farm{}, fmt.Errorf("ERROR: invalid data format\n")
}
