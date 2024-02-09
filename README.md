# LeMin

LeMin est un programme conçu pour simuler une fourmilière et guider les fourmis d'un point de départ à un point d'arrivée à travers des tunnels et des salles.

---

## Installation

1. Clonez le dépôt GitHub sur votre machine locale :

2. Accédez au répertoire du projet :
3. Compilez le programme en utilisant Go :

---

## Utilisation

Exécutez le programme en spécifiant le nom du fichier contenant la description de la fourmilière :

---

## Fichier de description de la fourmilière

Le fichier de description de la fourmilière doit respecter un format spécifique. Consultez la documentation pour plus de détails.

---

## Fonctionnalités

- Lecture et traitement du fichier de description de la fourmilière.
- Création de la structure de la fourmilière.
- Simulation du mouvement des fourmis à travers les tunnels et les salles.
- Affichage des résultats de la simulation.


---

## Exemple d'utilisation

Voici un exemple d'utilisation du programme LeMin avec le fichier example00.txt :

```bash
$ go run . example00.txt
4
##start
0 0 3
2 2 5
3 4 0
##end
1 8 3
0-2
2-3
3-1
```
## Contributeurs

- Cheikh Mounirou Coly Diouf
- Frachis Janel Fmokomba
- Massar Samb
- Fatima Keita

---

## Licence

Ce projet est sous licence [Licence](LICENCE.md). Consultez le fichier LICENCE.md pour plus de détails.
