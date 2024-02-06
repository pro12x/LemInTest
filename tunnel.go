package main

import "fmt"

func FillTunnel(tab []string) (Tunnel, error) {
	if len(tab) != 0 {
		var result Tunnel
		if len(tab) == 2 {
			result.From = tab[0]
			result.To = tab[1]
			return result, nil
		}
	}
	return Tunnel{}, fmt.Errorf("invalid tunnel")
}
