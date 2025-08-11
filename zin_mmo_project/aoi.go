package zin_mmo_project

import "fmt"

type AoiManager struct {
	MinX int
	MaxX int
	MinY int
	MaxY int
	// x方向网格数量
	CountX int
	// y方向网格数量
	CountY int
	grids  map[uint32]*grid
}

func NewAoi(cx, cy, minX, maxX, minY, maxY int) *AoiManager {
	manager := &AoiManager{
		MinX:   minX,
		MaxX:   maxX,
		MinY:   minY,
		MaxY:   maxY,
		CountX: cx,
		CountY: cy,
		grids:  make(map[uint32]*grid),
	}

	for i := 0; i < cy; i++ {
		for j := 0; j < cx; j++ {
			gridId := uint32(i*cx + j)
			newGrid := NewGrid(gridId, minX+j*manager.GetXWidth(), minX+(j+1)*manager.GetXWidth(),
				minY+i*manager.GetYWidth(), minY+(i+1)*manager.GetYWidth())
			manager.grids[gridId] = newGrid
		}
	}
	return manager
}

func (am *AoiManager) GetXWidth() int {
	return (am.MaxX - am.MinX) / am.CountX
}
func (am *AoiManager) GetYWidth() int {
	return (am.MaxY - am.MinY) / am.CountY
}

func (am *AoiManager) PrintInfo() string {
	s := fmt.Sprintf("AoiManager: MinX:%d,MaxX:%d,MinY:%d,MaxY:%d\n", am.MinX, am.MaxX, am.MinY, am.MaxY)
	for _, grid := range am.grids {
		s += grid.Print() + "\n"
	}

	return s
}

// GetSurroundGridsByGridId 求周围的格子（包含自己）
func (am *AoiManager) GetSurroundGridsByGridId(griId int) (res []*grid) {
	var resGridIds []int
	if griId < 0 {
		return
	}
	if _, ok := am.grids[uint32(griId)]; !ok {
		return
	}
	indexX := griId % am.CountX // 当前grid的x坐标
	var griIds []int
	griIds = append(griIds, griId)
	if indexX > 0 {
		//左边有格子
		griIds = append(griIds, griId-1)
	}
	if indexX < am.CountX-1 {
		//右边有格子
		griIds = append(griIds, griId+1)
	}
	resGridIds = append(resGridIds, griIds...)
	for _, v := range griIds {
		indexY := v / am.CountY
		if indexY > 0 {
			//上边有格子
			resGridIds = append(resGridIds, v-am.CountX)
		}
		if indexY < am.CountY-1 {
			resGridIds = append(resGridIds, v+am.CountX)
		}
	}
	for _, v := range resGridIds {
		grid := am.grids[uint32(v)]
		res = append(res, grid)
	}
	return

}

func (am *AoiManager) GetSurroundPlayIdsByPos(x, y float32) (res []uint32) {
	gridId := am.GetGridIdByPos(x, y)
	grids := am.GetSurroundGridsByGridId(int(gridId))
	for _, grid := range grids {
		res = append(res, grid.GetAllPlayId()...)
	}
	return
}

func (am *AoiManager) GetGridIdByPos(x float32, y float32) uint32 {
	indexX := int(x) / am.GetXWidth()
	indexY := int(y) / am.GetXWidth()
	return uint32(indexY*am.CountX + indexX)
}
