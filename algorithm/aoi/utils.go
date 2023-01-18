package aoi

func Map2Slice(results []AOIEntityID, m map[AOIEntityID]AOIEntityID) []AOIEntityID {
	for _, id := range m {
		results = append(results, id)
	}
	return results
}

func TypeMap2Slice(results map[int][]AOIEntityID, m map[int]map[AOIEntityID]AOIEntityID, types []int) map[int][]AOIEntityID {
	for i := 0; i < len(types); i++ {
		typ := types[i]
		ids, ok := m[typ]
		if !ok {
			continue
		}
		if _, ok = results[typ]; !ok {
			results[typ] = make([]AOIEntityID, 0)
		}
		for _, id := range ids {
			results[typ] = append(results[typ], id)
		}
	}
	return results
}
