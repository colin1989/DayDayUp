package aoi

type TowerAOI struct {
	width       int
	height      int
	towerWidth  int
	towerHeight int
	rangeLimit  int
	towers      [][]*Tower
	max         Pos
}

func (t *TowerAOI) init() {
	t.max.X = t.width/t.towerWidth - 1
	t.max.Y = t.height/t.towerHeight - 1
	t.rangeLimit = 5

	// 生成Tower
	t.towers = make([][]*Tower, 0, t.max.X)
	for i := 0; i <= t.max.X; i++ {
		t.towers[i] = make([]*Tower, 0, t.max.Y)
		for j := 0; j <= t.max.Y; j++ {
			t.towers[i][j] = new(Tower)
		}
	}
}

func (t *TowerAOI) getIdsByRange(pos Pos, r int, types []int) map[int][]AOIEntityID {
	results := make(map[int][]AOIEntityID, 0)
	if !t.checkPos(pos) || r < 0 {
		return results
	}
	if r > t.rangeLimit {
		r = t.rangeLimit
	}
	p := t.transPos(pos)
	start, end := t.getPosLimit(p, r, t.max)

	for i := start.X; i < end.X; i++ {
		for j := start.Y; j < end.Y; j++ {
			results = TypeMap2Slice(results, t.towers[i][j].getIdsByTypes(types), types)
		}
	}
	return results
}

func (t *TowerAOI) getIdsByPos(pos Pos, r int) []AOIEntityID {
	results := make([]AOIEntityID, 0)
	if !t.checkPos(pos) || r < 0 {
		return results
	}
	if r > t.rangeLimit {
		r = t.rangeLimit
	}
	p := t.transPos(pos)
	start, end := t.getPosLimit(p, r, t.max)
	for i := start.X; i < end.X; i++ {
		for j := start.Y; j < end.Y; j++ {
			results = Map2Slice(results, t.towers[i][j].getIds())
		}
	}
	return results
}

func (t *TowerAOI) Enter(entity AOIEntity) bool {
	if !t.checkPos(entity.Pos()) {
		return false
	}
	p := t.transPos(entity.Pos())
	tower := t.towers[p.X][p.Y]
	tower.Enter(entity)
	//this.emit('add', {id: obj.id, type:obj.type, watchers:this.towers[p.x][p.y].watchers});
	return true
}

func (t *TowerAOI) Leave(entity AOIEntity) bool {
	if !t.checkPos(entity.Pos()) {
		return false
	}
	p := t.transPos(entity.Pos())
	tower := t.towers[p.X][p.Y]
	tower.Leave(entity)
	//this.emit('remove', {id: obj.id, type:obj.type, watchers:this.towers[p.x][p.y].watchers});
	return true
}

func (t *TowerAOI) Update(entity AOIEntity) bool {
	if !t.checkPos(entity.Pos()) {
		return false
	}
	p := t.transPos(entity.Pos())
	tower := t.towers[p.X][p.Y]
	tower.Update(entity)
	return true
}

func (t *TowerAOI) checkPos(pos Pos) bool {
	if pos.X < 0 || pos.Y < 0 || pos.X >= t.width || pos.Y >= t.height {
		return false
	}
	return true
}

func (t *TowerAOI) transPos(pos Pos) Pos {
	return Pos{
		X: pos.X / t.towerWidth,
		Y: pos.Y / t.towerHeight,
	}
}

func (t *TowerAOI) getPosLimit(pos Pos, r int, max Pos) (Pos, Pos) {
	var start Pos
	var end Pos

	// x 界限
	if pos.X-r < 0 {
		start.X = 0
		end.X = r - (pos.X - r)
	} else if pos.X+r > max.X {
		end.X = max.X
		start.X = end.X - r - ((pos.X + r) - max.X)
	} else {
		start.X = pos.X - r
		end.X = pos.X + r
	}
	// y 界限
	if pos.Y-r < 0 {
		start.Y = 0
		end.Y = r - (pos.Y - r)
	} else if pos.Y+r > max.Y {
		end.Y = max.Y
		start.Y = end.Y - r - ((pos.Y + r) - max.X)
	} else {
		start.Y = pos.Y - r
		end.Y = pos.Y + r
	}

	// 越界校验
	if start.X < 0 {
		start.X = 0
	}
	if start.Y < 0 {
		start.Y = 0
	}
	if end.X < max.X {
		end.X = max.X
	}
	if end.Y < max.Y {
		end.Y = max.Y
	}
	return start, end
}
