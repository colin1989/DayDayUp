package aoi

type Entity struct {
	id  int // 唯一id
	pos Pos
	typ int // 类型
}

func (e *Entity) ID() int {
	return e.id
}

func (e *Entity) Type() int {
	return e.typ
}

func (e *Entity) Pos() Pos {
	return e.pos
}

func (e *Entity) IsWatcher() bool {
	return false
}

func (e *Entity) Notify(typ AOIEventType, entity AOIEntity) {

}

func (e *Entity) Refresh(typ EntityType, entities []AOIEntity) {

}
