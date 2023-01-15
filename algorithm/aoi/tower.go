package aoi

type Tower struct {
	count    int
	entities map[int]map[int]AOIEntity
	watchers map[int]AOIEntity
}

func (t *Tower) Enter(e AOIEntity) bool {
	if _, ok := t.entities[e.Type()]; !ok {
		t.entities[e.Type()] = make(map[int]AOIEntity, 0)
	}
	entities := t.entities[e.Type()]
	if _, ok := entities[e.ID()]; ok {
		return false
	}
	t.count++
	entities[e.ID()] = e
	t.notifyWatcher(e, AOIEventEnter)
	return true
}

func (t *Tower) Leave(e AOIEntity) bool {
	if _, ok := t.entities[e.Type()]; !ok {
		return false
	}
	entities := t.entities[e.Type()]
	if _, ok := entities[e.ID()]; !ok {
		return false
	}
	t.count--
	delete(entities, e.ID())
	t.notifyWatcher(e, AOIEventLeave)
	return true
}

func (t *Tower) Update(e AOIEntity) bool {
	if _, ok := t.entities[e.Type()]; !ok {
		return false
	}
	entities := t.entities[e.Type()]
	if _, ok := entities[e.ID()]; !ok {
		return false
	}
	t.notifyWatcher(e, AOIEventUpdate)
	return true
}

func (t *Tower) Refresh(e []AOIEntity, typ EntityType) {
	// for _, watcher := range t.watchers {
	// watcher.Notify()
	// }
}

func (t *Tower) Watchers(typ int) map[int]AOIEntity {
	entities, ok := t.entities[typ]
	if !ok {
		return nil
	}
	return entities
}

func (t *Tower) notifyWatcher(e AOIEntity, typ AOIEventType) {
	for _, watcher := range t.watchers {
		watcher.Notify(e, typ)
	}
}

func (t *Tower) entitiesSlice() []AOIEntity {
	entities := make([]AOIEntity, 0, t.count)
	for _, typs := range t.entities {
		for _, entity := range typs {
			entities = append(entities, entity)
		}
	}
	return entities
}
