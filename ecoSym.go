package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"math/rand"
	"os"
)

func getRandomCoords(islandSize [2]int) (int, int) {
	randomPlantX := rand.Intn(islandSize[0])
	randomPlantY := rand.Intn(islandSize[1])
	return randomPlantX, randomPlantY
}

func makePlants(islandSizes int) [][]int {

	plants := make([][]int, islandSizes)
	for i := range plants {
		plants[i] = make([]int, islandSizes)
	}
	return plants
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

func updatePlants(plants [][]int, islandSizes int, growthRate int, food [][]int) [][]int {
	newPlants := make([][]int, len(plants))
	for i := range plants {
		newPlants[i] = make([]int, len(plants[i]))
		copy(newPlants[i], plants[i])
	}

	for i := 0; i < islandSizes; i++ {
		for j := 0; j < islandSizes; j++ {
			if plants[i][j] > 0 {
				if food[i][j] > 0 {
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
					food[i][j] -= 1
				}
			}
		}
	}

	return newPlants
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

func createHeatmap(plants [][]int, maxValue int) *image.Paletted {
	height := len(plants)
	width := len(plants[0])
	img := image.NewGray(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			value := plants[y][x]
			grayValue := uint8(255 * value / maxValue)
			img.SetGray(x, y, color.Gray{Y: grayValue})
		}
	}

	// Convert image to paletted format for GIF
	var palette []color.Color
	for i := 0; i < 256; i++ {
		palette = append(palette, color.Gray{uint8(i)})
	}
	palettedImage := image.NewPaletted(img.Rect, palette)
	draw.FloydSteinberg.Draw(palettedImage, img.Bounds(), img, image.Point{})

	return palettedImage
}

func saveGIF(frames []*image.Paletted, filename string) {
	outGif := &gif.GIF{}
	for _, img := range frames {
		outGif.Image = append(outGif.Image, img)
		outGif.Delay = append(outGif.Delay, 10) // Adjust delay as needed
	}

	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error: could not create GIF file.")
		return
	}
	defer f.Close()

	gif.EncodeAll(f, outGif)
}

func main() {
	islandSizes := 50
	islandSize := [2]int{islandSizes, islandSizes}
	growthRate := 5
	numberOfPlants := 1
	simLen := 500
	maxValue := 100

	plants := makePlants(islandSizes)

	food := makeFood(islandSizes)

	for i := 0; i < numberOfPlants; i++ {
		randomPlantX, randomPlantY := getRandomCoords(islandSize)
		plants[randomPlantX][randomPlantY] = 1
	}

	frames := []*image.Paletted{}

	for i := 0; i < simLen; i++ {
		plants = updatePlants(plants, islandSizes, growthRate, food)

		for i := 0; i < islandSizes; i++ {
			for j := 0; j < islandSizes; j++ {
				fmt.Printf("%2d ", plants[i][j])
			}
			fmt.Println()
		}
		fmt.Println("############################################################################")

		// Generate and save heatmap for each step
		heatmap := createHeatmap(plants, maxValue)
		frames = append(frames, heatmap)
	}

	fmt.Println(food)
	// Save the frames as a GIF
	saveGIF(frames, "heatmap1.gif")
	sumPlants(plants)
}
