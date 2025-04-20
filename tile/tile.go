package tile

// 0 - 3
const TILE_WIDTH = 8
const TILE_HEIGHT = 8
const TILE_SIZE_BYTES = (TILE_WIDTH * TILE_HEIGHT * 2) / 8

/*
Represents a single pixel brightness value between 0 - 3 inclusive.
*/
type BitPair uint8

type Tile [TILE_WIDTH][TILE_HEIGHT]BitPair

// convert to it's byte data
func (t *Tile) ConvertToBinary() [TILE_SIZE_BYTES]byte {
	var out [TILE_SIZE_BYTES]byte
	byteIndex := 0
	for _, pixelRow := range t {
		var lhsByte, rhsByte byte = 0x0, 0x0
		for pixelPosition, pixelValue := range pixelRow {
			lhs, rhs := bytePairAtPosition(uint(pixelPosition), pixelValue)
			lhsByte = lhs | lhsByte
			rhsByte = rhs | rhsByte
		}
		out[byteIndex] = lhsByte
		out[byteIndex+1] = rhsByte
		byteIndex += 2
	}
	return out
}

type bytePair struct {
	lhs byte
	rhs byte
}

var valuePairLookup = [4]bytePair{
	{0x00, 0x00},
	{0x01, 0x00},
	{0x00, 0x01},
	{0x01, 0x01},
}

func bytePairAtPosition(position uint, value BitPair) (byte, byte) {
	bytePair := valuePairLookup[value]
	return bytePair.lhs << byte(7-position), bytePair.rhs << byte(7-position)
}
