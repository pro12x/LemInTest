package main

func IsDuplicateTunnel(tunnel Tunnel, tunnels []Tunnel) bool {
	if len(tunnels) != 0 {
		for _, t := range tunnels {
			if t.From == tunnel.From && t.To == tunnel.To {
				return true
			}
		}
	}
	return false
}
