package viewarea

import "time"

type ViewArea struct {
	// 中心点坐标
	CenterX, CenterY int

	// 最小坐标
	MinX, MinY int

	// 最大坐标
	MaxX, MaxY int

	UpdateTime time.Time
}

// GetCenterDistant 获取中心点的距离的平方 这里不求根号的原因是 求根运算消耗比较高
func (va *ViewArea) GetCenterDistant(vaTo *ViewArea) int {
	x, y := vaTo.CenterX-va.CenterX, vaTo.CenterY-va.CenterY
	return x*x + y*y
}

func (va *ViewArea) CanSeePos(x, y int) bool {
	return x >= va.MinX && x <= va.MaxX && y >= va.MinY && y <= va.MaxY
}

func (va *ViewArea) CanSeeLine(x1, y1, x2, y2 int) bool {

	// 先简单判断，看得到点的，都能看到线
	if va.CanSeePos(x1, y1) || va.CanSeePos(x2, y2) {
		return true
	}

	// 快速排除，线段的包围盒与矩形是否相交（不相交，则线段与矩形也肯定不相交）
	minX := Min(x1, x2)
	minY := Min(y1, y2)
	maxX := Max(x1, x2)
	maxY := Max(y1, y2)
	if !isRectIntersect(minX, minY, maxX, maxY, va.MinX, va.MinY, va.MaxX, va.MaxY) {
		return false
	}

	// 线与矩形的4条边是否相交，是则相交，否则不相交
	if isSegmentIntersect(x1, y1, x2, y2, va.MinX, va.MinY, va.MinX, va.MaxY) {
		return true
	}
	if isSegmentIntersect(x1, y1, x2, y2, va.MinX, va.MinY, va.MaxX, va.MinY) {
		return true
	}
	if isSegmentIntersect(x1, y1, x2, y2, va.MinX, va.MaxY, va.MaxX, va.MaxY) {
		return true
	}
	if isSegmentIntersect(x1, y1, x2, y2, va.MaxX, va.MinY, va.MaxX, va.MaxY) {
		return true
	}

	return false
}

// 矩形相交
func isRectIntersect(minx1, miny1, maxx1, maxy1, minx2, miny2, maxx2, maxy2 int) bool {
	//假定矩形是用一对点表达的(minx, miny) (maxx, maxy)，那么两个矩形
	//rect1{(minx1, miny1)(maxx1, maxy1)}
	//rect2{(minx2, miny2)(maxx2, maxy2)}
	//相交的结果一定是个矩形，构成这个相交矩形rect{(minx, miny) (maxx, maxy)}的点对坐标是：
	minx := Max(minx1, minx2)
	miny := Max(miny1, miny2)
	maxx := Min(maxx1, maxx2)
	maxy := Min(maxy1, maxy2)

	//如果两个矩形不相交，那么计算得到的点对坐标必然满足：
	//（ minx  >  maxx ） 或者 （ miny  >  maxy ）
	return minx <= maxx && miny <= maxy
}

// 线段相交
func isSegmentIntersect(ax, ay, bx, by, cx, cy, dx, dy int) bool {
	// 线段的两个端点分别在另一条线段的两侧
	return cross(ax, ay, bx, by, cx, cy)*cross(ax, ay, bx, by, dx, dy) <= 0 &&
		cross(cx, cy, dx, dy, ax, ay)*cross(cx, cy, dx, dy, bx, by) <= 0
}

// 求出向量 AB 与向量 AC 的向量积,返回 0 代表共线
// 叉积的概念： 设向量 a(x1, y1) 、 b(x2, y2)
// a x b = x1*y2 - x2*y1
// <0 左侧， >0 右侧， =0 同一直线
func cross(ax, ay, bx, by, cx, cy int) int {
	x1 := bx - ax
	y1 := by - ay
	x2 := cx - ax
	y2 := cy - ay

	return x1*y2 - x2*y1
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Min(x, y int) int {
	if x > y {
		return y
	}

	return x
}

func Minx(a int, b ...int) (min int) {
	min = a
	for _, x := range b {
		min = Min(min, x)
	}
	return
}
