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

func sumPlants(plants [][]int) int {
	count := 0
	for i := 0; i < len(plants); i++ {
		for j := 0; j < len(plants); j++ {
			count = count + plants[i][j]
		}
	}
	fmt.Println("#Sum of All plants ", count)
	return count

}

func makeFood(islandSizes int) [][]int {
	food := make([][]int, islandSizes)
	for i := range food {
		food[i] = make([]int, islandSizes)
		for j := range food[i] {
			food[i][j] = 10
		}
	}

	for i := 0; i < islandSizes; i++ {
		for j := 0; j < islandSizes; j++ {
			fmt.Printf("%2d ", food[i][j])
		}
		fmt.Println()

	}
	return food
}

func updatePlants(plants [][]int, islandSize int, food [][]int) [][]int {

	growthRate := 5
	newPlants := make([][]int, len(plants))

	for i := 0; i < islandSize; i++ {
		newPlants[i] = make([]int, len(plants[i]))
		copy(newPlants[i], plants[i])
	}

	for i := 0; i < islandSize; i++ {
		for j := 0; j < islandSize; j++ {
			if plants[i][j] > 0 {
				if food[i][j] > 0 {
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
					food[i][j] -= 1
				}
			}
		}
	}

	return newPlants
}

func main() {
	islandSize := 50
	numberOfPlants := 1
	simLen := 500

	plants := make([][]int, islandSize)

	for i := 0; i < islandSize; i++ {
		plants[i] = make([]int, islandSize)
	}

	for i := 0; i < numberOfPlants; i++ {
		randomX, randomY := getRandomCoords(islandSize)
		fmt.Println(randomX, randomY)
		plants[randomY][randomX] = 1
	}

	food := makeFood(islandSize)

	for i := 0; i < simLen; i++ {
		plants = updatePlants(plants, islandSize, food)
		for i := 0; i < islandSize; i++ {
			for j := 0; j < islandSize; j++ {
				fmt.Printf("%2d ", plants[i][j])
			}
			fmt.Println()
		}
		fmt.Println("############################################################################")
	}

	fmt.Println(food)

	sumPlants(plants)

}
