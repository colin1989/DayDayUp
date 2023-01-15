package aoi

//type AOI struct {
//	X int
//	Y int
//}

//type AOICallback interface {
//	OnEnterAOI(other *AOI)
//	OnLeaveAOI(other *AOI)
//}

type AOI interface {
	Enter()
	Leave()
	Moved()
}
