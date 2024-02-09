package main

import (
	"fmt"
	"os"
)

type Location struct {
	posPath int
	posRoom int
}

func MovingAnts(paths []Path, ants int) {
	if len(paths) != 0 {
		path := 0
		location := make(map[string]*Location)
		for i := 0; i < ants; i++ {
			next := (path + 1) % len(paths)
			ant := Ant{
				UniqueID: fmt.Sprintf("%v", i+1),
			}
			pathValue := len(paths[path].Ants) + len(paths[path].Rooms)
			nextValue := len(paths[next].Ants) + len(paths[next].Rooms)

			if nextValue <= pathValue {
				paths[next].Ants = append(paths[next].Ants, ant)
				path = next
			} else {
				paths[path].Ants = append(paths[path].Ants, ant)
			}
			i := len(paths[path].Ants) - 1
			if i != 0 {
				i = -1 * i
			}
			location[ant.UniqueID] = &Location{
				path,
				i,
			}
		}
		var str string
		for {
			// Nous supposons que toutes les fourmis sont déjà passées
			check := true
			// Parcours selon les positions des fourmis
			for id, pos := range location {
				if pos.posRoom >= 0 && pos.posRoom < len(paths[pos.posPath].Rooms) {
					// Récupérer la salle dans laquelle se trouve la fourmi
					room := paths[pos.posPath].Rooms[pos.posRoom]
					str += "L" + id + "-" + room.Name + " "
					// fmt.Println("Janel Movement:", str)
					pos.posRoom++
					// Il y a encore d'autre fourmis
					check = false
				} else {
					location[id].posRoom++
				}
			}
			if check {
				break
			} else {
				str += string(rune(10))
			}
		}
		fmt.Println(str)
		os.Exit(0)
	}
	fmt.Println("ERROR: invalid data format")
}
