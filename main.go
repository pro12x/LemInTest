package main

import (
	"fmt"
	"os"
)

// Ant struct represents an ant
type Ant struct {
	UniqueID string
}

// Room struct represents a room in the ant farm
type Room struct {
	Name string
	X    int
	Y    int
}

// Path struct represents a path in the ant farm
type Path struct {
	Rooms []Room
	Ants  []Ant
}

// Tunnel struct represents a tunnel connecting two rooms in the ant farm
type Tunnel struct {
	From string
	To   string
}

// Farm struct represents the overall structure of the ant farm
type Farm struct {
	Ants    int  // OK - Number of Ants
	Start   Room // starting room
	End     Room
	Li      []string // ending room
	Roo     []Room
	Rooms   map[string]Room // All rooms
	Tunnels []Tunnel        // All tunnels
}

//=====================================================================

func main() {

	str := os.Args[1:]

	if len(str) == 1 && ValidFile(str[0]) {
		//data, err := Read(os.Args[1:][0])
		data, err := Read(str[0])
		if err != nil {
			DisplayError(err)
		}
		farm, err := Input(data)
		if err != nil {
			DisplayError(err)
		}
		paths := FillPath(farm)
		MovingAnts(paths, farm.Ants)
		os.Exit(0)
	}
	fmt.Println("USAGE: /lemin <filename>")
}
