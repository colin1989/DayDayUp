package delete_slice

import (
	"reflect"
	"testing"
)

var (
	deleteF = func(data *DSlice) bool {
		return data.Id&1 == 1
	}

	want = []*DSlice{
		{2, "2"},
		{4, "4"},
		{6, "6"},
		{8, "8"},
		{10, "10"},
		{12, "12"},
	}
)

func BenchmarkDeleteSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testData := []*DSlice{
			{1, "1"},
			{2, "2"},
			{3, "3"},
			{4, "4"},
			{5, "5"},
			{6, "6"},
			{7, "7"},
			{8, "8"},
			{9, "9"},
			{10, "10"},
			{11, "11"},
			{12, "12"},
		}
		if got := DeleteSlice(testData, deleteF); !reflect.DeepEqual(got, want) {
			b.Errorf("DeleteSlice() = %v, want %v", got, want)
		}
	}
}

func BenchmarkDeleteSlice1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testData := []*DSlice{
			{1, "1"},
			{2, "2"},
			{3, "3"},
			{4, "4"},
			{5, "5"},
			{6, "6"},
			{7, "7"},
			{8, "8"},
			{9, "9"},
			{10, "10"},
			{11, "11"},
			{12, "12"},
		}
		if got := DeleteSlice1(testData, deleteF); !reflect.DeepEqual(got, want) {
			b.Errorf("DeleteSlice1() = %v, want %v", got, want)
		}
	}
}

func BenchmarkDeleteSlice2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testData := []*DSlice{
			{1, "1"},
			{2, "2"},
			{3, "3"},
			{4, "4"},
			{5, "5"},
			{6, "6"},
			{7, "7"},
			{8, "8"},
			{9, "9"},
			{10, "10"},
			{11, "11"},
			{12, "12"},
		}
		if got := DeleteSlice2(testData, deleteF); !reflect.DeepEqual(got, want) {
			b.Errorf("DeleteSlice2() = %v, want %v", got, want)
		}
	}
}

func Test(t *testing.T) {
	//your mock code...
	var s []int

	s = append(s, 1, 2, 3)

	a := s[len(s):]
	// b := s[len(s)]
	c := s[len(s)+1:]

	println(a)
	// println(b)
	println(c)
}
