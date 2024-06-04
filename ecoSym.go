package main

import (
	"fmt"

	"math/rand"

	"image"
	"image/color"
	"image/gif"
	"os"
)

func getRandomCoords(islandSize [2]int) (int, int) {
	randomPlantX := rand.Intn(islandSize[0])
	randomPlantY := rand.Intn(islandSize[1])
	return randomPlantX, randomPlantY
}

func updatePlants(plants [][]int, islandSizes int, growthRate int) [][]int {
	newPlants := make([][]int, len(plants))
	for i := range plants {
		newPlants[i] = make([]int, len(plants[i]))
		copy(newPlants[i], plants[i])
	}

	for i := 0; i < islandSizes; i++ {
		for j := 0; j < islandSizes; j++ {
			if newPlants[i][j] > 0 {
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
	islandSizes := 500
	islandSize := [2]int{islandSizes, islandSizes}
	plantStart := 5
	growthRate := 1
	numberOfPlants := 5
	simLen := 500

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

	palette := []color.Color{
		color.RGBA{0x00, 0x00, 0xFF, 0xFF}, // Blue for lower concentration
		color.RGBA{0x00, 0x33, 0xFF, 0xFF},
		color.RGBA{0x00, 0x66, 0xFF, 0xFF},
		color.RGBA{0x00, 0x99, 0xFF, 0xFF}, // Light blue
		color.RGBA{0x00, 0xCC, 0xFF, 0xFF},
		color.RGBA{0x00, 0xFF, 0xFF, 0xFF}, // Cyan
		color.RGBA{0x00, 0xFF, 0xCC, 0xFF},
		color.RGBA{0x00, 0xFF, 0x99, 0xFF},
		color.RGBA{0x00, 0xFF, 0x66, 0xFF},
		color.RGBA{0x00, 0xFF, 0x33, 0xFF},
		color.RGBA{0x00, 0xFF, 0x00, 0xFF}, // Green
		color.RGBA{0x33, 0xFF, 0x00, 0xFF},
		color.RGBA{0x66, 0xFF, 0x00, 0xFF},
		color.RGBA{0x99, 0xFF, 0x00, 0xFF},
		color.RGBA{0xCC, 0xFF, 0x00, 0xFF},
		color.RGBA{0xFF, 0xFF, 0x00, 0xFF}, // Yellow
		color.RGBA{0xFF, 0xCC, 0x00, 0xFF},
		color.RGBA{0xFF, 0x99, 0x00, 0xFF}, // Orange
		color.RGBA{0xFF, 0x66, 0x00, 0xFF},
		color.RGBA{0xFF, 0x33, 0x00, 0xFF},
		color.RGBA{0xFF, 0x00, 0x00, 0xFF}, // Red for higher concentration
		color.RGBA{0xCC, 0x00, 0x00, 0xFF},
		color.RGBA{0x99, 0x00, 0x00, 0xFF},
		color.RGBA{0x66, 0x00, 0x00, 0xFF},
		color.RGBA{0x33, 0x00, 0x00, 0xFF},
		color.RGBA{0x00, 0x00, 0x00, 0xFF}, // Darker red
	}

	// Create an array of images
	images := []*image.Paletted{}
	delays := []int{}

	// Size of the image
	width, height := islandSizes, islandSizes

	for i := 0; i < simLen; i++ {
		// Create a new paletted image with the palette defined earlier
		img := image.NewPaletted(image.Rect(0, 0, width, height), palette)

		// Draw a colored rectangle
		maxPlantValue := 10000
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				// Calculate index based on the plant value scaled to the length of the palette
				plantValue := plants[y][x] // Retrieve the plant concentration value
				idx := int(float64(plantValue) / float64(maxPlantValue) * float64(len(palette)-1))
				if idx < 0 {
					idx = 0
				} else if idx >= len(palette) {
					idx = len(palette) - 1
				}
				c := palette[idx]
				img.Set(x, y, c) // Set the color at position x, y
			}
		}

		images = append(images, img)
		delays = append(delays, 0) // Delay between frames in 100ths of a second
		plants = updatePlants(plants, islandSizes, growthRate)
	}

	// Create the GIF
	f, err := os.Create("output.gif")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	gif.EncodeAll(f, &gif.GIF{
		Image: images,
		Delay: delays,
	})
}
