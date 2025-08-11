package zin_mmo_project

import (
	"fmt"
	"testing"
)

func TestNewAoi(t *testing.T) {
	aoi := NewAoi(5, 5, 0, 250, 0, 250)
	println(aoi.PrintInfo())
}

func TestAoiManager_GetSurroundGridsByGridId(t *testing.T) {
	aoi := NewAoi(5, 5, 0, 250, 0, 250)
	grids := aoi.GetSurroundGridsByGridId(24)

	for i := 0; i < len(grids); i++ {
		g := grids[i]
		fmt.Print(g.Id, "---")
	}

}
