package main

import (
	"fmt"
	"math/rand"
)

func getRandomCoords(islandSize [2]int) (int, int) {
	randomPlantX := rand.Intn(islandSize[0])
	randomPlantY := rand.Intn(islandSize[1])
	return randomPlantX, randomPlantY
}

func updatePlants(plants [][]int, islandSizes int, growthRate int) [][]int {
	newPlants := make([][]int, len(plants))
	copy(newPlants, plants)

	for i := 0; i < islandSizes; i++ {
		for j := 0; j < islandSizes; j++ {
			if newPlants[i][j] > 0 {
				// Update current plant and spread to adjacent cells with boundary checks
				newPlants[i][j] += growthRate // Assuming growthRate should be an integer

				// Spread to adjacent cells with boundary checks
				if i > 0 { // Check upper boundary
					newPlants[i-1][j] += growthRate
				}
				if i < islandSizes-1 { // Check lower boundary
					newPlants[i+1][j] += growthRate
				}
				if j > 0 { // Check left boundary
					newPlants[i][j-1] += growthRate
				}
				if j < islandSizes-1 { // Check right boundary
					newPlants[i][j+1] += growthRate
				}
			}
		}
	}

	return plants
}
func main() {
	islandSizes := 20
	islandSize := [2]int{islandSizes, islandSizes}
	plantStart := 5
	growthRate := 1
	numberOfPlants := 5

	plants := make([][]int, islandSize[0])

	for i := range plants {
		plants[i] = make([]int, islandSize[1])
	}

	plantsLine := make([]int, 0, 5)

	for i := 0; i < numberOfPlants; i++ {
		randomPlantX, randomPlantY := getRandomCoords(islandSize)
		plants[randomPlantX][randomPlantY] = 1
	}

	plants = updatePlants(plants, islandSizes, growthRate)
	fmt.Println(plants, islandSize, plantStart, growthRate, plantsLine)
}
