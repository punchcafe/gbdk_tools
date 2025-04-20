package main

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"punchcafe.dev/gbdk_tools/tile"
)

func main() {

	existingImageFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	defer existingImageFile.Close()

	// Calling the generic image.Decode() will tell give us the data
	// and type of image it is as a string. We expect "png"
	imageData, err := png.Decode(existingImageFile)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	invalidWidth := imageData.Bounds().Max.X%8 != 0
	invalidHeight := imageData.Bounds().Max.Y%8 != 0
	if invalidHeight || invalidWidth {
		panic("Invalid image dimensions")
	}

	tile, err := tile.DefaultParser().ParseTile(imageData, image.Point{40, 40})
	if err != nil {
		panic("Failed to parse image")
	}
	fmt.Printf("Here's tile: %v", tile.ConvertToBinary())
}
