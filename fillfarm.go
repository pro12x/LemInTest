package main

import (
	"fmt"
	"strconv"
	"strings"
)

func FillFarm(tab []string) (Farm, error) {
	if len(tab) != 0 {
		var (
			err     error
			farm    Farm
			rooms   = make(map[string]Room)
			slice   []string
			isStart bool
			isEnd   bool
			room    Room
			start   Room
			end     Room
			tunnel  Tunnel
			tunnels []Tunnel
			ants    int
		)

		ants, err = strconv.Atoi(tab[0])
		if err != nil {
			return Farm{}, fmt.Errorf("ERROR: invalid data format\n")
		}
		tab = tab[1:]
		// fmt.Println("Janel ants:", ants)
		// farm.Ants = ants

		for _, str := range tab {
			str = strings.TrimSpace(str)
			if len(str) == 0 {
				continue
			}

			if str == "##start" {
				isStart = true
				continue
			}

			if str == "##end" {
				isEnd = true
				continue
			}

			if strings.HasPrefix(str, "#") {
				continue
			}

			if strings.Contains(str, "-") {
				slice = strings.Fields(Replace(str, "-", " "))
			} else {
				slice = strings.Fields(str)
			}

			if len(slice) > 3 {
				return Farm{}, fmt.Errorf("ERROR: invalid data format\n")
			} else if len(slice) == 3 {
				room, err = FillRoom(slice)
				if err != nil {
					return Farm{}, fmt.Errorf("ERROR: invalid data format\n")
				}
				rooms[room.Name] = room

				if isStart {
					start = room
					isStart = false
					// fmt.Println("Janel start:", start)
					continue
				}

				if isEnd {
					end = room
					isStart = false
					// fmt.Println("Janel end:", end)
					continue
				}
				continue
			} else if len(slice) == 2 {
				tunnel, err = FillTunnel(slice)
				// fmt.Println("Janel Tunnel:", tunnel)
				if err != nil {
					return Farm{}, fmt.Errorf("ERROR: invalid data format\n")
				}
				// Check if tunnel exists
				_, to := rooms[tunnel.To]
				_, from := rooms[tunnel.From]

				if to && from {
					if !IsDuplicateTunnel(tunnel, tunnels) {
						tunnels = append(tunnels, tunnel)
					}
					// fmt.Println("Janel Tunnels:", tunnels)
				} else {
					return Farm{}, fmt.Errorf("ERROR: invalid data format\n")
				}
			}
		}
		farm = Farm{
			Ants:    ants,
			Start:   start,
			End:     end,
			Rooms:   rooms,
			Tunnels: tunnels,
		}
		return farm, nil
	}
	return Farm{}, fmt.Errorf("ERROR: invalid data format\n")
}
