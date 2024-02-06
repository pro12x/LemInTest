package main

import (
	"fmt"
	"strconv"
	"strings"
)

func FillRoom(tab []string) (Room, error) {
	if len(tab) != 0 {
		var result Room
		var err error
		if len(tab) == 3 {
			if strings.HasPrefix(tab[0], "L") {
				return Room{}, fmt.Errorf("room cannot begin with 'L'")
			}
			result.Name = tab[0]
			result.X, err = strconv.Atoi(tab[1])
			if err != nil {
				return Room{}, fmt.Errorf("cannot parse %v to int", tab[1])
			}
			result.Y, err = strconv.Atoi(tab[2])
			if err != nil {
				return Room{}, fmt.Errorf("cannot parse %v to int", tab[2])
			}
			return result, nil
		}
	}
	return Room{}, fmt.Errorf("invalid room")
}
