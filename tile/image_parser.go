package tile

import (
	"errors"
	"image"
	"image/color"
)

type ImageParser struct {
	colorIndex [4]color.Color
}

func DefaultParser() *ImageParser {
	parser := ImageParser{
		[4]color.Color{
			color.RGBA{224, 248, 208, 255},
			color.RGBA{136, 192, 112, 255},
			color.RGBA{52, 104, 86, 255},
			color.RGBA{8, 24, 32, 255},
		},
	}
	return &parser
}

func (ip *ImageParser) ParseTile(Image image.Image, TopLeftCorner image.Point) (*Tile, error) {
	var bitPairs [TILE_WIDTH * TILE_HEIGHT]BitPair
	cursor := 0
	for x := TopLeftCorner.X; x < TopLeftCorner.X+TILE_WIDTH; x++ {
		for y := TopLeftCorner.Y; y < TopLeftCorner.Y+TILE_HEIGHT; y++ {
			bitPair, err := ip.mapColor(Image.At(x, y))
			if err != nil {
				return nil, err
			}
			bitPairs[cursor] = bitPair
			cursor++
		}
	}

	tile, err := FromBytePairArray(bitPairs)
	if err != nil {
		panic("unexpected failure, bitPairs should not be able to fail.")
	}
	return tile, nil
}

func (ip *ImageParser) mapColor(c color.Color) (BitPair, error) {
	for i, color := range ip.colorIndex {
		if c == color {
			return BitPair(i), nil
		}
	}

	return 0, errors.New("Unknown color found.")
}
