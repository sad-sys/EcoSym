package main

import (
	"fmt"
	//"math/rand"
)

/*
func getRandomCoords(islandSize [2]int) (int, int) {
	randomPlantX := rand.Intn(islandSize[0])
	randomPlantY := rand.Intn(islandSize[1])
	return randomPlantX, randomPlantY
}
*/

func updatePlants(plants [][]int, islandSizes int, growthRate int) [][]int {
	newPlants := make([][]int, len(plants))
	for i := range plants {
		newPlants[i] = make([]int, len(plants[i]))
		copy(newPlants[i], plants[i])
	}

	for i := 0; i < islandSizes; i++ {
		for j := 0; j < islandSizes; j++ {
			if plants[i][j] > 0 {
				newPlants[i][j] += growthRate
				if i > 0 {
					newPlants[i-1][j] += growthRate
				}
				if i < islandSizes-1 {
					newPlants[i+1][j] += growthRate
				}
				if j > 0 {
					newPlants[i][j-1] += growthRate
				}
				if j < islandSizes-1 {
					newPlants[i][j+1] += growthRate
				}
			}
		}
	}

	return newPlants
}

func main() {
	islandSizes := 20
	islandSize := [2]int{islandSizes, islandSizes}
	growthRate := 1
	numberOfPlants := 1
	simLen := 3

	plants := make([][]int, islandSize[0])

	for i := range plants {
		plants[i] = make([]int, islandSize[1])
	}

	for i := 0; i < numberOfPlants; i++ {
		//randomPlantX, randomPlantY := getRandomCoords(islandSize)
		randomPlantX := 0
		randomPlantY := 0
		plants[randomPlantX][randomPlantY] = 1
	}

	// Size of the image
	for i := 0; i < simLen; i++ {
		plants = updatePlants(plants, islandSizes, growthRate)
	}
	for i := 0; i < islandSizes; i++ {
		for j := 0; j < islandSizes; j++ {
			fmt.Printf("%2d ", plants[i][j])
		}
		fmt.Println()
	}
}
