// Copyright 2015 The Golang Plus Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package mathp is a plus to the standard "math" package.
*/
package mathp

// Min finds the mininum value within Len number of elements and returns the
// index of it, given a less function. Less is similar to sort.Interface.Less.
// If Len is zero, 0 is returned.
func Min(Len int, Less func(i, j int) bool) int {
	mn := 0
	for i := 1; i < Len; i++ {
		if Less(i, mn) {
			mn = i
		}
	}
	return mn
}

// Max finds the maximum value within Len number of elements and returns the
// index of it, given a less function. Less is similar to sort.Interface.Less.
// If Len is zero, 0 is returned.
func Max(Len int, Less func(i, j int) bool) int {
	mx := 0
	for i := 1; i < Len; i++ {
		if Less(mx, i) {
			mx = i
		}
	}
	return mx
}

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
