package Power

func power1(x float64, n int) float64 {
	if x >= -0.00000001 && x <= 0.00000001 && n < 0 {
		return 0.0
	}
	if n < 0 {
		return 1.0 / mathPower(x, -n)
	} else {
		return mathPower(x, n)
	}
}

func mathPower(x float64, n int) float64 {
	if n < 1 {
		return 1
	}
	p := mathPower(x, n>>1)
	if n&1 == 1 {
		return p * p * x
	} else {
		return p * p
	}
}

func myPower(x float64, n int) float64 {
	if x >= -0.00000001 && x <= 0.00000001 && n < 0 {
		return 0.0
	}
	if n < 0 {
		return 1.0 / mathPower1(x, -n)
	} else {
		return mathPower1(x, n)
	}
}

func mathPower1(x float64, n int) float64 {
	out := 1.0
	for n > 0 {
		if n&0x1 == 1 {
			out *= x
		}
		x *= x
		n >>= 1
	}
	return out
}
