package mathp

// MaxI returns the larger of x or y.
func MaxI(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// MaxI32 returns the larger of x or y.
func MaxI32(x, y int32) int32 {
	if x > y {
		return x
	}
	return y
}

// MaxI64 returns the larger of x or y.
func MaxI64(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

// MinI returns the smaller of x or y.
func MinI(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// MinI32 returns the smaller of x or y.
func MinI32(x, y int32) int32 {
	if x < y {
		return x
	}
	return y
}

// MinI64 returns the smaller of x or y.
func MinI64(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}
