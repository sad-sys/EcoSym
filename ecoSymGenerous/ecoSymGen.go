package main

import (
	"fmt"
	"math/rand"
)

func getRandomCoords(islandSize int) (int, int) {
	randomX := rand.Intn(islandSize - 1)
	randomY := rand.Intn(islandSize - 1)
	return randomX, randomY
}

func sumPlants(plants [][]int) int{
	count := 0
	for i:= 0; i < len(plants); i++{
		for j:= 0; j < len(plants); j++{
			count = count + plants[i][j]
		}
	}
	fmt.Println("#Sum of All plants ", count)
	return count 

} 

func updatePlants(plants [][]int, islandSize int) [][]int {

	growthRate := 5
	newPlants := make([][]int, len(plants))

	for i := 0; i < islandSize; i++ {
		newPlants[i] = make([]int, len(plants[i]))
		copy(newPlants[i], plants[i])
	}

	for i := 0; i < islandSize; i++ {
		for j := 0; j < islandSize; j++ {
			if plants[i][j] > 0 {
				newPlants[i][j] += growthRate / 5
				if i > 0 {
					newPlants[i-1][j] += growthRate + growthRate/5
				}
				if i < islandSize-1 {
					newPlants[i+1][j] += growthRate + growthRate/5
				}
				if j > 0 {
					newPlants[i][j-1] += growthRate + growthRate/5
				}
				if j < islandSize-1 {
					newPlants[i][j+1] += growthRate + growthRate/5
				}
			}
		}
	}

	return newPlants
}

func main() {
	islandSize := 20
	numberOfPlants := 1
	simLen := 3

	plants := make([][]int, islandSize)

	for i := 0; i < islandSize; i++ {
		plants[i] = make([]int, islandSize)
	}

	for i := 0; i < numberOfPlants; i++ {
		randomX, randomY := getRandomCoords(islandSize)
		fmt.Println(randomX, randomY)
		plants[randomY][randomX] = 1
	}

	for i := 0; i < simLen; i++ {
		plants = updatePlants(plants, islandSize)
		for i := 0; i < islandSize; i++ {
			for j := 0; j < islandSize; j++ {
				fmt.Printf("%2d ", plants[i][j])
			}
			fmt.Println()
		}
		fmt.Println("############################################################################")
	}

	sumPlants(plants)
}
