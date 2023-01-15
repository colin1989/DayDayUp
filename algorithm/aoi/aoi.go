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

type Pos struct {
	X int
	Y int
}

type AOI interface {
	Enter(e AOIEntity) bool
	Leave(e AOIEntity) bool
	Update(e AOIEntity) bool
	Refresh(e []AOIEntity, typ EntityType)
	Watchers(typ int) map[int]AOIEntity
}

type AOIEntity interface {
	ID() int
	Type() int
	Pos() Pos
	Notify(t AOIEntity, typ AOIEventType)
	IsWatcher() bool
	Refresh(e []AOIEntity, typ EntityType)
}
