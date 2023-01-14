package delete_slice

type DSlice struct {
	Id   int
	name string
}

func DeleteSlice(a []*DSlice, f func(data *DSlice) bool) []*DSlice {
	for i := 0; i < len(a); i++ {
		if f(a[i]) {
			a = append(a[:i], a[i+1:]...)
			i--
		}
	}
	return a
}

func DeleteSlice1(a []*DSlice, f func(data *DSlice) bool) []*DSlice {
	ret := make([]*DSlice, 0, len(a))
	for _, val := range a {
		if !f(val) {
			ret = append(ret, val)
		}
	}
	return ret
}

func DeleteSlice2(a []*DSlice, f func(data *DSlice) bool) []*DSlice {
	j := 0
	for _, val := range a {
		if !f(val) {
			a[j] = val
			j++
		}
	}
	c := a[:j]
	_ = c
	b := a[j:]
	_ = b
	return a[:j]
}
