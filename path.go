package main

func FillPath(farm Farm) []Path {
	visitedRoom := make(map[string]bool)
	var paths [][]string
	var funcPath func(room string, path []string)

	funcPath = func(room string, path []string) {
		visitedRoom[room] = true
		path = append(path, room)

		if room == farm.End.Name {
			anotherPath := make([]string, len(path))
			copy(anotherPath, path)

			pathIsShared := false
			pathID := 0

			for i, p := range paths {
				for _, r := range anotherPath {
					if (r != farm.Start.Name && r != farm.End.Name) && IsExists(p, r) {
						pathIsShared = true
						pathID = i
						break
					}
				}
				if pathIsShared {
					break
				}
			}

			if !pathIsShared {
				paths = append(paths, anotherPath)
				// Test
				pathID = len(paths) - 1
			}
			if len(paths[pathID]) > len(anotherPath) {
				paths[pathID] = anotherPath
			}
		} else {
			for _, tunnel := range farm.Tunnels {
				if tunnel.From == room && !visitedRoom[tunnel.To] {
					funcPath(tunnel.To, path)
				}
				if tunnel.To == room && !visitedRoom[tunnel.From] {
					funcPath(tunnel.From, path)
				}
			}
		}
		visitedRoom[room] = false
	}
	funcPath(farm.Start.Name, []string{})

	return ChangeFarmToPath(farm, paths)
}
