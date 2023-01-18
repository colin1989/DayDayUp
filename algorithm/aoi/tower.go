package aoi

type Tower struct {
	count    int
	ids      map[AOIEntityID]AOIEntityID
	idsType  map[int]map[AOIEntityID]AOIEntityID
	watchers map[AOIEntityID]AOIEntityID
}

func (t *Tower) Init() {
	t.ids = make(map[AOIEntityID]AOIEntityID)
}

func (t *Tower) Enter(e AOIEntity) bool {
	t.ids[e.ID()] = e.ID()
	if _, ok := t.idsType[e.Type()]; !ok {
		t.idsType[e.Type()] = make(map[AOIEntityID]AOIEntityID)
	}
	entities := t.idsType[e.Type()]
	if _, ok := entities[e.ID()]; ok {
		return false
	}
	t.count++
	entities[e.ID()] = e.ID()
	//t.notifyWatcher(e, AOIEventEnter)
	if e.IsWatcher() {
		t.watchers[e.ID()] = e.ID()
	}
	return true
}

func (t *Tower) Leave(e AOIEntity) bool {
	if _, ok := t.ids[e.ID()]; !ok {
		return false
	}
	delete(t.ids, e.ID())
	if _, ok := t.idsType[e.Type()]; !ok {
		return false
	}
	entities := t.idsType[e.Type()]
	if _, ok := entities[e.ID()]; !ok {
		return false
	}
	t.count--
	delete(entities, e.ID())
	if e.IsWatcher() {
		delete(t.watchers, e.ID())
	}
	//t.notifyWatcher(e, AOIEventLeave)
	return true
}

func (t *Tower) Update(e AOIEntity) bool {
	if _, ok := t.idsType[e.Type()]; !ok {
		return false
	}
	entities := t.idsType[e.Type()]
	if _, ok := entities[e.ID()]; !ok {
		return false
	}
	//t.notifyWatcher(e, AOIEventUpdate)
	return true
}

func (t *Tower) Refresh(typ EntityType, e []AOIEntity) {
	// for _, watcher := range t.watchers {
	// watcher.Notify()
	// }
}

func (t *Tower) Watchers(typ int) map[AOIEntityID]AOIEntityID {
	entities, ok := t.idsType[typ]
	if !ok {
		return nil
	}
	return entities
}

func (t *Tower) getIdsByTypes(types []int) map[int]map[AOIEntityID]AOIEntityID {
	results := make(map[int]map[AOIEntityID]AOIEntityID, 0)
	for i := 0; i < len(types); i++ {
		typ := types[i]
		results[typ] = t.idsType[typ]
	}
	return results
}

func (t *Tower) getIds() map[AOIEntityID]AOIEntityID {
	return t.ids
}

func (t *Tower) entitiesSlice() []AOIEntityID {
	entities := make([]AOIEntityID, 0, t.count)
	for _, typs := range t.idsType {
		for _, entity := range typs {
			entities = append(entities, entity)
		}
	}
	return entities
}
