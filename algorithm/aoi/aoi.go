package aoi

type AOIEventType int

const (
	AOIEventEnter AOIEventType = iota
	AOIEventLeave
	AOIEventUpdate
)

type EntityType int

const (
	EntityTypePlayer EntityType = iota
	EntityTypeResource
	EntityTypeMonster
)

type AOIEntityID string

type Pos struct {
	X int
	Y int
}

type AOI interface {
	//
	// Enter
	//  @Description: 实例进入AOI范围
	//  @param e
	//  @return bool
	//
	Enter(e AOIEntity) bool
	//
	// Leave
	//  @Description: 实例离开AOI范围
	//  @param e
	//  @return bool
	//
	Leave(e AOIEntity) bool
	//
	// Update
	//  @Description: AOI范围内的实例更新
	//  @param e
	//  @return bool
	//
	Update(e AOIEntity) bool
	//
	// Refresh
	//  @Description: 按类型重置刷新实例
	//  @param typ
	//  @param e
	//
	Refresh(typ EntityType, e []AOIEntity)
	//
	// Watchers
	//  @Description: 获取AOI内的所有观察者
	//  @param typ
	//  @return map[int]AOIEntity
	//
	Watchers(typ int) map[AOIEntityID]AOIEntityID
}

type AOIEntity interface {
	ID() AOIEntityID
	Type() int
	Pos() Pos
	IsWatcher() bool
	Notify(typ AOIEventType, entity AOIEntity)
	Refresh(typ EntityType, entities []AOIEntity)
}
