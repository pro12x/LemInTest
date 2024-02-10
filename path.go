package main

import (
	"fmt"
	"os"
	"strings"
)

var PathValide [][]Room

func FillPath(farm Farm) []Path {
	CurrentPath := []Room{farm.Start}
	findPath(farm, farm.Roo, farm.Start, farm.End, farm.Li, CurrentPath, farm.Ants)
	PathValide = Doublons(PathValide, farm.Start, farm.End)
	pathss := Transform(PathValide)
	Sorting(pathss)
	if len(pathss) == 0 {
		fmt.Println("ERROR: invalid data format")
		os.Exit(0)
	}
	return ChangeFarmToPath(farm, pathss)
}

func findPath(Farms Farm, Rooms []Room, currentRoom Room, endRoom Room, link []string, currentPath []Room, antsLeft int) {
	if currentRoom.Name == endRoom.Name {
		// Nous avons trouvé un chemin valide jusqu'à la pièce d'arrivée
		PathValide = append(PathValide, append([]Room{}, currentPath...))
		return
	}
	// Obtenir les liens de la pièce actuelle
	linksOfCurrentRoom := RoomLink(currentRoom, Farms.Li)
	if len(linksOfCurrentRoom) != 0 {
		for _, linkedRoomName := range linksOfCurrentRoom {
			// Obtenir la pièce connectée
			linkedRoomName = strings.Split(linkedRoomName, "-")[1]
			linkedRoom := findRoomByName(Rooms, linkedRoomName)

			// Vérifier si la pièce connectée est déjà dans le chemin
			if !roomInPath(linkedRoom, currentPath) {

				// Vérifier si nous pouvons déplacer une fourmi à travers ce lien
				if antsLeft > 0 {
					// Mettre à jour le chemin actuel
					newPath := append(currentPath, linkedRoom)

					// Appeler récursivement pour explorer le chemin avec la pièce connectée
					findPath(Farms, Rooms, linkedRoom, endRoom, link, newPath, antsLeft-1)
				}
			}
		}
	}
}

// roomInPath vérifie si une pièce est déjà présente dans un chemin.
func roomInPath(room Room, path []Room) bool {
	for _, r := range path {
		if room.Name == r.Name {
			return true
		}
	}
	return false
}

// findRoomByName trouve une pièce par son nom dans la liste des pièces.
func findRoomByName(rooms []Room, name string) Room {
	for _, room := range rooms {
		if room.Name == name {
			return room
		}
	}
	return Room{} // Retourne une pièce vide si aucune correspondance n'est trouvée
}

// RoomLink retourne les liens d'une pièce spécifique.
func RoomLink(room Room, link []string) []string {
	var linkofroom []string
	for i := 0; i < len(link); i++ {
		parts := strings.Split(link[i], "-")
		if parts[0] == room.Name {
			linkofroom = append(linkofroom, link[i])
		} else if parts[1] == room.Name {
			linkofroom = append(linkofroom, parts[1]+"-"+parts[0])
		}
	}
	return linkofroom
}

func Doublons(path [][]Room, start, end Room) [][]Room {
	var Path [][]Room

	for j := 0; j < len(path); j++ {
		// Indicateur pour savoir si la liste de rooms doit être ajoutée à Path
		addToList := true

		// Parcourir les listes précédentes pour vérifier les doublons
		for k := 0; k < j; k++ {
			// Vérifier si les listes ont des éléments en commun
			if hasCommonElements(path[j], path[k], start, end) {
				path[j] = []Room{}
				addToList = false
				break
			}
		}

		// Si addToList est vrai, ajouter la liste de rooms à Path
		if addToList {
			Path = append(Path, path[j])
		}
	}

	return Path
}
func hasCommonElements(list1 []Room, list2 []Room, start Room, end Room) bool {
	// Utiliser une map pour garder une trace des noms de salles dans list1
	roomNames := make(map[string]bool)

	// Remplir la map avec les noms de salles de list1
	for _, room := range list1 {
		roomNames[room.Name] = true
	}

	// Vérifier si list2 a des éléments en commun avec list1
	for _, room := range list2 {
		if roomNames[room.Name] && end.Name != room.Name && start.Name != room.Name {
			return true
		}
	}

	// Aucun élément en commun trouvé
	return false
}

func Transform(paths [][]Room) [][]string {
	var Paths [][]string
	for _, i := range paths {
		path := []string{}
		for _, j := range i {
			path = append(path, j.Name)
		}
		Paths = append(Paths, path)
	}
	return Paths
}
