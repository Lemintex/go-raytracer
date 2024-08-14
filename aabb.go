package main

type AABB struct {
	X, Y, Z Interval
}

func NewAABBFromPoints(a, b Vec3) AABB {
	aabb := AABB{}
	if a.X <= b.X {
		aabb.X = Interval{a.X, b.X}
	} else {
		aabb.X = Interval{b.X, a.X}
	}

	if a.Y <= b.Y {
		aabb.Y = Interval{a.Y, b.Y}
	} else {
		aabb.Y = Interval{b.Y, a.Y}
	}

	if a.Z <= b.Z {
		aabb.Z = Interval{a.Z, b.Z}
	} else {
		aabb.Z = Interval{b.Z, a.Z}
	}
	aabb.AddPadding()
	return aabb
}

func NewAABBFromAABB(a, b AABB) AABB {
	xMin, xMax := NewInterval(a.X, b.X)
	yMin, yMax := NewInterval(a.Y, b.Y)
	zMin, zMax := NewInterval(a.Z, b.Z)
	return AABB{
		X: Interval{xMin, xMax},
		Y: Interval{yMin, yMax},
		Z: Interval{zMin, zMax},
	}
}

func (aabb AABB) AxisInterval(axis int) Interval {
	switch axis {
	case 0:
		return aabb.X
	case 1:
		return aabb.Y
	case 2:
		return aabb.Z
	}
	panic("Invalid axis")
}

func (aabb AABB) Hit(r Ray, i Interval) bool {
	for index := 0; index < 3; index++ {
		ax := aabb.AxisInterval(index)
		var invD, t0, t1 float64
		switch index {
		case 0:
			invD = 1.0 / r.Direction.X
			t0 = (ax.Min - r.Origin.X) * invD
			t1 = (ax.Max - r.Origin.X) * invD
		case 1:
			invD = 1.0 / r.Direction.Y
			t0 = (ax.Min - r.Origin.Y) * invD
			t1 = (ax.Max - r.Origin.Y) * invD
		case 2:
			invD = 1.0 / r.Direction.Z
			t0 = (ax.Min - r.Origin.Z) * invD
			t1 = (ax.Max - r.Origin.Z) * invD
		}

		if t0 < t1 {
			if t0 > i.Min {
				i.Min = t0
			}
			if t1 < i.Max {
				i.Max = t1
			}
		} else {
			if t1 > i.Min {
				i.Min = t1
			}
			if t0 < i.Max {
				i.Max = t0
			}
		}

		if i.Max <= i.Min {
			return false
		}

	}
	return true
}

func (aabb AABB) LongestAxis() int {
	xSpan := aabb.X.Max - aabb.X.Min
	ySpan := aabb.Y.Max - aabb.Y.Min
	zSpan := aabb.Z.Max - aabb.Z.Min
	if xSpan > ySpan && xSpan > zSpan {
		return 0
	} else if ySpan > zSpan {
		return 1
	}
	return 2
}

// make the bounding box wider if it is too skinny
func (aabb *AABB) AddPadding() {
	delta := 0.00001
	if aabb.X.Size() < delta {
		aabb.X.Expand(delta)
	}
	if aabb.Y.Size() < delta {
		aabb.Y.Expand(delta)
	}
	if aabb.Z.Size() < delta {
		aabb.Z.Expand(delta)
	}
}
