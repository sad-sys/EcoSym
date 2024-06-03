package main

import (
	"fmt"
	"image/color"
	"image/png"
	"math/rand"
	"os"

	"github.com/fogleman/gg"
)

// Constants for island size and growth rate
const (
	islandSize   = 20
	growthRate   = 1
	updatePeriod = 10
)

// Generates a heatmap from a 20x20 grid of int values and saves it as a PNG.
func GenerateHeatmap(grid [][]int, filename string) error {
	const scale = 20 // Scale factor for each cell in the grid

	// Create a new image
	dc := gg.NewContext(islandSize*scale, islandSize*scale)

	// Set color gradient: from blue (cool) to red (hot)
	for y := 0; y < islandSize; y++ {
		for x := 0; x < islandSize; x++ {
			value := float64(grid[y][x])
			// Normalize the value to 0-1 for color mapping
			normValue := value / 255.0 // Assuming grid values are in the range [0, 255]
			r := uint8(255 * normValue)
			b := uint8(255 * (1 - normValue))
			color := color.RGBA{R: r, G: 0, B: b, A: 255}
			dc.SetColor(color)
			dc.DrawRectangle(float64(x*scale), float64(y*scale), float64(scale), float64(scale))
			dc.Fill()
		}
	}

	// Save the image to a file
	img := dc.Image()
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	png.Encode(file, img)

	return nil
}

func getRandomCoords() (int, int) {
	randomPlantX := rand.Intn(islandSize)
	randomPlantY := rand.Intn(islandSize)
	return randomPlantX, randomPlantY
}

func updatePlants(plants [][]int) [][]int {
	newPlants := make([][]int, len(plants))
	for i := range plants {
		newPlants[i] = make([]int, len(plants[i]))
		copy(newPlants[i], plants[i])
	}

	for i := 0; i < islandSize; i++ {
		for j := 0; j < islandSize; j++ {
			if newPlants[i][j] > 0 {
				// Update current plant and spread to adjacent cells with boundary checks
				newPlants[i][j] += growthRate // Assuming growthRate should be an integer

				// Spread to adjacent cells with boundary checks
				if i > 0 { // Check upper boundary
					newPlants[i-1][j] += growthRate
				}
				if i < islandSize-1 { // Check lower boundary
					newPlants[i+1][j] += growthRate
				}
				if j > 0 { // Check left boundary
					newPlants[i][j-1] += growthRate
				}
				if j < islandSize-1 { // Check right boundary
					newPlants[i][j+1] += growthRate
				}
			}
		}
	}
	return newPlants
}

func main() {
	// Initialize plants
	plants := make([][]int, islandSize)
	for i := range plants {
		plants[i] = make([]int, islandSize)
	}

	// Plant initial plants
	for i := 0; i < islandSize/2; i++ {
		randomPlantX, randomPlantY := getRandomCoords()
		plants[randomPlantX][randomPlantY] = 1
	}

	// Main loop
	for updates := 1; ; updates++ {
		// Update plants
		plants = updatePlants(plants)

		// Generate heatmap every 10 updates
		if updates%updatePeriod == 0 {
			filename := fmt.Sprintf("heatmap_update_%d.png", updates)
			err := GenerateHeatmap(plants, filename)
			if err != nil {
				panic(err)
			}
			fmt.Println("Generated heatmap for update", updates)
		}

		// Simulate delay (optional)
		// time.Sleep(time.Second)
	}
}
