package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func DisplayError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func ValidFile(file string) bool {
	if len(file) != 0 {
		reg1 := regexp.MustCompile(`^[a-zA-Z0-9_]+(?P<ext>.txt)$`)
		reg2 := regexp.MustCompile(`^[a-zA-Z0-9_]+(?P<ext>)$`)
		return reg1.MatchString(file) || reg2.MatchString(file)
	}
	return false
}

func OnlyAlphaNumeric(s string) bool {
	if len(strings.TrimSpace(s)) != 0 {
		cpt := 0
		for _, c := range strings.TrimSpace(s) {
			if c >= 97 && c <= 122 {
				cpt++
			}
		}
		if cpt == len(strings.TrimSpace(s)) {
			return true
		}
	}
	return false
}

func Sorting(paths [][]string) {
	index := 0
	for index < len(paths)-1 {
		if len(paths[index]) > len(paths[index+1]) {
			paths[index], paths[index+1] = paths[index+1], paths[index]
			index *= 0
		} else {
			index++
		}
	}
}

func IsExists(tab []string, s string) bool {
	if len(tab) != 0 {
		for _, v := range tab {
			if v == s {
				return true
			}
		}
	}
	return false
}

func ChangeFarmToPath(farm Farm, tab [][]string) []Path {
	if len(tab) != 0 {
		var paths []Path
		for _, path := range tab {
			// fmt.Println("Janel Path:", path)
			var rooms []Room
			for _, p := range path {
				if p != farm.Start.Name {
					rooms = append(rooms, farm.Rooms[p])
					// fmt.Println("Janel Room:", path)
				}
			}
			paths = append(paths, Path{Rooms: rooms})
		}
		return paths
	}
	return nil
}
