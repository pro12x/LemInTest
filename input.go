package main

import (
	"fmt"
	"strconv"
	"strings"
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

func Inputs(data []string) (Farm, error) {
	if len(data) != 0 {
		// Initialisation des variables
		var (
			err     error
			farm    Farm
			start   []Room
			end     []Room
			tunnel  Tunnel
			tab     []string
			isStart bool
			isEnd   bool
		)

		farm.Rooms = make(map[string]Room)
		farm.Ants, err = strconv.Atoi(data[0])

		if err != nil {
			return Farm{}, fmt.Errorf("cannot parse %v to int", data[0])
		}

		data = data[1:]

		for _, s := range data {
			s = strings.TrimSpace(s)
			if len(s) == 0 {
				continue
			}

			if s == "##start" {
				isEnd = false
				isStart = true
				continue
			}

			if s == "##end" {
				isStart = false
				isEnd = true
				continue
			}

			if strings.HasPrefix(s, "#") {
				continue
			}

			if strings.Contains(s, "-") {
				isEnd = false
				tab = strings.Fields(Replace(s, "-", " "))
			} else {
				tab = strings.Fields(s)
				fmt.Println(len(tab))
			}

			if len(tab) > 3 {
				return Farm{}, fmt.Errorf("invalid room")
			} else if len(tab) == 3 {
				farm.Rooms[tab[0]], err = FillRoom(tab)
			} else if len(tab) == 2 {
				tunnel, err = FillTunnel(tab)
				if err != nil {
					return Farm{}, err
				}
				farm.Tunnels = append(farm.Tunnels, tunnel)
			}

			fmt.Println(farm.Tunnels)

			if isStart {
				start = append(start, farm.Rooms[tab[0]])
			}

			if isEnd {
				end = append(end, farm.Rooms[tab[0]])
			}
			// fmt.Println(tab)
			// fmt.Println(farm.Rooms)
		}
		fmt.Println(start)
		fmt.Println(end)
		farm.Start = start[0]
		farm.End = end[0]

		return farm, nil
	}
	return Farm{}, fmt.Errorf("ERROR: invalid data format")
}
