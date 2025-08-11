package zin_mmo_project

import (
	"fmt"
	"sync"
)

// 网格结构体
type grid struct {
	Id        uint32
	MinX      int
	MaxX      int
	MinY      int
	MaxY      int
	PlayerIds map[uint32]bool
	PIdLock   sync.RWMutex
}

func NewGrid(id uint32, minX int, maxX int, minY int, maxY int) *grid {
	return &grid{
		Id:        id,
		MinX:      minX,
		MaxX:      maxX,
		MinY:      minY,
		MaxY:      maxY,
		PlayerIds: make(map[uint32]bool),
	}
}

// 添加玩家到网格中

func (g *grid) Add(playerId uint32) {
	g.PIdLock.Lock()
	defer g.PIdLock.Unlock()
	g.PlayerIds[playerId] = true
}

// Del 删除玩家
func (g *grid) Del(playerId uint32) {
	g.PIdLock.Lock()
	defer g.PIdLock.Unlock()
	delete(g.PlayerIds, playerId)
}

// GetOne 得到一个
func (g *grid) GetOne(playerId uint32) bool {
	g.PIdLock.RLock()
	defer g.PIdLock.RUnlock()
	_, ok := g.PlayerIds[playerId]
	return ok
}

// GetAllPlayId  获取所有玩家
func (g *grid) GetAllPlayId() []uint32 {
	g.PIdLock.RLock()
	defer g.PIdLock.RUnlock()
	var playerIds []uint32
	for id, ok := range g.PlayerIds {
		if ok {
			playerIds = append(playerIds, id)
		}
	}
	return playerIds
}

// Print 打印grid 信息
func (g *grid) Print() string {
	return fmt.Sprintf("Grid ID: %d, MinX: %d, MaxX: %d, MinY: %d, MaxY: %d,Plays:%v ", g.Id, g.MinX, g.MaxX, g.MinY, g.MaxY, g.PlayerIds)
}
