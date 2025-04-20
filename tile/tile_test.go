package tile

import (
	"testing"
)

func Test_tileRendersBlackTile(t *testing.T) {
	blackSample := Tile{
		{3, 3, 3, 3, 3, 3, 3, 3},
		{3, 3, 3, 3, 3, 3, 3, 3},
		{3, 3, 3, 3, 3, 3, 3, 3},
		{3, 3, 3, 3, 3, 3, 3, 3},
		{3, 3, 3, 3, 3, 3, 3, 3},
		{3, 3, 3, 3, 3, 3, 3, 3},
		{3, 3, 3, 3, 3, 3, 3, 3},
		{3, 3, 3, 3, 3, 3, 3, 3},
	}

	expectedResponse := [16]byte{
		255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255,
	}

	if expectedResponse != blackSample.ConvertToBinary() {
		t.Fail()
	}
}

func Test_tileRendersWhiteTile(t *testing.T) {
	whiteSample := Tile{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}

	expectedResponse := [16]byte{
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
	}

	actualResponse := whiteSample.ConvertToBinary()

	if expectedResponse != actualResponse {
		t.Errorf("Failed to match: %v", actualResponse)
		t.Fail()
	}
}

func Test_tileRendersLightGreyTile(t *testing.T) {
	lightGreySample := Tile{
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
	}

	expectedResponse := [16]byte{
		255, 0, 255, 0, 255, 0, 255, 0,
		255, 0, 255, 0, 255, 0, 255, 0,
	}

	actualResponse := lightGreySample.ConvertToBinary()

	if expectedResponse != actualResponse {
		t.Errorf("Failed to match: %v", actualResponse)
		t.Fail()
	}
}

func Test_tileRendersDarkGreyTile(t *testing.T) {
	darkGreySample := Tile{
		{2, 2, 2, 2, 2, 2, 2, 2},
		{2, 2, 2, 2, 2, 2, 2, 2},
		{2, 2, 2, 2, 2, 2, 2, 2},
		{2, 2, 2, 2, 2, 2, 2, 2},
		{2, 2, 2, 2, 2, 2, 2, 2},
		{2, 2, 2, 2, 2, 2, 2, 2},
		{2, 2, 2, 2, 2, 2, 2, 2},
		{2, 2, 2, 2, 2, 2, 2, 2},
	}

	expectedResponse := [16]byte{
		0, 255, 0, 255, 0, 255, 0, 255,
		0, 255, 0, 255, 0, 255, 0, 255,
	}

	actualResponse := darkGreySample.ConvertToBinary()

	if expectedResponse != actualResponse {
		t.Errorf("Failed to match: %v", actualResponse)
		t.Fail()
	}
}

func Test_tileRendersComplexTile(t *testing.T) {
	darkGreySample := Tile{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 0, 1, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 2, 1, 2, 1, 2, 1, 2},
		{2, 2, 2, 2, 2, 2, 2, 2},
		{2, 3, 2, 3, 2, 3, 2, 3},
		{3, 3, 3, 3, 3, 3, 3, 3},
		{3, 3, 3, 3, 3, 3, 3, 3},
	}

	expectedResponse := [16]byte{0, 0, 170, 0, 255, 0, 85, 170, 0, 255, 170, 255, 255, 255, 255, 255}

	actualResponse := darkGreySample.ConvertToBinary()

	if expectedResponse != actualResponse {
		t.Errorf("Failed to match: %v", actualResponse)
		t.Fail()
	}
}

func Test_tileRendersSmileyFaceTile(t *testing.T) {
	smileyFaceSample := Tile{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 1, 0, 0, 0, 0, 0},
		{0, 1, 1, 0, 0, 0, 0, 0},
		{0, 1, 1, 0, 0, 1, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 1, 0, 0, 1, 1, 0},
		{0, 0, 1, 1, 1, 1, 0, 0},
	}

	expectedResponse := [16]byte{0, 0, 0, 0, 6, 0, 6, 0, 102, 0, 0, 0, 102, 0, 60, 0}

	actualResponse := smileyFaceSample.ConvertToBinary()

	if expectedResponse != actualResponse {
		t.Errorf("Failed to match: %v", actualResponse)
		t.Fail()
	}
}
