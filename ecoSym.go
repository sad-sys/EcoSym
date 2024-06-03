package main

import (
	"fmt"
	"image/color"
	"image/png"
	"math/rand"
	"os"

	"github.com/fogleman/gg"
)

// Generates a heatmap from a 20x20 grid of float64 values and saves it as a PNG.
func GenerateHeatmap(grid [][]int, filename string) error {
	const size = 20
	const scale = 20 // Scale factor for each cell in the grid

	// Create a new image
	dc := gg.NewContext(size*scale, size*scale)

	// Set color gradient: from blue (cool) to red (hot)
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
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

	err := GenerateHeatmap(plants, "heatmap.png")
	if err != nil {
		panic(err)
	}

	fmt.Println(plants, islandSize, plantStart, growthRate, plantsLine)
}
